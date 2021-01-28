# Exercise 3.12

Write a function that reports whether two strings are anagrams of each other, that is, they contain the same letters in a different order.

## Test

```shell
⇒  go test -v .
=== RUN   TestIsAnagram
--- PASS: TestIsAnagram (0.00s)
PASS
ok  	github.com/fpdevil/goprog/random-stuff/gopl/ex3.12	0.090s

⇒  go test -v .
=== RUN   TestIsAnagram
    main_test.go:18: isAnagram("oranges", "potatoes"), received false and needed true
--- FAIL: TestIsAnagram (0.00s)
FAIL
FAIL	github.com/fpdevil/goprog/random-stuff/gopl/ex3.12	0.155s
FAIL
```
