package collections

import (
	"testing"
)

func TestLinkedListCreation(t *testing.T) {
	l := NewLinkedList[string]()

	if l.Head != nil {
		t.Errorf("expected head to be nil")
	}
	if l.Tail != nil {
		t.Errorf("expected tail to be nil")
	}
}

func TestLinkedListAdd(t *testing.T) {
	l := NewLinkedList[string]()
	l.Add("foo")

	exp := "foo"
	val := l.Head.Value
	if exp != val {
		t.Errorf("expected %s, got %s", exp, val)
	}
	val = l.Tail.Value
	if exp != val {
		t.Errorf("expected %s, got %s", exp, val)
	}

	l.Add("bar")
	exp = "bar"
	val = l.Head.Value
	if exp != val {
		t.Errorf("expected %s, got %s", exp, val)
	}
	exp = "foo"
	val = l.Tail.Value
	if exp != val {
		t.Errorf("expected %s, got %s", exp, val)
	}

	l.Add("baz")
	exp = "baz"
	val = l.Head.Value
	if exp != val {
		t.Errorf("expected %s, got %s", exp, val)
	}
	exp = "foo"
	val = l.Tail.Value
	if exp != val {
		t.Errorf("expected %s, got %s", exp, val)
	}
}

func TestLinkedListAppend(t *testing.T) {
	l := NewLinkedList[string]()
	l.Append("foo")

	exp := "foo"
	val := l.Head.Value
	if exp != val {
		t.Errorf("expected %s, got %s", exp, val)
	}
	val = l.Tail.Value
	if exp != val {
		t.Errorf("expected %s, got %s", exp, val)
	}

	l.Append("bar")
	exp = "foo"
	val = l.Head.Value
	if exp != val {
		t.Errorf("expected %s, got %s", exp, val)
	}
	exp = "bar"
	val = l.Tail.Value
	if exp != val {
		t.Errorf("expected %s, got %s", exp, val)
	}

	l.Append("baz")
	exp = "foo"
	val = l.Head.Value
	if exp != val {
		t.Errorf("expected %s, got %s", exp, val)
	}
	exp = "baz"
	val = l.Tail.Value
	if exp != val {
		t.Errorf("expected %s, got %s", exp, val)
	}

}

func createList() *LinkedList[string] {
	l := NewLinkedList[string]()
	l.Append("000")
	l.Append("111")
	l.Append("222")
	l.Append("333")
	l.Append("444")
	l.Append("555")
	return l
}

func TestLinkedListInsert(t *testing.T) {
	// insert at 0
	l := createList()
	err := l.Insert("foo", 0)
	if err != nil {
		t.Error(err)
	}

	exp := "foo"
	val := l.Head.Value
	if exp != val {
		t.Errorf("expected %s, got %s", exp, val)
	}
	val, _ = l.GetValueAt(0)
	if exp != val {
		t.Errorf("expected %s, got %s", exp, val)
	}

	// insert at 1
	l = createList()
	err = l.Insert("foo", 1)
	if err != nil {
		t.Error(err)
	}
	exp = "foo"
	val, _ = l.GetValueAt(1)
	if exp != val {
		t.Errorf("expected %s, got %s", exp, val)
	}

	exp = "000"
	val = l.Head.Value
	if exp != val {
		t.Errorf("expected %s, got %s", exp, val)
	}
	val, _ = l.GetValueAt(0)
	if exp != val {
		t.Errorf("expected %s, got %s", exp, val)
	}

	// insert at 5
	l = createList()
	err = l.Insert("foo", 5)
	if err != nil {
		t.Error(err)
	}
	exp = "foo"
	val, _ = l.GetValueAt(5)
	if exp != val {
		t.Errorf("expected %s, got %s", exp, val)
	}

	exp = "000"
	val = l.Head.Value
	if exp != val {
		t.Errorf("expected %s, got %s", exp, val)
	}
	val, _ = l.GetValueAt(0)
	if exp != val {
		t.Errorf("expected %s, got %s", exp, val)
	}
	exp = "555"
	val, _ = l.GetValueAt(6)
	if exp != val {
		t.Errorf("expected %s, got %s", exp, val)
	}

	// insert at 6
	l = createList()
	err = l.Insert("foo", 6)
	if err != nil {
		t.Error(err)
	}
	exp = "foo"
	val, _ = l.GetValueAt(6)
	if exp != val {
		t.Errorf("expected %s, got %s", exp, val)
	}

	// insert at 7
	l = createList()
	err = l.Insert("foo", 7)
	if err == nil {
		t.Error("expected error, got none.")
	}

	// insert at -1
	l = createList()
	err = l.Insert("foo", -1)
	if err == nil {
		t.Error("expected error, got none.")
	}
}

func TestLinkedListGet(t *testing.T) {
	// insert at 0
	l := createList()

	exp := "000"
	val, err := l.GetValueAt(0)
	if err != nil {
		t.Error(err)
	}
	if exp != val {
		t.Errorf("expected %s, got %s", exp, val)
	}

	exp = "111"
	val, err = l.GetValueAt(1)
	if err != nil {
		t.Error(err)
	}
	if exp != val {
		t.Errorf("expected %s, got %s", exp, val)
	}

	exp = "555"
	val, err = l.GetValueAt(5)
	if err != nil {
		t.Error(err)
	}
	if exp != val {
		t.Errorf("expected %s, got %s", exp, val)
	}

	_, err = l.GetValueAt(-1)
	if err == nil {
		t.Error("expected error, got none")
	}

	_, err = l.GetValueAt(6)
	if err == nil {
		t.Error("expected error, got none")
	}
}
