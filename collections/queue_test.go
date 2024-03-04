package collections

import "testing"

func TestQueueCreation(t *testing.T) {
	q := NewQueue[string]()
	size := q.Size()
	exp := 0
	if size != exp {
		t.Errorf("expected size %d, got %d", exp, size)
	}

	b := q.IsEmpty()
	if !b {
		t.Errorf("expected empty to be %t, got %t", true, b)
	}

	b = q.HasItems()
	if b {
		t.Errorf("expected HasItems to be %t, got %t", false, b)
	}
}

func TestEnqueue(t *testing.T) {
	q := NewQueue[string]()
	q.Enqueue("foo")

	size := q.Size()
	exp := 1
	if size != exp {
		t.Errorf("expected size %d, got %d", exp, size)
	}

	q.Enqueue("bar")
	size = q.Size()
	exp = 2
	if size != exp {
		t.Errorf("expected size %d, got %d", exp, size)
	}

	b := q.IsEmpty()
	if b {
		t.Errorf("expected empty to be %t, got %t", false, b)
	}

	b = q.HasItems()
	if !b {
		t.Errorf("expected HasItems to be %t, got %t", true, b)
	}
}

func TestDequeue(t *testing.T) {
	q := NewQueue[string]()
	q.Enqueue("foo")
	q.Enqueue("bar")

	size := q.Size()
	exp := 2
	if size != exp {
		t.Errorf("expected size %d, got %d", exp, size)
	}

	item := q.Dequeue()
	expItem := "foo"
	if item != expItem {
		t.Errorf("expected item %s, got %s", expItem, item)
	}

	item = q.Dequeue()
	expItem = "bar"
	if item != expItem {
		t.Errorf("expected item %s, got %s", expItem, item)
	}

	b := q.IsEmpty()
	if !b {
		t.Errorf("expected empty to be %t, got %t", true, b)
	}

	b = q.HasItems()
	if b {
		t.Errorf("expected HasItems to be %t, got %t", false, b)
	}
}

func TestEnqueueAll(t *testing.T) {
	q := NewQueue[string]()
	items := []string{"foo", "bar", "baz"}
	q.EnqueueAll(items)

	size := q.Size()
	exp := 3
	if size != exp {
		t.Errorf("expected size %d, got %d", exp, size)
	}

	item := q.Dequeue()
	expItem := "foo"
	if item != expItem {
		t.Errorf("expected item %s, got %s", expItem, item)
	}

	item = q.Dequeue()
	expItem = "bar"
	if item != expItem {
		t.Errorf("expected item %s, got %s", expItem, item)
	}

	item = q.Dequeue()
	expItem = "baz"
	if item != expItem {
		t.Errorf("expected item %s, got %s", expItem, item)
	}

	b := q.IsEmpty()
	if !b {
		t.Errorf("expected empty to be %t, got %t", true, b)
	}

	b = q.HasItems()
	if b {
		t.Errorf("expected HasItems to be %t, got %t", false, b)
	}
}
