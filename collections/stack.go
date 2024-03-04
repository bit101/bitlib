// Package collections contains custom collection types.
package collections

// Stack is a LIFO stack of items.
type Stack[T any] []T

// NewStack creates a new stack.
func NewStack[T any]() Stack[T] {
	return Stack[T]{}
}

// Push pushes a single item onto the stack.
func (q *Stack[T]) Push(item T) {
	*q = append(*q, item)
}

// PushAll pushes a list of items onto the stack.
func (q *Stack[T]) PushAll(items []T) {
	*q = append(*q, items...)
}

// Pop pops an item off of the stack.
func (q *Stack[T]) Pop() T {
	end := len(*q) - 1
	item := (*q)[end]
	*q = (*q)[:end]
	return item
}

// IsEmpty returns whether or not this stack is empty
func (q Stack[T]) IsEmpty() bool {
	return len(q) == 0
}

// HasItems returns whether or not this stack has items.
// Is the opposite of IsEmpty.
func (q Stack[T]) HasItems() bool {
	return len(q) > 0
}

// Size returns the size of the stack.
func (q Stack[T]) Size() int {
	return len(q)
}
