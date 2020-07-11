# select statement match probability

When multiple channels are being checked for simultaneously in the select statement, the order or the priority at which the match happens for the channels happens at equal proportions between all.

In the example code we start `2` channels which are made ready simultaneously in the select statement.

Here is the output of program run:

```go
go run main.go
* Probability of c1: 49.2%
* Probability of c2: 50.9%

Total time taken: 72.125µs
```

As can be seen the select statement almost gives equal priority to both matches and there is no precedence of one over the other.

Go's runtime performs a pseudo random uniform selection over the set of case statements inside select. Hence each set of case statements will have an equal chance of being selected as all the others as can be evidenced from the sample code.