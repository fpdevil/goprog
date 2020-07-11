# 01 Multiple Prime Searchers Running Concurrently

This is an exercise from the excellent [repo] _striversity's GO Language for Tourist_

## Details of the exercise

Write a complete golang application which uses **Multiple Prime Searcher functions** running concurrently to find all the prime numbers between `1` and `100,000`.

## TODO 1 - Write function which returns a slice of `100,000` integers.

Your data generation function has to do the following:

1. Creates a slice of capacity `100,000`
2. Initialize each element starting at the number `1` until `100,000`

## TODO 2 - Write a Prime Search function

Your Prime Search function has to do the following:

1. Should take as `parameters` a _worker id_ and a _slice of integers_
   
    * TIP: Use the provided `subSlice()` function in `subslice.go`

      (_note: `subslice.go`is already provided, but we will customize it_)
2. For each integer, check if it is a prime number or not.
   
    * NOTE: The order the numbers printed is not important for this exercise.

## TODO 3 - Run between `3` to `10` Prime Searchers

Check your results:

1. Compile and run your program
2. Count the number of lines in the file, it should be equal to `9,592`.
    * TIP
        * Mac/Linux Users:  `go build -o prime && ./prime | wc -l`
        * Windows Users:    `go build -o prime && ./prime > results.txt && code results.txt`
    * See the entry for number of primes between 1 and 100,000 at https://primes.utm.edu/howmany.html

[repo]: github.com/striversity/glft/