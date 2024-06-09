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

func TestPointListCull(t *testing.T) {
	list := NewPointList()
	for range 100 {
		list.Add(RandomPointInRect(0, 0, 100, 100))
	}
	for range 150 {
		list.Add(RandomPointInRect(100, 0, 100, 100))
	}
	count := len(list)
	exp := 250

	if count != exp {
		t.Errorf("expected %d, got %d", exp, count)
	}

	list.Cull(func(p *Point) bool {
		return p.X >= 100
	})

	count = len(list)
	exp = 150

	if count != exp {
		t.Errorf("expected %d, got %d", exp, count)
	}
}

func TestPointListCulled(t *testing.T) {
	list := NewPointList()
	for range 100 {
		list.Add(RandomPointInRect(0, 0, 100, 100))
	}
	for range 150 {
		list.Add(RandomPointInRect(100, 0, 100, 100))
	}
	count := len(list)
	exp := 250

	if count != exp {
		t.Errorf("expected %d, got %d", exp, count)
	}

	list2 := list.Culled(func(p *Point) bool {
		return p.X >= 100
	})

	// no change to original list
	count = len(list)

	if count != exp {
		t.Errorf("expected %d, got %d", exp, count)
	}

	// new list is smaller
	count = len(list2)
	exp = 150

	if count != exp {
		t.Errorf("expected %d, got %d", exp, count)
	}
}

func TestPointListSplit(t *testing.T) {
	list := NewPointList()
	for range 100 {
		list.Add(RandomPointInRect(0, 0, 100, 100))
	}
	for range 150 {
		list.Add(RandomPointInRect(100, 0, 100, 100))
	}
	count := len(list)
	exp := 250

	if count != exp {
		t.Errorf("expected %d, got %d", exp, count)
	}

	removed := list.Split(func(p *Point) bool {
		return p.X >= 100
	})

	count = len(list)
	exp = 150

	countRemoved := len(removed)
	expRemoved := 100

	if count != exp {
		t.Errorf("expected %d, got %d", exp, count)
	}

	if countRemoved != expRemoved {
		t.Errorf("expected %d, got %d", expRemoved, countRemoved)
	}
}
