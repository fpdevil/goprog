# Exercise 5.10

Rewrite the `topoSort` to use `maps` instead of `slices` and eliminate the _initial sort_. Verify the results, though nondeterministic, are valid topological orderings.

## Description of original problem

Consider the problem of computing a sequence of computer science courses that satisfy the prerequisite requirements of each one. The prerequisites given in a table mapping each course to the list of courses that must be completed first is given. From which we will have to order the priority of the courses indicating which one has to be completed first in order.

This is a type of problem called as a `topological sorting`. It forms a directed graph with a node for each course and edges from each course to the courses that depend over it.