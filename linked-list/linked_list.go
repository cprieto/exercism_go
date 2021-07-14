package linkedlist

import "errors"

var ErrEmptyList = errors.New("list is empty")

// Node represents a node element in a linked list
type Node struct {
	Val  interface{}
	next *Node
	prev *Node
}

// Next returns the next element to current node
func (n *Node) Next() *Node {
	return n.next
}

// Prev returns the previous element to current node
func (n *Node) Prev() *Node {
	return n.prev
}

// List represents a double linked list
type List struct {
	head *Node
	tail *Node
}

// NewList returns a new double linked list with given elements
func NewList(elems ...interface{}) *List {
	l := List{}
	for _, elem := range elems {
		l.PushBack(elem)
	}

	return &l
}

// First returns the first element in the list
func (l *List) First() *Node {
	return l.head
}

// Last returns the last element in the list
func (l *List) Last() *Node {
	return l.tail
}

// PushFront inserts a new item at the beginning of the list
func (l *List) PushFront(elem interface{}) {
	node := &Node{Val: elem, next: l.head}

	if l.head != nil {
		l.head.prev = node
	}

	if l.tail == nil {
		l.tail = node
	}

	l.head = node
}

// PushBack inserts a new item at the end of the list
func (l *List) PushBack(elem interface{}) {
	node := &Node{Val: elem, prev: l.tail}

	if l.tail != nil {
		l.tail.next = node
	}

	if l.head == nil {
		l.head = node
	}

	l.tail = node
}

// Reverse mutates a list with it reverse and returns it
// based on https://www.geeksforgeeks.org/reverse-a-doubly-linked-list/
func (l *List) Reverse() *List {
	current := l.head
	l.head, l.tail = l.tail, l.head

	for current != nil {
		current.prev, current.next = current.next, current.prev
		current = current.prev
	}

	return l
}

// Returns the first element and removes it from the list
func (l *List) PopFront() (interface{}, error) {
	if l.head == nil {
		return nil, ErrEmptyList
	}

	val := l.head.Val
	l.head = l.head.next

	if l.head == nil {
		l.tail = nil
	} else {
		l.head.prev = nil
	}

	return val, nil
}

// PopBack returns the last element and removes it from the list
func (l *List) PopBack() (interface{}, error) {
	if l.tail == nil {
		return nil, ErrEmptyList
	}

	val := l.tail.Val
	l.tail = l.tail.prev

	if l.tail == nil {
		l.head = nil
	} else {
		l.tail.next = nil
	}

	return val, nil
}
