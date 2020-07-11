# No channel is ready with work to perform in the middle.

We can use a `default` match in the select statement which can be matched allowing to exit the block without blocking.

Output:

```go
go run main.go
/// No Channels are ready, but need to perform work ///

Matched `default` after 4.633µs
```