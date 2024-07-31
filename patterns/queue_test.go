package patterns

import (
	"errors"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	q := NewQueue[string, string]()
	if q == nil {
		t.Fatal("expected queue to be created")
	}

	if q.Size() != 0 {
		t.Fatalf("expected size to be 0, got %d", q.Size())
	}

	q.Push(1, "key1", "value1")
	q.Push(10, "key2", "value2")
	q.Push(0, "key3", "value3")
	q.Push(MaxPriority, "key4", "value4")
	q.Push(MinPriority, "key5", "value5")

	n, err := q.Pop()
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if n.Key != "key4" {
		t.Fatalf("expected key to be key4, got %s", n.Key)
	}

	if n.Value != "value4" {
		t.Fatalf("expected value to be value4, got %s", n.Value)
	}

	n, err = q.PopByKey("key2")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if n.Key != "key2" {
		t.Fatalf("expected key to be key2, got %s", n.Key)
	}

	if n.Value != "value2" {
		t.Fatalf("expected value to be value2, got %s", n.Value)
	}

	if q.Size() != 3 {
		t.Fatalf("expected size to be 3, got %d", q.Size())
	}

	q.Push(0, "key3", "value3")
	q.Push(100, "key3", "value3")
	q.Push(100, "key3", "value33")

	if q.Size() != 3 {
		t.Fatalf("expected size to be 3, got %d", q.Size())
	}

	n, err = q.Pop()
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if n.Key != "key3" {
		t.Fatalf("expected key to be key3, got %s", n.Key)
	}

	if n.Value != "value33" {
		t.Fatalf("expected value to be value3, got %s", n.Value)
	}

	n, err = q.Pop()
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if n.Key != "key1" {
		t.Fatalf("expected key to be key1, got %s", n.Key)
	}

	if n.Value != "value1" {
		t.Fatalf("expected value to be value1, got %s", n.Value)
	}

	n, err = q.Pop()
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if n.Key != "key5" {
		t.Fatalf("expected key to be key5, got %s", n.Key)
	}

	if n.Value != "value5" {
		t.Fatalf("expected value to be value5, got %s", n.Value)
	}

	if q.Size() != 0 {
		t.Fatalf("expected size to be 0, got %d", q.Size())
	}

	_, err = q.Pop()
	if !errors.Is(err, ErrEmptyQueue) {
		t.Fatalf("expected error to be ErrEmptyQueue, got %v", err)
	}

	_, err = q.PopByKey("key1")
	if !errors.Is(err, ErrKeyDoesNotExist) {
		t.Fatalf("expected error to be ErrEmptyQueue, got %v", err)
	}

	_, err = q.PopByChannel(500 * time.Millisecond)
	if !errors.Is(err, ErrEmptyChannel) {
		t.Fatalf("expected error to be ErrEmptyChannel, got %v", err)
	}

	q.Close()
}

func TestQueueWithChannel(t *testing.T) {
	q := NewQueue[string, string](WithChannel[string, string]())
	if q == nil {
		t.Fatal("expected queue to be created")
	}

	if q.ch == nil {
		t.Fatal("expected channel to be created")
	}

	if q.closeCh == nil {
		t.Fatal("expected close channel to be created")
	}

	if q.Size() != 0 {
		t.Fatalf("expected size to be 0, got %d", q.Size())
	}

	q.Push(1, "key1", "value1")
	q.Push(2, "key2", "value2")

	ch, err := q.PopByChannel(500 * time.Millisecond)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	n := <-ch
	if n.Key != "key2" {
		t.Fatalf("expected key to be key1, got %s", n.Key)
	}

	if n.Value != "value2" {
		t.Fatalf("expected value to be value1, got %s", n.Value)
	}

	n = <-ch
	if n.Key != "key1" {
		t.Fatalf("expected key to be key2, got %s", n.Key)
	}

	if n.Value != "value1" {
		t.Fatalf("expected value to be value2, got %s", n.Value)
	}

	go func() {
		n = <-ch
		if n.Key != "key3" {
			t.Fatalf("expected key to be key3, got %s", n.Key)
		}

		if n.Value != "value3" {
			t.Fatalf("expected value to be value3, got %s", n.Value)
		}
	}()

	q.Push(3, "key3", "value3")

	time.Sleep(time.Second)

	go func() {
		n = <-ch
		if n.Key != "key4" {
			t.Fatalf("expected key to be key4, got %s", n.Key)
		}

		if n.Value != "value4" {
			t.Fatalf("expected value to be value4, got %s", n.Value)
		}
	}()

	q.Push(4, "key4", "value4")

	time.Sleep(time.Second)

	q.Close()
}
