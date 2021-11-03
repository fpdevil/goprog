# Prime number generation

A naive and inefficient of generating a stream of random prime numbers using the pipeline approach. The stream is converted into an integer stream and then passed into the `findPrimes` stage. The function divides the number provided as the input stream by every number between `2` and the `number` below it. If it divides (_not a prime_) the value is passed onto the next stage.

```sh
⇒  go run --race main.go
2021/06/07 22:22:32 * [enter] main *

***************************************************************
A Naive and an inefficient approach of generating prime numbers
using the pipeline patterns.
number of primes: 10 with upper limit 50000000
***************************************************************
Generating Prime Numbers:
	24941317
	36122539
	6410693
	10128161
	25511527
	2107939
	14004383
	7190363
	45931967
	2393161
2021/06/07 22:23:00 * [exit] main took 27.563007668s *
```
