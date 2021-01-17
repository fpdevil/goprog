# Trailing Number of Zeros

A trailing zero is a zero digit in the representation of a number which has no non-zero digits that are less significant than the zero digit. Put more simply, it is a zero digit with no non-zero digits to the right of it.

Underlying mathematics available from [here](https://brilliant.org/wiki/trailing-number-of-zeros/)

## Example

If we wanted to count the number of trailing zeros in 20!

```python
In [5]: math.factorial(20)
Out[5]: 2432902008176640000

# 20! has clearly 4 trailing zeros
```

If an integer is divisible by `10^k`, then it has `k` trailing zeros.

## Usage

```shell
⇒  go run main.go
usage main <start> <end> <step>
usage main 5.0 100 5

⇒  go run main.go 5.0 100 5
  5! has   1 trailing zeros
 10! has   2 trailing zeros
 15! has   3 trailing zeros
 20! has   4 trailing zeros
 25! has   6 trailing zeros
 30! has   7 trailing zeros
 35! has   8 trailing zeros
 40! has   9 trailing zeros
 45! has  10 trailing zeros
 50! has  12 trailing zeros
 55! has  13 trailing zeros
 60! has  14 trailing zeros
 65! has  15 trailing zeros
 70! has  16 trailing zeros
 75! has  18 trailing zeros
 80! has  19 trailing zeros
 85! has  20 trailing zeros
 90! has  21 trailing zeros
 95! has  22 trailing zeros
100! has  24 trailing zeros
```
