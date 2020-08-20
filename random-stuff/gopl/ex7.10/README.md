# Exercise 7.10

The `sort.Interface` type can be adapted to other uses. Write a function `IsPalindrome(s sort.Interface) bool` that reports whether the sequence is a palindrome, in other words, reversing the sequence would not change it.

Assume that the elements at indices `i` and `j` are equal if `!s.Less(i, j) && !s.Less(j, i)`.