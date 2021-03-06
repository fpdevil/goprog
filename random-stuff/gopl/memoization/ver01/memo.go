// Package memo provides a concurrency unsafe version of
// memoization of a function of type func
package memo

// Memo caches the results of calling a Func.
type Memo struct {
	f     Func              // function to memoize
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
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}
