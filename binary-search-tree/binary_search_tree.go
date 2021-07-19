package binarysearchtree

// SearchTreeData is a node in a binary search tree
type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

// Bst creates a new binary search tree
func Bst(start int) *SearchTreeData {
	return &SearchTreeData{data: start}
}

// Insert insert an element in a binary search tree
func (b *SearchTreeData) Insert(elem int) {
	if elem <= b.data {
		if b.left == nil {
			b.left = &SearchTreeData{data: elem}
		} else {
			b.left.Insert(elem)
		}
	} else {
		if b.right == nil {
			b.right = &SearchTreeData{data: elem}
		} else {
			b.right.Insert(elem)
		}
	}
}

// MapString one day Go will grow up and support generics
func (b *SearchTreeData) MapString(fn func(int) string) []string {
	var result []string

	if b.left == nil {
		result = make([]string, 0)
	} else {
		result = b.left.MapString(fn)
	}

	result = append(result, fn(b.data))

	if b.right != nil {
		result = append(result, b.right.MapString(fn)...)
	}

	return result
}

// MapInt Depth-first in-order traversal of tree returning an array of int
func (b *SearchTreeData) MapInt(fn func(int) int) []int {
	var result []int

	if b.left == nil {
		result = make([]int, 0)
	} else {
		result = b.left.MapInt(fn)
	}

	result = append(result, fn(b.data))

	if b.right != nil {
		result = append(result, b.right.MapInt(fn)...)
	}

	return result
}
