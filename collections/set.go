// Package collections contains custom collection types.
package collections

// Set is a set of unique items.
type Set[T comparable] map[T]bool

// NewSet creates a new set.
func NewSet[T comparable]() Set[T] {
	return Set[T]{}
}

// Add adds an item to a set.
func (s Set[T]) Add(item T) {
	s[item] = true
}

// Delete adds an item to a set.
func (s Set[T]) Delete(item T) {
	delete(s, item)
}

// Has checks if a set has an item.
func (s Set[T]) Has(item T) bool {
	return s[item]
}

// Items returns all the items in this set.
func (s Set[T]) Items() []T {
	keys := make([]T, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	return keys
}

// Size returns the size of this set.
func (s Set[T]) Size() int {
	return len(s)
}

// HasItems returns whether or not this set has items.
// Is the opposite of IsEmpty.
func (s Set[T]) HasItems() bool {
	return len(s) > 0
}

// IsEmpty returns whether or not this set is empty.
func (s Set[T]) IsEmpty() bool {
	return len(s) == 0
}
