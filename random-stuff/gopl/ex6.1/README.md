# Exercise 6.1: Implement these additional methods:

```go
func (*IntSet) Len() int      // return the number of elements
func (*IntSet) Remove(x int)  // remove x from the set
func (*IntSet) Clear()        // remove all elements from the set
func (*IntSet) Copy() *IntSet // return a copy of the set
```

This exercise is in continuation with the `bit vector` discussion from section **6.5**.


## Bitwise operations in GO & Implementation of Set

In a problem involving sets, the elements are selected from some given set called the universal set for that problem. For example, if the set of vowels or the set `{X, Y, Z}` is being considered, the universal set might be the set of all letters. If the universal set is the set of names of months of the year, then one might use the set of summer months `{June, July, August}`; the set of months whose names do not contain the letter **r** `{May, June, July, August}`; or the set of all months having fewer than **30** days `{February}`


`Sets` whose elements are selected from a finite universal set can be represented in computer memory by `bit strings` in which the number of `bits` is equal to the `number of elements` in this *universal set*. Each `bit` corresponds to exactly one element of the `universal set`.

A given `set` is then represented by a `bit string` in which the `bits` corresponding to the elements of that `set` are **1** and all other `bits` are **0**.

A `bitset` is therefore an obvious data structure to use to implement such a `set`.

For instance, if the `universal set` is the set of uppercase letters, then any `set` of
**uppercase** letters may be represented by a string of **26** `bits` with one `bit` correspnding to
the letter **A**, another corrsponding to letter **B** and so on.

Thus, the set of vowels may be represented by the `bit` string,
```sh
1   0   0   0   1   0   0   0   1   0   0   0   0   0   1   0   0   0   0   0   1   0   0   0   0   0
|   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |
A   B   C   D   E   F   G   H   I   J   K   L   M   N   O   P   Q   R   S   T   U   V   W   X   Y   Z
```
and the `empty set` is represented by
```sh
0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0
```

The `bitwise` operations `& | ^` can be used to implement the basic operations in a set like *intersection*, *union* and *compliment*

## Application of bitwise operators

Applying the bitwise operand **&** to the `bit strings` representing two `sets` yields a `bit string` representing the *intersection* of these `sets`. For instance, consider the `sets` `S = {A, B, C, D}` and `T = {A, C, E, G, I}`, where the `universal set` is the `set` of *uppercase* letters. The **26-bit** string representations of these are as below:

```sh
S: 1   1   1   1   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0

T: 1   0   1   0   1   0   1   0   1   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0
   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |
   A   B   C   D   E   F   G   H   I   J   K   L   M   N   O   P   Q   R   S   T   U   V   W   X   Y   Z
```

Now, performing the bitwise `&` operation over the `bit strings` `S` and `T` gives below:

```sh
1   0   1   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0
|   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |
A   B   C   D   E   F   G   H   I   J   K   L   M   N   O   P   Q   R   S   T   U   V   W   X   Y   Z
```
which represents the set `{A, C}` which is the **intersection** of `S` and `T` or `S ∩ T`.

The `bitwise` `|` operation to the `bit strings` representing two sets `S` and `T`
yields the representation of  `S` **union** `T` or `S ∪ T` as below, which is the set `{A, B, C, D, E, G, I}`.

```sh
1   1   1   1   1   0   1   0   1   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0
|   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |
A   B   C   D   E   F   G   H   I   J   K   L   M   N   O   P   Q   R   S   T   U   V   W   X   Y   Z
```

## Here are few examples of the way *bitwise* operations can be useful

- Use bitwise **OR |** to get the bits that are in `1 OR 2` which is *set intersection*.

```go
// 1     = 00000001
// 2     = 00000010
// 1 | 2 = 00000011 = 3
fmt.Println(1 | 2)
```

- Use bitwise **OR |** to get the bits that are in `1 OR 5`

```go
// 1     = 00000001
// 5     = 00000101
// 1 | 5 = 00000101 = 5
fmt.Println(1 | 5)
```

- Use bitwise **XOR ^** to get the bits that are in `3 OR 6` *BUT NOT BOTH*.

```go
// 3     = 00000011
// 6     = 00000110
// 3 ^ 6 = 00000101 = 5
fmt.Println(3 ^ 6)
```

- Use bitwise **AND &** to get the bits that are in `3 AND 6`.

```go
// 3     = 00000011
// 6     = 00000110
// 3 & 6 = 00000010 = 2
fmt.Println(3 & 6)
```

- Use bit clear **AND NOT &^** to get the bits that are in `3 AND NOT 6` (order matters)

```go
// 3      = 00000011
// 6      = 00000110
// 3 &^ 6 = 00000001 = 1
fmt.Println(3 &^ 6)
```


