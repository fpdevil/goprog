# Depth First Search

A `depth first search` traversal of the tree starts at the root, plunges down the leftmost path, and backtracks only when it gets stuck, returning to the root at the end.

A recursive implementation of the same would follow something like below:

```python
DFS(G,v)   (v is the vertex where the search starts from)
    Stack S := {};   (start with an empty stack)
    for each vertex u, set visited[u] := false;
    push S, v;
    while (S is not empty) do
        u := pop S;
        if (not visited[u]) then
            visited[u] := true;
            for each unvisited neighbour w of u
                push S, w;
        end if
    end while
END DFS()
```
## Running time

Running time of `DFS` on a tree with `n` nodes is given as `T(n) = Θ(1) + ΣT(k)`
