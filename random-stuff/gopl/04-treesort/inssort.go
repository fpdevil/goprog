package treesort

type tree struct {
	value       int
	left, right *tree // recursive data structure
}

// Sort function performs an inplace sorting of the values
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = insert(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues function appends the elements of the tree t
// to values in order and returns the resulting slice
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func insert(t *tree, value int) *tree {
	if t == nil {
		t = &tree{value: value, left: nil, right: nil}
		// else t = new(tree)
		// t.value = value
		return t
	}

	if value < t.value {
		t.left = insert(t.left, value)
	} else {
		t.right = insert(t.right, value)
	}
	return t
}
