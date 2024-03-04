package collections

import (
	"slices"
	"testing"
)

func TestSetCreation(t *testing.T) {
	set := NewSet[string]()

	size := set.Size()
	exp := 0

	if size != exp {
		t.Errorf("expected size %d, got %d", exp, size)
	}

	b := set.IsEmpty()
	if !b {
		t.Errorf("expected empty to be %t, got %t", true, b)
	}

	b = set.Has("foo")
	if b {
		t.Errorf("expected set to NOT have %s", "foo")
	}

	b = set.HasItems()
	if b {
		t.Errorf("expected HasItems to be %t, got %t", false, b)
	}

	items := set.Items()
	size = len(items)
	exp = 0
	if size != exp {
		t.Errorf("expected size %d, got %d", exp, size)
	}
}

func TestSetAdd(t *testing.T) {
	set := NewSet[string]()
	set.Add("foo")

	size := set.Size()
	exp := 1
	if size != exp {
		t.Errorf("expected size %d, got %d", exp, size)
	}

	b := set.Has("foo")
	if !b {
		t.Errorf("expected set to have %s", "foo")
	}

	set.Add("bar")
	set.Add("baz")
	b = set.Has("bar")
	if !b {
		t.Errorf("expected set to have %s", "bar")
	}
	b = set.Has("baz")
	if !b {
		t.Errorf("expected set to have %s", "baz")
	}

	size = set.Size()
	exp = 3
	if size != exp {
		t.Errorf("expected size %d, got %d", exp, size)
	}

	items := set.Items()
	b = slices.Contains(items, "foo")
	if !b {
		t.Errorf("expected set to contain %s", "foo")
	}
	b = slices.Contains(items, "bar")
	if !b {
		t.Errorf("expected set to contain %s", "bar")
	}
	b = slices.Contains(items, "baz")
	if !b {
		t.Errorf("expected set to contain %s", "baz")
	}

	b = set.IsEmpty()
	if b {
		t.Errorf("expected empty to be %t, got %t", false, b)
	}

	b = set.HasItems()
	if !b {
		t.Errorf("expected HasItems to be %t, got %t", true, b)
	}
}

func TestSetDuplicateAdd(t *testing.T) {
	set := NewSet[string]()
	set.Add("foo")
	set.Add("foo")
	set.Add("foo")
	set.Add("foo")
	set.Add("bar")
	set.Add("bar")
	set.Add("bar")

	size := set.Size()
	exp := 2
	if size != exp {
		t.Errorf("expected size %d, got %d", exp, size)
	}
}

func TestSetDelete(t *testing.T) {
	set := NewSet[string]()
	set.Add("foo")
	size := set.Size()
	exp := 1

	if size != exp {
		t.Errorf("expected size %d, got %d", exp, size)
	}
	b := set.Has("foo")
	if !b {
		t.Errorf("expected set to have %s", "foo")
	}

	set.Delete("foo")
	size = set.Size()
	exp = 0

	if size != exp {
		t.Errorf("expected size %d, got %d", exp, size)
	}
	b = set.Has("foo")
	if b {
		t.Errorf("expected set to NOT have %s", "foo")
	}

}
