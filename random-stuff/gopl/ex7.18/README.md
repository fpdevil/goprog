# Chapter 7, Exercise 7.18

Using token based code `API`, write a program that will read an arbitrary `XML` document and construct a `tree` of generic `nodes` that represents it. `Nodes` are of two kinds:

`CharData nodes` represent text strings, and
`Element nodes` represent `named elements` and their attributes.

Each element `node` has a slice of `child nodes`.

You may find the following declarations helpful.

```go
import "encoding/xml

type Node interface {}  // CharData or *Element

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}
```