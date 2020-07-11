# No channels are ready for select

When no channels are ready or when all the channels are blocked, we may get out of the condition using a timeout match as in the code.

When ran here is what it produces:

```go
go run main.go
/// no channels are ready ///

Timed out
```

We have `2` cases in the select for match from a `nil` channel `c` as below

```go
var ch <-chan int
select {
case <-ch:
case <- time.After(1 * time.Second):
    fmt.Println("Timed Out!)
}
```

Being selected from an empty channel, it goes into a blocked state and once the match for `time.After` happens, it just prints the message _Timed Out_.