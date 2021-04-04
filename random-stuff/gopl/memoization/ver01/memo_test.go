package memo_test

import (
	"testing"

	"github.com/fpdevil/goprog/random-stuff/gopl/memoization/memotest"
	memo "github.com/fpdevil/goprog/random-stuff/gopl/memoization/ver01"
)

var httpGetBody = memotest.HTTPGetBody

func TestMemoSequential(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

func TestMemoConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}

/* Testing Results
-------------------------------------------------------------
⇒  go test -run=TestMemoConcurrent -v
=== RUN   TestMemoConcurrent
* https://play.golang.org       , 381.193531ms, 6315 bytes
* https://golang.org            , 381.97088ms, 11298 bytes
* https://play.golang.org       , 384.735519ms, 6315 bytes
* https://golang.org            , 385.812983ms, 11298 bytes
* https://godoc.org             , 464.531564ms, 10413 bytes
* https://godoc.org             , 467.965934ms, 10413 bytes
* http://gopl.io                , 487.781125ms, 4154 bytes
* http://gopl.io                , 490.691155ms, 4154 bytes
--- PASS: TestMemoConcurrent (0.49s)
PASS
ok  	github.com/fpdevil/goprog/random-stuff/gopl/memoization/ver01	0.658s

⇒  go test -run=TestMemoSequential -v
=== RUN   TestMemoSequential
* https://golang.org            , 375.895027ms, 11298 bytes
* https://godoc.org             , 197.676418ms, 10413 bytes
* https://play.golang.org       , 78.333181ms, 6315 bytes
* http://gopl.io                , 411.421531ms, 4154 bytes
* https://golang.org            , 604ns, 11298 bytes
* https://godoc.org             , 1.601µs, 10413 bytes
* https://play.golang.org       , 469ns, 6315 bytes
* http://gopl.io                , 309ns, 4154 bytes
--- PASS: TestMemoSequential (1.06s)
PASS
ok  	github.com/fpdevil/goprog/random-stuff/gopl/memoization/ver01	1.232s
*/
