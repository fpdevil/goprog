package treesort

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

// ⇒  go test
// PASS
// ok  	github.com/fpdevil/goprog/random-stuff/gopl/04-treesort	0.232s

func TestSort(t *testing.T) {
	now := time.Now().UTC().UnixNano()
	rand.Seed(now)
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Intn(50 + 1)
	}

	Sort(data)
	got := sort.IntsAreSorted(data)
	if !got {
		t.Errorf("data not sorted: %v", data)
	}
}
