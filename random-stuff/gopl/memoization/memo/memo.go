// Package memo provides a concurrency-unsafe
// memoization of a function of type Func.
package memo

// A Memo caches the results of calling a Func.
type Memo struct {
	f     Func
	cache map[string]result
}

// Func will be the type of the function to be memod
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	resp, ok := memo.cache[key]
	if !ok {
		resp.value, resp.err = memo.f(key)
		memo.cache[key] = resp
	}
	return resp.value, resp.err
}
