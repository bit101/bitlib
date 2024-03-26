// Package collections contains custom collection types.
package collections

import (
	"cmp"
	"fmt"
)

// MaxPQ is a priority queue that allows getting the maximum value in the queue.
type MaxPQ[T cmp.Ordered] struct {
	pq []T
	n  int
}

// NewMaxPQ creates a new MaxPQ.
func NewMaxPQ[T cmp.Ordered]() *MaxPQ[T] {
	q := MaxPQ[T]{
		pq: []T{},
		n:  0,
	}
	return &q
}

// Insert inserts a new value into the queue.
func (m *MaxPQ[T]) Insert(value T) {
	m.n++
	if len(m.pq) <= m.n {
		newPQ := make([]T, m.n*2)
		copy(newPQ, m.pq)
		m.pq = newPQ
	}
	m.pq[m.n] = value
	m.swim(m.n)
}

// DelMax removes and returns the maximum value in the queue.
func (m *MaxPQ[T]) DelMax() T {
	max := m.pq[1]
	m.pq[1], m.pq[m.n] = m.pq[m.n], m.pq[1]
	m.n--
	m.sink(1)
	return max
}

// Max returns the maximum value in the queue, but does not remove it.
func (m *MaxPQ[T]) Max() T {
	return m.pq[1]
}

// IsEmpty returns whether or not there are any items in the queue.
func (m *MaxPQ[T]) IsEmpty() bool {
	return m.n == 0
}

// Size returns the number of items in the queue.
func (m *MaxPQ[T]) Size() int {
	return m.n
}

func (m *MaxPQ[T]) sink(k int) {
	for k*2 <= m.n {
		j := k * 2
		if j < m.n && m.less(j, j+1) {
			j++
		}
		if !m.less(k, j) {
			break
		}
		m.pq[j], m.pq[k] = m.pq[k], m.pq[j]
		k = j
	}
}

func (m *MaxPQ[T]) swim(k int) {
	for k > 1 && m.less(k/2, k) {
		m.pq[k], m.pq[k/2] = m.pq[k/2], m.pq[k]
		k = k / 2
	}
}

func (m *MaxPQ[T]) less(i, j int) bool {
	return m.pq[i] < m.pq[j]
}

func (m *MaxPQ[T]) String() string {
	s := "["
	for i := 1; i <= m.n; i++ {
		v := m.pq[i]
		s += fmt.Sprintf(" %v", v)
	}
	return s + " ]"
}
