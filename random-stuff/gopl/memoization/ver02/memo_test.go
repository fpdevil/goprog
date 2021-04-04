package memo_test

import (
	"testing"

	"github.com/fpdevil/goprog/random-stuff/gopl/memoization/memotest"
	memo "github.com/fpdevil/goprog/random-stuff/gopl/memoization/ver02"
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
⇒  go test -run=TestMemoSequential -v
=== RUN   TestMemoSequential
* https://golang.org            , 424.409403ms, 11298 bytes
* https://godoc.org             , 302.443252ms, 10413 bytes
* https://play.golang.org       , 127.39973ms, 6315 bytes
* http://gopl.io                , 394.293968ms, 4154 bytes
* https://golang.org            , 987ns, 11298 bytes
* https://godoc.org             , 546ns, 10413 bytes
* https://play.golang.org       , 271ns, 6315 bytes
* http://gopl.io                , 401ns, 4154 bytes
--- PASS: TestMemoSequential (1.25s)
PASS
ok  	github.com/fpdevil/goprog/random-stuff/gopl/memoization/ver02	1.488s

⇒  go test -run=TestMemoConcurrent -race -v
=== RUN   TestMemoConcurrent
* https://golang.org            , 467.90977ms, 11298 bytes
* https://godoc.org             , 651.605236ms, 10413 bytes
* https://play.golang.org       , 720.871692ms, 6315 bytes
* http://gopl.io                , 1.139616137s, 4154 bytes
* https://golang.org            , 1.139423056s, 11298 bytes
* https://godoc.org             , 1.139285349s, 10413 bytes
* https://play.golang.org       , 1.139071471s, 6315 bytes
* http://gopl.io                , 1.138892521s, 4154 bytes
--- PASS: TestMemoConcurrent (1.14s)
PASS
ok  	github.com/fpdevil/goprog/random-stuff/gopl/memoization/ver02	1.380s
*/
