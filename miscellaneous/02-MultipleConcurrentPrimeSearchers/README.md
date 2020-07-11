# 02 Optimal Timing for Multiple Prime Searchers Running Concurrently

This is the exercise _02 of section 06_ from the excellent [repo] _striversity's GO Language for Tourist_

## Details of the exercise

Write a complete golang application which uses **Multiple Prime Searcher** functions running **concurrently** to find all the prime numbers between `1` and `100,000`.

> However, the application _MUST_ end immediately after the _last Prime Search finishes_.

## TODO 1 - Copy the code written for 01-MultiplePrimeSearchers.

_note from earlier code:_
Your data generation function has to do the following:

1. Creates a slice of capacity `100,000`
2. Initialize each element starting at the number `1` until `100,000`

## TODO 2 - Use sync.WaitGroup in stead of time.Sleep in main()

We will replace the `time.Sleep` at the end of the `main` with `sync.WorkGroup.Wait()`

The below should follow an order for `var wg sync.WorkGroup`

- wg.Add(1)
- wg.Done()
- wg.Wait()

_note from earlier code:_
Your Prime Search function has to do the following:

1. Should take as `parameters` a _worker id_ and a _slice of integers_
   
    * TIP: Use the provided `subSlice()` function in `subslice.go`

      (_note: `subslice.go`is already provided, but we will customize it_)
2. For each integer, check if it is a prime number or not.
   
    * NOTE: The order the numbers printed is not important for this exercise.

## TODO 3 - Run between 3 to 10 Prime Searchers


## Check your results:

After running the program, perform the following to check ad confirm:

1. Compile and run your program
2. Count the number of lines in the file, it should be equal to `9,592`.
    * TIP
        * Mac/Linux Users:  `go build -o prime && ./prime | wc -l`
        * Windows Users:    `go build -o prime && ./prime > results.txt && code results.txt`
    * See the entry for number of primes between 1 and 100,000 at https://primes.utm.edu/howmany.html

[repo]: github.com/striversity/glft/