# Finding duplicate files

The `showduplicates.go` module searches for all duplicate files at a given location. It uses the
`SHA1` value of files to check if more than one files are same or not. The `SHA-1` secure hash
algorithm produces a *20-byte* hash value for any given chunk of data, such as a file. Similar or
identical files will have the same *SHA-1* values and different files will almost always have
different _SHA-1_ values.

```haskell
⇒  go run showduplicates.go ~/go/src/github.com/fpdevil/goprog/
3 duplicate files [41 bytes]:
	~/go/src/github.com/fpdevil/goprog/.git/ORIG_HEAD
	~/go/src/github.com/fpdevil/goprog/.git/refs/heads/master
	~/go/src/github.com/fpdevil/goprog/.git/refs/remotes/origin/master
2 duplicate files [41 bytes]:
	~/go/src/github.com/fpdevil/goprog/.git/refs/heads/v1
	~/go/src/github.com/fpdevil/goprog/.git/refs/remotes/origin/v1
2 duplicate files [410 bytes]:
	~/go/src/github.com/fpdevil/goprog/miscellaneous/Bank-CSV-Transactions/abc.csv
	~/go/src/github.com/fpdevil/goprog/miscellaneous/Bank-CSV-Transactions/bank_transactions.csv
3 duplicate files [1,168 bytes]:
	~/go/src/github.com/fpdevil/goprog/random-stuff/gopl/memoization/ver03/memo_test.go
	~/go/src/github.com/fpdevil/goprog/random-stuff/gopl/memoization/ver04/memo_test.go
	~/go/src/github.com/fpdevil/goprog/random-stuff/gopl/memoization/ver05/memo_test.go
2 duplicate files [1,237 bytes]:
	~/go/src/github.com/fpdevil/goprog/miscellaneous/01-MultiplePrimeSearchers/subslice.go
	~/go/src/github.com/fpdevil/goprog/miscellaneous/02-MultipleConcurrentPrimeSearchers/subslice.go
3 duplicate files [1,333 bytes]:
	~/go/src/github.com/fpdevil/goprog/miscellaneous/Concurrent-Word-Count/data.txt
	~/go/src/github.com/fpdevil/goprog/random-stuff/Word-Frequencies/02/data.txt
	~/go/src/github.com/fpdevil/goprog/random-stuff/Word-Frequencies/03/data.txt
2 duplicate files [2,874 bytes]:
	~/go/src/github.com/fpdevil/goprog/random-stuff/gopl/05-outline/README.md
	~/go/src/github.com/fpdevil/goprog/random-stuff/gopl/05-outline2/README.md
2 duplicate files [1,675,839 bytes]:
	~/go/src/github.com/fpdevil/goprog/random-stuff/gopl/ex3.1/out.svg
	~/go/src/github.com/fpdevil/goprog/random-stuff/gopl/ex3.2/out.svg
2 duplicate files [1,895,751 bytes]:
	~/go/src/github.com/fpdevil/goprog/random-stuff/gopl/ex3.3/out.svg
	~/go/src/github.com/fpdevil/goprog/random-stuff/gopl/ex8.5/out.svg
```
