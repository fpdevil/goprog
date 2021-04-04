# Hofstadter sequence

In mathematics, a Hofstadter sequence is a member of a family of related integer sequences defined by non-linear recurrence relations.

## Hofstadter Female and Male sequences

The _Hofstadter_ **Female** `(F)` and **Male** `(M)` sequences are defined as follows:

```js
F(0) = 1 ; M(0) = 0
F(n) = n − M(F(n − 1)) if n > 0
M(n) = n − F(M(n − 1)) if n > 0

// The first few terms of these sequences are
F: 1, 1, 2, 2, 3, 3, 4, 5, 5, 6, 6, 7, 8, 8, 9, 9, 10, 11, 11, 12, 13, ...
M: 0, 0, 1, 2, 2, 3, 4, 4, 5, 6, 6, 7, 7, 8, 9, 9, 10, 11, 11, 12, 12, ...
```


