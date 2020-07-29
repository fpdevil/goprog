# To find the duplicate lines

This is an exerpt from the *Go Programming Lamguage Book*.

The program is similar to the unix program `uniq` which looks for duplicate lines adjacent to the current.

## Sample run

* 01

```bash
⇒  cat ../data1.txt | go run main.go
3	cats
2	dogs
2	pigeons
```

* 02

```bash
⇒  go run main.go ../data1.txt
another variant of uniq

3	cats
2	dogs
2	pigeons
```

## Exercise

`ex1.4` has the below exercise from `Exercise 1.4`

```bash
⇒  go run main.go ../uniq/data1.txt ../uniq/data2.txt
3	[../data2.txt]	carrots
2	[../data1.txt]	dogs
2	[../data1.txt ../data2.txt]	rats
2	[../data1.txt ../data2.txt]	mice
4	[../data1.txt ../data2.txt]	pigeons
2	[../data1.txt ../data2.txt]	puppies
6	[../data1.txt ../data2.txt]	cats
2	[../data1.txt ../data2.txt]	squirrels
2	[../data1.txt ../data2.txt]	crows
2	[../data1.txt ../data2.txt]	sparrows
```

Modify the `uniq02` to print the names of all files in which each duplicated line occurs.
