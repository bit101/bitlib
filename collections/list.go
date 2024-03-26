// Package collections contains custom collection types.
package collections

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/bit101/bitlib/random"
)

// List is a general purpose List class based on slices.
type List[T cmp.Ordered] struct {
	arr []T
}

// NewList creates a new list.
func NewList[T cmp.Ordered]() *List[T] {
	return &List[T]{
		[]T{},
	}
}

// Append appends a value to the end of the list.
func (l *List[T]) Append(value T) {
	l.arr = append(l.arr, value)
}

// Prepend prepends a value to the start of the list.
func (l *List[T]) Prepend(value T) {
	l.arr = append([]T{value}, l.arr...)
}

// Get returns the indexed value from the list.
func (l *List[T]) Get(index int) (T, error) {
	if index >= len(l.arr) {
		var value T
		return value, fmt.Errorf("index %d out of bounds, list length %d", index, len(l.arr))
	}
	if index < 0 {
		var value T
		return value, fmt.Errorf("negative index not possible")
	}
	return l.arr[index], nil
}

// GetRandom returns a random element from the list.
func (l *List[T]) GetRandom() T {
	index := random.IntRange(0, len(l.arr))
	return l.arr[index]
}

// GetFirst returns the first element in the list.
func (l *List[T]) GetFirst() T {
	return l.arr[0]
}

// GetLast returns the last element in the list.
func (l *List[T]) GetLast() T {
	return l.arr[len(l.arr)-1]
}

// Remove removes and returns the indexed value from the list.
func (l *List[T]) Remove(index int) (T, error) {
	if index >= len(l.arr) {
		var value T
		return value, fmt.Errorf("index %d out of bounds, list length %d", index, len(l.arr))
	}
	if index < 0 {
		var value T
		return value, fmt.Errorf("negative index not possible")
	}
	value := l.arr[index]
	l.arr = slices.Delete(l.arr, index, index+1)
	return value, nil
}

// RemoveFirst removes and returns the first value on the list.
func (l *List[T]) RemoveFirst() T {
	value := l.arr[0]
	l.arr = l.arr[1:]
	return value
}

// RemoveLast removes and returns the last value on the list.
func (l *List[T]) RemoveLast() T {
	i := len(l.arr) - 1
	value := l.arr[i]
	l.arr = l.arr[:i]
	return value
}

// Insert inserts a value at the given index.
func (l *List[T]) Insert(index int, value T) error {
	if index >= len(l.arr) {
		return fmt.Errorf("index %d out of bounds, list length %d", index, len(l.arr))
	}
	if index < 0 {
		return fmt.Errorf("negative index not possible")
	}
	l.arr = slices.Insert(l.arr, index, value)
	return nil
}

// Shuffle randomly shuffles the list.
func (l *List[T]) Shuffle() {
	for i := 0; i < len(l.arr); i++ {
		n := random.IntRange(0, i+1)
		l.arr[i], l.arr[n] = l.arr[n], l.arr[i]
	}
}

// Length returns the number of elements in the list.
func (l *List[T]) Length() int {
	return len(l.arr)
}

// Contains returns whether or not the list contains the specified value.
func (l *List[T]) Contains(value T) bool {
	return slices.Contains(l.arr, value)
}

// IndexOf returns the index of the speified value if it is in the list, or -1 if not.
func (l *List[T]) IndexOf(value T) int {
	for i, val := range l.arr {
		if val == value {
			return i
		}
	}
	return -1
}

// Reverse reverses the list.
func (l *List[T]) Reverse() {
	slices.Reverse(l.arr)
}

// Sort sorts the list.
func (l *List[T]) Sort() {
	slices.Sort(l.arr)
}

// Filter returns a new list with the elements that pass the filter function.
func (l *List[T]) Filter(filterFunc func(T) bool) *List[T] {
	ret := NewList[T]()
	for i := 0; i < len(l.arr); i++ {
		if filterFunc(l.arr[i]) {
			ret.Append(l.arr[i])
		}
	}
	return ret
}

// Map returns a new list with elements obtained by applying the map function to each element.
func (l *List[T]) Map(mapFunc func(T) T) *List[T] {
	ret := NewList[T]()
	for i := 0; i < len(l.arr); i++ {
		ret.Append(mapFunc(l.arr[i]))
	}
	return ret
}

// String returns a string representation of the list.
func (l *List[T]) String() string {
	return fmt.Sprintf("%v", l.arr)
}

// Slice returns a slice containing all the values in this list.
func (l *List[T]) Slice() []T {
	return slices.Clone(l.arr)
}

// Clear removes all elements from the list.
func (l *List[T]) Clear() {
	l.arr = []T{}
}
