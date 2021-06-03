package main

import "fmt"

/*
Set Declaration
GO has no built-in data structure for Set, so we will try to mock
the same using a map with the below signature
Set = map[T]bool // $6.5 Example: Bit Vector Type Go Programming Language

for an empty set creation
s := make(Set)

create set with some initial valuess
s := Set{9: true, 13: true, 11: false, 21: true, 30: false}
_, ok := s[13]			// check existence of element
s[25] = true			// add an element to set
delete(s, 25)			// delete an element from set
*/
type Set map[interface{}]bool

// Graph data structure definition
type Graph struct {
	adjacencyList map[interface{}]Set
}

// New function creates a new Graph and initializes the adjacency list
// and then returns the pointer pointing to it
func New() *Graph {
	var gr = Graph{
		adjacencyList: make(map[interface{}]Set),
	}
	return &gr
}

// addEdge method adds an Edge between two nodes provided
// If a node does not exist yet, it will be created and added to Graph
func (gr *Graph) addEdge(nodex, nodey interface{}) {
	if _, ok := gr.adjacencyList[nodex]; !ok {
		gr.adjacencyList[nodex] = make(Set)
	}
	gr.adjacencyList[nodex][nodey] = true
	if _, ok := gr.adjacencyList[nodey]; !ok {
		gr.adjacencyList[nodey] = make(Set)
	}
	gr.adjacencyList[nodey][nodex] = true
}

func dfs(gr *Graph, start interface{}, visited Set, visitFunc func(interface{})) {
	if _, ok := visited[start]; ok {
		return
	}

	visited[start] = true
	visitFunc(start)
	for neighbour := range gr.adjacencyList[start] {
		dfs(gr, neighbour, visited, visitFunc)
	}
}

func main() {
	gr := New()
	gr.addEdge(5, 7)
	gr.addEdge(5, 6)
	gr.addEdge(2, 4)
	gr.addEdge(2, 3)
	gr.addEdge(1, 5)
	gr.addEdge(1, 2)
	visited := make(Set)

	fmt.Println()
	dfs(gr, 1, visited, func(node interface{}) { fmt.Print(node.(int), " ") })

	gr = New()
	gr.addEdge("A", "B")
	gr.addEdge("A", "C")
	gr.addEdge("B", "A")
	gr.addEdge("B", "C")
	gr.addEdge("B", "D")
	gr.addEdge("C", "A")
	gr.addEdge("C", "B")
	gr.addEdge("C", "D")
	gr.addEdge("C", "E")
	gr.addEdge("D", "E")
	gr.addEdge("D", "B")
	gr.addEdge("D", "C")
	gr.addEdge("D", "Y")
	gr.addEdge("E", "C")
	gr.addEdge("E", "D")
	gr.addEdge("F", "D")
	visited = make(Set)
	fmt.Println()
	dfs(gr, "E", visited, func(node interface{}) { fmt.Print(node.(string), " ") })
}
