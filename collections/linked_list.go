// Package collections contains custom collection types.
package collections

import (
	"errors"
	"fmt"
)

// ListNode represents a single element in a linked list.
type ListNode[T any] struct {
	Value T
	Next  *ListNode[T]
	Prev  *ListNode[T]
}

// NewListNode creates a new ListNode.
func NewListNode[T any](value T) *ListNode[T] {
	return &ListNode[T]{
		Value: value,
	}
}

// LinkedList is a doubly linked list.
type LinkedList[T any] struct {
	Head *ListNode[T]
	Tail *ListNode[T]
}

// NewLinkedList creates a new linked list.
func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Add adds an element to the linked list.
// This element will be added to the top of the list for speed.
func (l *LinkedList[T]) Add(value T) {
	node := NewListNode[T](value)
	if l.Head == nil {
		l.Head = node
		l.Tail = l.Head
	} else {
		node.Next = l.Head
		l.Head.Prev = node
		l.Head = node
	}
}

// Append adds an element to the linked list.
// This element will be added to the end of the list.
func (l *LinkedList[T]) Append(value T) {
	if l.Head == nil {
		l.Add(value)
	} else {
		node := NewListNode[T](value)
		l.Tail.Next = node
		node.Prev = l.Tail
		l.Tail = node
	}
}

// Insert adds an element to the linked list.
// This element will be added at the specified index.
func (l *LinkedList[T]) Insert(value T, index int) error {
	if index < 0 {
		return fmt.Errorf("unable to insert to negative index")
	}
	// just add as head
	if index == 0 {
		l.Add(value)
		return nil
	}
	node := NewListNode[T](value)
	currentNode := l.Head
	// go to the node before what we are looking for (index-1)
	for i := 0; i < index-1; i++ {
		currentNode = currentNode.Next
		if currentNode == nil {
			// if we get here, index is too high.
			return fmt.Errorf("unable to add item at index %d. max is %d", index, i)
		}
	}
	// adjust node
	node.Next = currentNode.Next
	node.Prev = currentNode
	// adjust next (if exists)
	if node.Next != nil {
		node.Next.Prev = node
	}
	// adjust current
	currentNode.Next = node
	if currentNode == l.Tail {
		l.Tail = node
	}
	return nil
}

// GetValueAt returns the value in the node at the specified index.
func (l *LinkedList[T]) GetValueAt(index int) (T, error) {
	if index < 0 {
		var ret T
		return ret, errors.New("unable to get item at negative index")
	}
	if index == 0 {
		return l.Head.Value, nil
	}

	currentNode := l.Head
	// go to the node we are looking for
	for i := 0; i < index; i++ {
		currentNode = currentNode.Next
		if currentNode == nil {
			// if we get here, index is too high.
			var ret T
			return ret, fmt.Errorf("unable to get item at index %d. max is %d", index, i)
		}
	}
	return currentNode.Value, nil
}
