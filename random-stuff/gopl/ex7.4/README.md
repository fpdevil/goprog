# Exercise 7.4

The `stings.NewReader` function returns a value that satisfies the `io.Reader` interface (_and others_) by reading from the argument, a string. Implement a simple version of `NewReader` yourself, and use it to make the *HTML* parser take input from a string.