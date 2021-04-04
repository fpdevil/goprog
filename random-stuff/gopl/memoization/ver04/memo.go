// Package memo provides  a concurrency safe version of  memoization of a
// function  of  type  func.  Requests  for  different  keys  proceed  in
// parallel. The  concurrent requests  for the same  key block  until the
// first one completes.
// This also handles duplicate suppression
package memo

import "sync"

// result represents a pair of results returned by a call to f
// which is a value and an error
type result struct {
	value interface{} // value part of result
	err   error       // error part of result
}

type entry struct {
	res   result        // result of the run
	ready chan struct{} // this channel is closed when res is ready
}

// Memo caches the results of calling a Func.
type Memo struct {
	f     Func              // function to memoize
	mu    sync.Mutex        // a mutex lock
	cache map[string]*entry // cache map from string to entry results
}

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

func New(f Func) *Memo {
	return &Memo{
		f:     f,
		cache: make(map[string]*entry),
	}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// This is the first reqiest for this key.
		// This goroutine becomes responsible for computing
		// the value and broadcasting the ready condition.
		e = &entry{
			ready: make(chan struct{}),
		}
		memo.cache[key] = e
		memo.mu.Unlock()
		e.res.value, e.res.err = memo.f(key)
		close(e.ready)
	} else {
		// this is a repeat request for this key
		memo.mu.Unlock()
		<-e.ready // now wait for ready condition
	}

	return e.res.value, e.res.err
}
