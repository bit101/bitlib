package collections

import "testing"

func TestStackCreation(t *testing.T) {
	s := NewStack[string]()
	size := s.Size()
	exp := 0
	if size != exp {
		t.Errorf("expected size %d, got %d", exp, size)
	}

	b := s.IsEmpty()
	if !b {
		t.Errorf("expected empty to be %t, got %t", true, b)
	}

	b = s.HasItems()
	if b {
		t.Errorf("expected HasItems to be %t, got %t", false, b)
	}
}

func TestPush(t *testing.T) {
	s := NewStack[string]()
	s.Push("foo")

	size := s.Size()
	exp := 1
	if size != exp {
		t.Errorf("expected size %d, got %d", exp, size)
	}

	s.Push("bar")
	size = s.Size()
	exp = 2
	if size != exp {
		t.Errorf("expected size %d, got %d", exp, size)
	}

	b := s.IsEmpty()
	if b {
		t.Errorf("expected empty to be %t, got %t", false, b)
	}

	b = s.HasItems()
	if !b {
		t.Errorf("expected HasItems to be %t, got %t", true, b)
	}
}

func TestPop(t *testing.T) {
	s := NewStack[string]()
	s.Push("foo")
	s.Push("bar")

	size := s.Size()
	exp := 2
	if size != exp {
		t.Errorf("expected size %d, got %d", exp, size)
	}

	item := s.Pop()
	expItem := "bar"
	if item != expItem {
		t.Errorf("expected item %s, got %s", expItem, item)
	}

	item = s.Pop()
	expItem = "foo"
	if item != expItem {
		t.Errorf("expected item %s, got %s", expItem, item)
	}

	b := s.IsEmpty()
	if !b {
		t.Errorf("expected empty to be %t, got %t", true, b)
	}

	b = s.HasItems()
	if b {
		t.Errorf("expected HasItems to be %t, got %t", false, b)
	}
}

func TestPushAll(t *testing.T) {
	s := NewStack[string]()
	items := []string{"foo", "bar", "baz"}
	s.PushAll(items)

	size := s.Size()
	exp := 3
	if size != exp {
		t.Errorf("expected size %d, got %d", exp, size)
	}

	item := s.Pop()
	expItem := "baz"
	if item != expItem {
		t.Errorf("expected item %s, got %s", expItem, item)
	}

	item = s.Pop()
	expItem = "bar"
	if item != expItem {
		t.Errorf("expected item %s, got %s", expItem, item)
	}

	item = s.Pop()
	expItem = "foo"
	if item != expItem {
		t.Errorf("expected item %s, got %s", expItem, item)
	}

	b := s.IsEmpty()
	if !b {
		t.Errorf("expected empty to be %t, got %t", true, b)
	}

	b = s.HasItems()
	if b {
		t.Errorf("expected HasItems to be %t, got %t", false, b)
	}
}

func TestStackInt(t *testing.T) {
	s := NewStack[int]()
	s.Push(99)

	size := s.Size()
	exp := 1
	if size != exp {
		t.Errorf("expected size %d, got %d", exp, size)
	}

	i := s.Pop()
	if i != 99 {
		t.Errorf("expected size %d, got %d", i, 99)
	}
}

func TestStackObject(t *testing.T) {
	s := NewStack[Queue[string]]()
	q := NewQueue[string]()
	q.Enqueue("success")
	s.Push(q)

	size := s.Size()
	exp := 1
	if size != exp {
		t.Errorf("expected size %d, got %d", exp, size)
	}

	r := s.Pop()
	v := r.Dequeue()

	if v != "success" {
		t.Errorf("expected %s, got %s", "success", v)
	}

}
