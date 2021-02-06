# Exercise 8.5

Take an existing CPU-bound sequential program, such as the `Mandelbrot` program of _Section 3.3_ or the _3-D surface_ computation of _Section 3.2_, and execute its `main` loop in parallel using channels for communication. How much faster does it run on a multiprocessor machine? What is the optimal number of goroutines to use?

## Usage

```shell
go run main.go > out.svg
```
