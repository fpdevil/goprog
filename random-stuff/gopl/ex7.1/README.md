# Exercise 7.1

Using the ideas from `ByteCounter`, implement counters for words and for lines.
You will find `bufio.ScanWords` useful.

```go
// ByteCounter for counting bytes
type ByteCounter int

// Write method satisfies the io.Writer interface
func (bc *ByteCounter) Write(p []byte) (int, error) {
	*bc += ByteCounter(len(p))
	return len(p), nil
}
```

`
