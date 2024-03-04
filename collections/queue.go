// Package collections contains custom collection types.
package collections

// Queue is a FIFO queue of items.
type Queue[T any] []T

// NewQueue creates a new queue.
func NewQueue[T any]() Queue[T] {
	return Queue[T]{}
}

// Enqueue enqueues a single item.
func (q *Queue[T]) Enqueue(item T) {
	*q = append(*q, item)
}

// EnqueueAll enqueues a list of items.
func (q *Queue[T]) EnqueueAll(items []T) {
	*q = append(*q, items...)
}

// Dequeue dequeues the next item on the list.
func (q *Queue[T]) Dequeue() T {
	item := (*q)[0]
	*q = (*q)[1:]
	return item
}

// IsEmpty returns whether or not this queue is empty
func (q Queue[T]) IsEmpty() bool {
	return len(q) == 0
}

// HasItems returns whether or not this queue has items.
// Is the opposite of IsEmpty.
func (q Queue[T]) HasItems() bool {
	return len(q) > 0
}

// Size returns the size of the queue.
func (q Queue[T]) Size() int {
	return len(q)
}
