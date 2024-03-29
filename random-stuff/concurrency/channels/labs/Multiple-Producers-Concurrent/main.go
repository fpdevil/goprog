package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	maxMsgsProducer = 100                         // maximum number of messages per producer
	numOfProducers  = 3                           // number of producers
	numOfConsumers  = 2                           // number of consumers
	t               = time.Now().UTC().UnixNano() // timestamp for random seed
	ch              chan string                   // string channel
	wgProducers     sync.WaitGroup                // waitgroup for producers
	wgConsumers     sync.WaitGroup                // waitgroup for consumers
	consumerMutex   sync.Mutex                    // mutex for consumer
)

// PData struct holds the sum and count of the Producer
type PData struct {
	count int
	sum   int
}

func init() {
	flag.IntVar(&maxMsgsProducer, "m", maxMsgsProducer, "provide max number of messages per producer (m > 0)")
	flag.IntVar(&numOfProducers, "np", numOfProducers, "provide number of producers (np > 0)")
	flag.IntVar(&numOfConsumers, "nc", numOfConsumers, "provide number of consumers (nc > 0)")
	flag.Parse()
}

func main() {
	fmt.Println("/// Multiple Producers Running Sequentially ///")
	fmt.Println()

	// show usage message for an invalid value like -1
	if maxMsgsProducer < 1 || numOfProducers < 1 || numOfConsumers < 1 {
		flag.Usage()
		return
	}

	// seed for random number generation
	rand.Seed(t)

	// string channel capable of holding all of the messages
	// that can be produced, which is (np * m)
	ch = make(chan string, maxMsgsProducer*numOfProducers)
	// launch 'n' producers
	for i := 0; i < numOfProducers; i++ {
		producer(i+1, ch)
	}

	// launch 'n' consumers
	for i := 0; i < numOfConsumers; i++ {
		consumer(i+1, ch)
	}

	wgProducers.Wait()
	close(ch)
	wgConsumers.Wait()
}

// producer function takes an integer `id` and generates a random number
// of  messages  between `1` to `m` messages to the `channel` containing
// its `name` or `id` and a `random number`
func producer(id int, out chan string) {
	wgProducers.Add(1)

	go func() {
		m := rand.Intn(maxMsgsProducer) + 1
		for i := 0; i < m; i++ {
			randomN := rand.Intn(25) + 1
			out <- fmt.Sprintf("Producer #%.3v, num: %v", id, randomN)
		}
		wgProducers.Done()
	}()
}

// consumer function is capable of  extracting the producer's name/id and
// random number from  the message, print the number of  messages and the
// sum from  each producer, print the  total number of messages  sent and
// total sum across producers
func consumer(id int, in chan string) {
	wgConsumers.Add(1)

	go func() {
		var (
			sum, count int
		)

		result := make(map[string]PData)

		for p := range in {
			fields := strings.Fields(p)
			l := len(fields)
			pid := fields[1]
			pid = strings.TrimSuffix(pid, ",")
			num, _ := strconv.Atoi(fields[l-1])

			pd := result[pid]
			pd.count++
			pd.sum += num
			result[pid] = pd
		}

		// fmt.Printf("%#v\n", result)
		consumerMutex.Lock()
		fmt.Printf("Consumer %v\n", id)
		for k, v := range result {
			fmt.Printf("\tProducer %v\n\t\tNumber of Elements: %v\n\t\tSub-total: %d\n", k, v.count, v.sum)
			sum += v.sum
			count += v.count
		}
		fmt.Printf("\tTotal count: %d\n", count)
		fmt.Printf("\tTotal sum: %d\n", sum)

		consumerMutex.Unlock()
		wgConsumers.Done()
	}()
}
