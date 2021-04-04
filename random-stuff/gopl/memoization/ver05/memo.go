// Package memo provides  a concurrency safe version of  memoization of a
// function  of  type  func.  Requests  for  different  keys  proceed  in
// parallel.  Concurrent requests  for  same key  block  until the  first
// completes. The implementation uses a monitor goroutine.
package memo

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

// result represents a pair of results returned by a call to f
// which is a value and an error
type result struct {
	value interface{} // value part of result
	err   error       // error part of result
}

type entry struct {
	res   result        // memoized result of f
	ready chan struct{} // closed when res is ready
}

// A request is a message requesting that the Func be applied to key.
type request struct {
	key      string        // argument to the memoized function
	response chan<- result // the client wants a single result
}

// Memo caches the results of calling a Func.
type Memo struct {
	requests chan request
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// This is the first request for this key
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	// broadcast ready condition
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.res
}
