# concurrency patterns

`GO` follows a model of concurrency called the **fork-join** model. *fork* refers to any path in the program which can split off a child branch of execution to be run concurrently with it's parent. While *join* refers to the fact that at some point of time in the future the concurrent branches of execution will again join back the parent `main`.

In example `02` we observe that there was no output displayed when we run the program as below:

```bash
go run main.go
/// SINK: Golang Concurrency Patterns ///

INFO[0000] Counter - starting to work...
```

We call the `numGen` and `counter` which trigger their own `goroutines`, but there was no `join` point. The goroutines executing `numGen` and `counter` will simply exit after some undetermined time in the future and the rest of the program will have already continued executing. Infact it's undetermined whether those `goroutines` will ever be run at all. We can introduce a _sleep_ timer to get the output but it doesn't actually create a join point, it just creates a race condition to increase the probability of running the `goroutines` before main exits. But it does not guarantee the same.

`join` points will guarantee our program's correctness and will remove the race conditions. In order to create `join` points we have to synchronize the `main` goroutine and the `numGen` & `counter` goroutines, which can be done in many ways. But the easiest would be by using a `sync.WaitGroup` which is implemented in example `03`.





```go
func main() {
	fmt.Println("/// SINK: Golang Concurrency Patterns ///")
	fmt.Println()

	// seed for random number generator
	rand.Seed(t)

	// sink - count numbers
	done := make(chan bool)
	d := numGen(done)
	counter(d)
	time.Sleep(1 * time.Second)

	// send a signal to go routine to exit gracefully
	done <- true
}
```