# A fast Fibonacci sequence using Channels

This is the naive `Fibonacci` sequence the computes result fast using `channels`.

```sh
⇒  go run fibonacci.go
usage go run fibonacci <limit>%

# sequence from 0 till 25
⇒  go run fibonacci.go 25
Fibonacci sequence from 0 to 25
* Fibonacci( 0):          0
* Fibonacci( 1):          1
* Fibonacci( 2):          1
* Fibonacci( 3):          2
* Fibonacci( 4):          3
* Fibonacci( 5):          5
* Fibonacci( 6):          8
* Fibonacci( 7):         13
* Fibonacci( 8):         21
* Fibonacci( 9):         34
* Fibonacci(10):         55
* Fibonacci(11):         89
* Fibonacci(12):        144
* Fibonacci(13):        233
* Fibonacci(14):        377
* Fibonacci(15):        610
* Fibonacci(16):        987
* Fibonacci(17):       1597
* Fibonacci(18):       2584
* Fibonacci(19):       4181
* Fibonacci(20):       6765
* Fibonacci(21):      10946
* Fibonacci(22):      17711
* Fibonacci(23):      28657
* Fibonacci(24):      46368
* Fibonacci(25):      75025

Total time elapsed: 133.13µs
```
