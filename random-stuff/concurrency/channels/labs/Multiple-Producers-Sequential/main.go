package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var (
	maxMsgsProducer = 100                         // maximum number of messages per producer
	numOfProducers  = 3                           // number of producers
	t               = time.Now().UTC().UnixNano() // timestamp for random seed
	ch              chan string                   // string channel
)

// PData struct holds the sum and count of the Producer
type PData struct {
	count int
	sum   int
}

func init() {
	flag.IntVar(&maxMsgsProducer, "m", maxMsgsProducer, "provide max number of messages per producer (m > 0)")
	flag.IntVar(&numOfProducers, "np", numOfProducers, "provide number of producers (np > 0)")
	flag.Parse()
}

func main() {
	fmt.Println("/// Multiple Producers Running Sequentially ///")
	fmt.Println()

	// show usage message for an invalid value like -1
	if maxMsgsProducer <= 0 || numOfProducers <= 0 {
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

	close(ch)
	consumer(ch)
}

// producer function takes an integer `id` and generates a random number
// of  messages  between `1` to `m` messages to the `channel` containing
// its `name` or `id` and a `random number`
func producer(id int, out chan string) {
	m := rand.Intn(maxMsgsProducer) + 1
	for i := 0; i < m; i++ {
		randomN := rand.Intn(25) + 1
		out <- fmt.Sprintf("Producer #%.3v, num: %v", id, randomN)
	}
}

// consumer function is capable of  extracting the producer's name/id and
// random number from  the message, print the number of  messages and the
// sum from  each producer, print the  total number of messages  sent and
// total sum across producers
func consumer(in chan string) {

	var (
		sum, count int
	)

	result := make(map[string]PData)

	for p := range in {
		fields := strings.Fields(p)
		l := len(fields)
		id := fields[1]
		id = strings.TrimSuffix(id, ",")
		num, _ := strconv.Atoi(fields[l-1])

		pd := result[id]
		pd.count++
		pd.sum += num
		result[id] = pd
	}
	// fmt.Printf("%#v\n", result)
	for k, v := range result {
		fmt.Printf("Producer %v\n \tNumber of Elements: %v\n\tSub-total: %d\n", k, v.count, v.sum)
		sum += v.sum
		count += v.count
	}
	fmt.Printf("Total count: %d\n", count)
	fmt.Printf("Total sum: %d\n", sum)
}
