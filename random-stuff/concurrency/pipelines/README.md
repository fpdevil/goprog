# Concurrency, Pipelines and GoRoutines

`goroutines` are the basic units of organization in a `go` program and every `go` program has atleast one `goroutine` which is the `main` and it is automatically created and started when the process begins. `goroutines` are unique to `go` with a deeper integration with the `go's` runtime. They do not define their own _suspension_ or _entry_ points. The runtime automatically _suspends_ them when they block and then resumes them when they are unblocked.

`goroutine` is a function that runs concurrently (but need not be parallel) alonside other code. We can make any function as a `goroutine` by simply qualifying it with a `go` keyword in front of it.

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