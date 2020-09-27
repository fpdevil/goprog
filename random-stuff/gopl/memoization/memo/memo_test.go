package memo_test

import (
	"testing"

	"github.com/fpdevil/goprog/random-stuff/gopl/memoization/memo"
	"github.com/fpdevil/goprog/random-stuff/gopl/memoization/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}
