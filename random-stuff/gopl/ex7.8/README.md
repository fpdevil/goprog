# Exercise 7.8

Many GUI's provide a table widget with a stateful multi-tier sort:

the _primary sort key_ is the most recently clicked column head,
the _secondary sort key_ is the second most recently clicked column head and so on.

Define an implementation of `sort.Interface` for use by such a table.
Compare that approach with repeated sorting using `sort.Stable`
