# Testing the pipeline 05

We create a filter for enforcing a condition over the random stream of numbers.  The condition or a predicate can be anything line `finding even numbers only` o `finding primes only` etc.

The counter and the proxyCounter are stages which just count the filtered values passed and the total values passed through the pipeline.

## Run

```bash
⇒  go run 05/main.go
/// Processor: Golang Concurrency Patterns Counting & Runnning a Filter ///

INFO[0000] Counter - starting wwork...
INFO[0000] Counter {Proxy} - starting work...
Counter {Proxy} processed 844043 items in 999.82542ms
Counter processed 421618 items in 999.912063ms
```