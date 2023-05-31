package structures

import (
	"errors"

	"github.com/nenormalka/lilith/methods"
)

type (
	Item[T any] struct {
		Value    T
		Priority float64
	}

	Items[T any] []*Item[T]

	PriorityQueue[T any] struct {
		items Items[T]
	}
)

var (
	ErrEmptyQueue = errors.New("empty queue")
)

func NewPriorityQueue[T any](items Items[T]) *PriorityQueue[T] {
	pq := &PriorityQueue[T]{
		items: items,
	}

	pq.init()

	return pq
}

func (pq *PriorityQueue[T]) Push(x *Item[T]) {
	pq.items.push(x)
	pq.up(pq.items.len() - 1)
}

func (pq *PriorityQueue[T]) Pop() (*Item[T], error) {
	n := pq.items.len()
	if n == 0 {
		return nil, ErrEmptyQueue
	}

	n--
	pq.items.swap(0, n)
	pq.down(0, n)

	return pq.items.pop(), nil
}

func (pq *PriorityQueue[T]) Len() int {
	return pq.items.len()
}

func (pq *PriorityQueue[T]) init() {
	n := pq.items.len()

	for i := n/2 - 1; i >= 0; i-- {
		pq.down(i, n)
	}
}

func (pq *PriorityQueue[T]) up(j int) {
	for {
		i := (j - 1) / 2
		if i == j || !pq.items.less(j, i) {
			break
		}

		pq.items.swap(i, j)
		j = i
	}
}

func (pq *PriorityQueue[T]) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 {
			break
		}

		j := j1
		if j2 := j1 + 1; j2 < n && pq.items.less(j2, j1) {
			j = j2
		}

		if !pq.items.less(j, i) {
			break
		}

		pq.items.swap(i, j)

		i = j
	}

	return i > i0
}

func (items Items[T]) len() int {
	return len(items)
}

func (items Items[T]) less(i, j int) bool {
	return items[i].Priority > items[j].Priority
}

func (items Items[T]) swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

func (items *Items[T]) push(item *Item[T]) {
	methods.ArrayPush[Items[T]](items, item)
}

func (items *Items[T]) pop() *Item[T] {
	return methods.ArrayPop[Items[T]](items)
}
