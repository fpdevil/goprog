// Package memo provides a concurrency unsafe version of
// memoization of a function of type func
package memo

import "sync"

// Memo caches the results of calling a Func.
type Memo struct {
	f     Func              // function to memoize
	mu    sync.Mutex        // a mutex lock
	cache map[string]result // cache map from string to results
}

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

// result represents a pair of results returned by a call to f
// which is a value and an error
type result struct {
	value interface{} // value part of result
	err   error       // error part of result
}

func New(f Func) *Memo {
	return &Memo{
		f:     f,
		cache: make(map[string]result),
	}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()
	if !ok {
		res.value, res.err = memo.f(key)

		// between the two critical sections several goroutines
		// may race to compute f(key) in order to update map.
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	return res.value, res.err
}
