package main

/*
channel of channel
- a channel on which `channel of type T` is sent or received
*/

import (
	"fmt"
	"math/rand"
	"time"
)

// WorkUnit is a type for integer channel
type (
	WorkUnit chan int
)

var (
	chanCap = 5
)

func main() {
	fmt.Println("/// channel of channels ///")
	fmt.Println()

	// random number seeding
	t := time.Now().UTC().UnixNano()
	rand.Seed(t)

	// processing channel of channel
	queue := generateWork()
	// now create workers to do some work using the WorkUnit
	id := 1
	for wu := range queue {
		worker(id, wu)
		id++
	}
}

func generateWork() (out <-chan WorkUnit) {
	// ch is a channel which takes channels
	ch := make(chan WorkUnit, chanCap)
	for i := 0; i < cap(ch); i++ {
		ch <- genWorkUnit(rand.Intn(10))
	}
	close(ch)
	out = ch
	return
}

// genWorkUnit create some work (feeding channels) for a worker
func genWorkUnit(r int) (out WorkUnit) {
	out = make(WorkUnit, r)
	for i := 0; i < r; i++ {
		out <- rand.Intn(100)
	}
	close(out)
	return
}

// worker performs some calculations on a WorkUnit
func worker(id int, wu WorkUnit) {
	fmt.Printf("* Starting WorkUnit #%v\n", id)

	var (
		sum, count int
	)

	for v := range wu {
		count++
		sum += v
	}

	fmt.Printf("\tProcessing %v values for WorkUnit: %v\n", count, id)
	if count > 0 {
		fmt.Printf("\tTotal: %v, Average: %v\n", sum, sum/count)
	}
}
