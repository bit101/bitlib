// Package geom has geometry related structs and funcs.
package geom

import "testing"

func TestPointListNegativeGet(t *testing.T) {
	list := NewPointList()
	for i := 0.0; i < 10; i++ {
		list.AddXY(i*10, i*10)
	}
	p := list.Get(-1)
	exp := NewPoint(90, 90)
	if !p.Equals(exp) {
		t.Errorf("expected %v, got %v", exp, p)
	}

	p = list.Get(-3)
	exp = NewPoint(70, 70)
	if !p.Equals(exp) {
		t.Errorf("expected %v, got %v", exp, p)
	}

	p = list.Get(-10)
	exp = NewPoint(0, 0)
	if !p.Equals(exp) {
		t.Errorf("expected %v, got %v", exp, p)
	}
}

func TestPointListNegativeGetOutOfRange(t *testing.T) {
	list := NewPointList()
	for i := 0.0; i < 10; i++ {
		list.AddXY(i*10, i*10)
	}

	defer func() { _ = recover() }()
	list.Get(-11)
	t.Errorf("did not panic")
}
