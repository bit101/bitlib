// Package collections contains custom collection types.
package collections

import (
	"cmp"
	"fmt"
)

// MinPQ is a priority queue that allows getting the minimum value in the queue.
type MinPQ[T cmp.Ordered] struct {
	pq []T
	n  int
}

// NewMinPQ creates a new MinPQ.
func NewMinPQ[T cmp.Ordered]() *MinPQ[T] {
	q := MinPQ[T]{
		pq: []T{},
		n:  0,
	}
	return &q
}

// Insert inserts a new value into the queue.
func (m *MinPQ[T]) Insert(value T) {
	m.n++
	if len(m.pq) <= m.n {
		newPQ := make([]T, m.n*2)
		copy(newPQ, m.pq)
		m.pq = newPQ
	}
	m.pq[m.n] = value
	m.swim(m.n)
}

// DelMin removes and returns the minimum value in the queue.
func (m *MinPQ[T]) DelMin() T {
	max := m.pq[1]
	m.pq[1], m.pq[m.n] = m.pq[m.n], m.pq[1]
	m.n--
	m.sink(1)
	return max
}

// Min returns the minimum value in the queue, but does not remove it.
func (m *MinPQ[T]) Min() T {
	return m.pq[1]
}

// IsEmpty returns whether or not there are any items in the queue.
func (m *MinPQ[T]) IsEmpty() bool {
	return m.n == 0
}

// Size returns the number of items in the queue.
func (m *MinPQ[T]) Size() int {
	return m.n
}

func (m *MinPQ[T]) sink(k int) {
	for k*2 <= m.n {
		j := k * 2
		if j < m.n && !m.less(j, j+1) {
			j++
		}
		if m.less(k, j) {
			break
		}
		m.pq[j], m.pq[k] = m.pq[k], m.pq[j]
		k = j
	}
}

func (m *MinPQ[T]) swim(k int) {
	for k > 1 && !m.less(k/2, k) {
		m.pq[k], m.pq[k/2] = m.pq[k/2], m.pq[k]
		k = k / 2
	}
}

func (m *MinPQ[T]) less(i, j int) bool {
	return m.pq[i] < m.pq[j]
}

func (m *MinPQ[T]) String() string {
	s := "["
	for i := 1; i <= m.n; i++ {
		v := m.pq[i]
		s += fmt.Sprintf(" %v", v)
	}
	return s + " ]"
}
