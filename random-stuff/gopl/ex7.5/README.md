# Exercise 7.5

The `LimitReader` function in the `io` package accepts an `io.Reader r` and a number of bytes `n`, and returns another `Reader` that reads from `r` but reports an *end-of-file* condition after `n` bytes.

Implement it.
