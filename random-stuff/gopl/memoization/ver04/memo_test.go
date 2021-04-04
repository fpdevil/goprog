package memo_test

import (
	"testing"

	"github.com/fpdevil/goprog/random-stuff/gopl/memoization/memotest"
	memo "github.com/fpdevil/goprog/random-stuff/gopl/memoization/ver03"
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
⇒  go test -run=TestMemoConcurrent -race -v
=== RUN   TestMemoConcurrent
* https://play.golang.org       , 501.550999ms, 6315 bytes
* https://golang.org            , 505.265558ms, 11298 bytes
* https://play.golang.org       , 506.406803ms, 6315 bytes
* https://golang.org            , 506.514975ms, 11298 bytes
* http://gopl.io                , 641.025706ms, 4154 bytes
* http://gopl.io                , 641.224659ms, 4154 bytes
* https://godoc.org             , 681.479746ms, 10413 bytes
* https://godoc.org             , 755.456553ms, 10413 bytes
--- PASS: TestMemoConcurrent (0.76s)
PASS
ok  	github.com/fpdevil/goprog/random-stuff/gopl/memoization/ver03	1.030s
*/
