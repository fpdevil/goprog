# Pipelines

The task  performed by the `pipeline.go`  program is to generate  random numbers
with in  a given  range and  then stop when  any number  in the  random sequence
appears for a second time.

However, before terminating the program will print the sum of all random numbers
that appeared up  to the point where  the first random number  appeared a second
time. We  need three functions for  connecting the channels of  the program. The
logic of the  program is found in  these three functions, but the  data flows in
the channels of the pipeline.

The program will  have two channels. The  first channel `A` will be  used to get
random numbers from the first function and send them to the second function. The
second channel  `B` will be responsible  for getting the data  from channel `B`,
calculating and rendering the results.

## Test run results

```shell
⇒  go run pipeline.go 10 100
58 99 58
sum of random numbers is 157

⇒  go run pipeline.go 100 200
167 119 106 112 159 125 158 198 103 194 116 157 111 185 101 104 186 125 143
sum of random numbers is 2544
```
