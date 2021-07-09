package tree

import "errors"

// Node represents a node tree
type Node struct {
	ID int
	Children []*Node
}

// Record represents an entry
type Record struct {
	ID int
	Parent int
}

// Build construct a tree using a collection of Record
func Build(input []Record) (*Node, error) {
	num := len(input)

	if num == 0 {
		return nil, nil
	}

	parents := make([]int, num)
	nodes := make([]*Node, num)
	for _, rec := range input {
		switch {
		case rec.ID >= num:
			return nil, errors.New("this id is beyond the number of elements")
		case rec.ID == 0 && rec.Parent != 0:
			return nil, errors.New("root node cannot have a parent")
		case rec.ID != 0 && rec.ID <= rec.Parent:
			return nil, errors.New("recurrent or invalid parent id")
		case nodes[rec.ID] != nil:
			return nil, errors.New("duplicated entries")
		}

		parents[rec.ID] = rec.Parent
		nodes[rec.ID] = &Node{ID: rec.ID}
	}

	for idx, node := range nodes {
		if idx == 0 {
			continue
		}
		pid := parents[idx]
		nodes[pid].Children = append(nodes[pid].Children, node)
	}

	return nodes[0], nil
}