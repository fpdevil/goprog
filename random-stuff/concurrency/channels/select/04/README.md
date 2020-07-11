# Progress of goroutines while waiting for others

Here we consider a `goroutine` spawned which waits for `6` seconds before close.

In the select statement match the channel and the default will be matched and task will be performed occassionally during the default match

Output:

```go
go run main.go

	Allowing a goroutine to make some progress on work while
	waiting for another goroutine to report the result(s)


Performed 6 cycles of CPU (work) prior to exiting
Time elapsed: 6.025530884s
```