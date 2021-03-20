# Concurrency, Pipelines and GoRoutines

`goroutines` are the basic units of organization in a `go` program and every `go` program has at least one `goroutine` which is the `main` and it is automatically created and started when the process begins. `goroutines` are unique to `go` with a deeper integration with the `go's` run-time. They do not define their own _suspension_ or _entry_ points. The run-time automatically _suspends_ them when they block and then resumes them once they are unblocked.

## spawning processes as `goroutines`

Once spawned, a `goroutine` is just like a normal function except that it runs concurrently (but need not be parallel) along side the other code. We can make any function as a `goroutine` by simply qualifying it with the keyword `go` in front of it.

- Here is an example showing how the `blanket` function is spawned as a `goroutine` in `main`.
```go
func main() {
    // blanket spawned as a go routine here
    go blanket() {
        // code goes here...
    }
}

func blanket() {
    fmt.Print("a blanket function \n")
}
```
In fact `anonymous` functions may also be qualified as `goroutines` in the same way.

`goroutines` are not the same as `OS` level threads and they are not exactly _green_ threads (*that are threads managed by a language's run-time*). They are of a higher level known as *coroutines*.

`go's` mechanism for hosting the `goroutines` is an implementation of `m:n` scheduler which means it maps `m` `green` threads to `n` `os` level threads.

The `goroutines` are then scheduled onto the `green` threads. If the `goroutines` exceed `green` threads, the _scheduler_ will handle the distribution of `goroutines` across the available threads thereby ensuring that when these `goroutines` become blocked, other `goroutines` can be run.

## Garbage collection

It has to be noted that `goroutines` are not garbage collected automatically, so it's our responsibility to have them properly created and destroyed.

Since `goroutines` just occupy only a few `kb` of size, spawning millions of `goroutines` will not cause any severe impact during the processing and it's a common practice to spawn such a large number of `goroutines` in production systems depending on the requirement. It's the way concurrency in `go` is implemented.

We can check how much amount of memory is allocated before and after the `goroutine` creation with the code [here](random-stuff/generic/memusage/main.go)

It's an abstract from the excellent book [Concurrency In Go](https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/)

Once run, it shows the below interesting details while creating `30000` `goroutines`.
```shell
# run the go code
⇒  go run generic/memusage/main.go

# Memory statistics for spawning multiple goroutines
* GO Runtime: go1.14.4
* Spawning 30000 goroutines...
* Resources per each goroutine:
	Memory: 2.377kb
	Time: 110.299342ms
```

## `Waitgroup` from the `sync` package

Using `WaitGroup` from the `sync` package is an excellent way to ensure that a set of concurrent operations complete when we do not care about the result of the concurrent operation or we have other means of collecting the results of running.

In case if we do care about the results or we do not have any explicit way of capturing the results, it's better to use **channels** with **select** statements.

Here is a way of applying `WaitGroup` to the example we listed earlier.
```go
import "sync"

var wg sync.WaitGroup

func main() {
    go blanket() {
        // do something here
    }
    wg.Wait()
}

func blanket() {
    wg.Add(1)
    fmt.Print("a blanket function \n")
    defer wg.Done()
}
```

Essentially, `WaitGroup` may be considered as a *Concurrency Safe* counter. All the calls to `wg.Add()` method above, would increment the counter by an integer passed in and all calls to `wg.Done()` decrements the counter by `1`. Calls to the `wg.Wait()` will block until the counter is back to `0`.