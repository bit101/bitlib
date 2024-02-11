// Package geom has geometry related structs and funcs.
package geom

import "testing"

func TestSegEqual(t *testing.T) {
	// identical objects
	s0 := NewSegment(100, 100, 200, 200)
	s1 := s0
	b := s0.Equals(s1)
	if !b {
		t.Errorf("Expected %t, got %t\n", true, b)
	}

	// equal values
	s1 = NewSegment(100, 100, 200, 200)
	b = s0.Equals(s1)
	if !b {
		t.Errorf("Expected %t, got %t\n", true, b)
	}

	// unequal
	s1 = NewSegment(100, 100, 200, 200.0001)
	b = s0.Equals(s1)
	if b {
		t.Errorf("Expected %t, got %t\n", false, b)
	}

	// swapped points
	s1 = NewSegment(200, 200, 100, 100)
	b = s0.Equals(s1)
	if !b {
		t.Errorf("Expected %t, got %t\n", true, b)
	}

	// swapped unequal
	s1 = NewSegment(200, 200, 100, 100.0001)
	b = s0.Equals(s1)
	if b {
		t.Errorf("Expected %t, got %t\n", false, b)
	}

}
