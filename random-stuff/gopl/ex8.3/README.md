# Exercise 8.3

 In `netcat3`, the interface value `conn` has the concrete type `net.TCPConn`, which represents a `TCP` connection. A TCP connection consists of two halves that may be closed independently using its `CloseRead` and `CloseWrite` methods. Modify the main goroutine of `netcat3` to close only the write half of the connection so that the program will continue to print the final  `echoes` from the `reverb` server even after the standard input has been closed.

## Documentation from go docs

```go
func (c *TCPConn) CloseRead() error
    CloseRead shuts down the reading side of the TCP connection. Most callers
    should just use Close.

func (c *TCPConn) CloseWrite() error
    CloseWrite shuts down the writing side of the TCP connection. Most callers
    should just use Close.
```
