package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

var (
	wg            sync.WaitGroup
	counter       int
	c             = make(chan interface{})
	numGoRoutines = flag.Int("n", 3*1e4, "Number of goroutines to spawn")
)

func noop() {
	wg.Done()
	counter++
	<-c
}

// resourceConsumed returns the totoal number of bytes consumed
// from the underlying OS
func resourceConsumed() uint64 {
	runtime.GC()
	var s runtime.MemStats
	runtime.ReadMemStats(&s)
	return s.Sys
}

func main() {
	flag.Parse()
	if *numGoRoutines <= 0 {
		fmt.Fprintf(os.Stderr, "invalid value for number of goroutines. exiting")
		os.Exit(1)
	}

	fmt.Println("/// Memory statistics for spawning multiple goroutines ///")
	fmt.Printf("* GO Runtime: %s\n", runtime.Version())

	// limit the max. number of CPU's executing simultaneously to 1
	runtime.GOMAXPROCS(1)

	wg.Add(*numGoRoutines)
	start := time.Now()
	before := resourceConsumed()
	for i := *numGoRoutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()

	// now yield the processor to allow other goroutines to run
	runtime.Gosched()
	after := resourceConsumed()
	elapsed := time.Since(start)

	if counter != *numGoRoutines {
		fmt.Fprintf(os.Stderr, "Failed to start goroutine execution")
		os.Exit(1)
	}

	fmt.Printf("* Spawning %d goroutines...\n", *numGoRoutines)
	fmt.Printf("* Resources per each goroutine:\n")
	fmt.Printf("\tMemory: %.3fkb\n", float64(after-before)/float64(*numGoRoutines)/1000)
	fmt.Printf("\tTime: %v\n", elapsed)
}
