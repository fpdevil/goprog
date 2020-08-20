# Exercise 7.2

Write a fcuntion `CountingWriter` with the signature below that, given an `io.Writer`, returns a new `Writer` that wraps the original and a pointer to an `int64` variable that at any moment contains the number of bytes written to the new `Writer`.

```go
func CountingWriter(w io.Writer) (io.Writer, *int64)
```