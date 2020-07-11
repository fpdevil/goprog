# Concurrency, Pipelines and GoRoutines

`goroutines` are the basic units of organization in a `go` program and every `go` program has atleast one `goroutine` which is the `main` and it is automatically created and started when the process begins. `goroutines` are unique to `go` with a deeper integration with the `go's` runtime. They do not define their own _suspension_ or _entry_ points. The runtime automatically _suspends_ them when they block and then resumes them when they are unblocked.


## spawning goroutines

Once spawned, a `goroutine` is just a function that runs concurrently (but need not be parallel) alonside the other code. We can make any function as a `goroutine` by simply qualifying it with a `go` keyword in front of it.

- example
```
func main() {
    go blanket() {
        // do something here
    }
}

func blanket() {
    fmt.Print("a blanket function \n")
}
```

Infact `anonymous` functions can also be qualified as `goroutines` in the same way.

`goroutines` are not the same as `OS` threads and they are not exactly _green_ threads (_which are threads managed by a language's runtime_). They are of a higher level known as `coroutines`.

`go's` mechanism for hosting goroutines is an implementation of `m:n` scheduler which means it maps `m` `green` threads to `n` os level threads. `goroutines` are then scheduled onto the `green` threads. If the `goroutines` exceeed `green` threads, the _scheduler_ will handle the disttribution of `goroutines` across the available threads thereby ensuring that when these `goroutines` become blocked, other `goroutines` can be run.

## Garbage collection

It has to be noted that `goroutines` are not garbage collected automatically, so it's our responsibility to have them properly created and destroyed.

Since goroutines just occupy only a few `kb` of size, spawning millions of `goroutines` will not cause any severe impact during the processing and it's a common practice to spawn such a large number of `goroutines` in production systems depending on the requirement. It's the way concurrency in `go` is implemented.

We can check how much amount of memory is allocated before and after the `goroutine` creation with the code [here](random-stuff/generic/memusage/main.go)

It's an abstract from the excellent book [Concurrency In Go](https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/)

Once run, it shows the below interesting details while creating `30000` goroutines.

```bash
⇒  go run generic/memusage/main.go

/// Memory statistics for spawning multiple goroutines ///
* GO Runtime: go1.14.4
* Spawning 30000 goroutines...
* Resources per each goroutine:
	Memory: 2.377kb
	Time: 110.299342ms
```
