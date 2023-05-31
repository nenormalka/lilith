package structures

import "testing"

func TestPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue[int](Items[int]{&Item[int]{Value: 5, Priority: 5}})

	pq.Push(&Item[int]{Value: 3, Priority: 7})
	pq.Push(&Item[int]{Value: 1, Priority: 10})
	pq.Push(&Item[int]{Value: 6, Priority: 3})

	item, err := pq.Pop()
	if err != nil {
		t.Error(err)
	}

	if item.Value != 1 {
		t.Errorf("Expected 1, got %d", item.Value)
	}

	item, err = pq.Pop()
	if err != nil {
		t.Error(err)
	}

	if item.Value != 3 {
		t.Errorf("Expected 3, got %d", item.Value)
	}

	pq.Push(&Item[int]{Value: 5, Priority: 5})

	if pq.Len() != 3 {
		t.Errorf("Expected 3, got %d", pq.Len())
	}

	item, err = pq.Pop()
	if err != nil {
		t.Error(err)
	}

	if item.Value != 5 {
		t.Errorf("Expected 5, got %d", item.Value)
	}

	item, err = pq.Pop()
	if err != nil {
		t.Error(err)
	}

	if item.Value != 5 {
		t.Errorf("Expected 5, got %d", item.Value)
	}

	item, err = pq.Pop()
	if err != nil {
		t.Error(err)
	}

	if item.Value != 6 {
		t.Errorf("Expected 6, got %d", item.Value)
	}

	if pq.Len() != 0 {
		t.Errorf("Expected 0, got %d", pq.Len())
	}

	item, err = pq.Pop()
	if err != ErrEmptyQueue {
		t.Errorf("Expected ErrEmptyQueue, got %v", err)
	}
}
