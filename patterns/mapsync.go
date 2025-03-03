package patterns

import (
	"sync"
	"sync/atomic"
)

type (
	MapSync[T comparable, M any] struct {
		l atomic.Int64
		m sync.Map
	}
)

func NewMapSync[T comparable, M any]() *MapSync[T, M] {
	return &MapSync[T, M]{
		m: sync.Map{},
		l: atomic.Int64{},
	}
}

func (ms *MapSync[T, M]) Clear() {
	ms.m.Clear()
	ms.l.Store(0)
}

func (ms *MapSync[T, M]) CompareAndDelete(key T, value M) (deleted bool) {
	deleted = ms.m.CompareAndDelete(key, value)
	if deleted {
		ms.l.Add(-1)
	}
	return
}

func (ms *MapSync[T, M]) CompareAndSwap(key T, old, new M) (swapped bool) {
	return ms.m.CompareAndSwap(key, old, new)
}

func (ms *MapSync[T, M]) Delete(key T) {
	ms.m.Delete(key)
	ms.l.Add(-1)
}

func (ms *MapSync[T, M]) Load(key T) (value M, ok bool) {
	v, ok := ms.m.Load(key)
	if !ok {
		return
	}

	value, ok = v.(M)
	return
}

func (ms *MapSync[T, M]) LoadOrStore(key T, value M) (actual M, loaded bool) {
	v, loaded := ms.m.LoadOrStore(key, value)
	actual = v.(M)
	if !loaded {
		ms.l.Add(1)
	}
	return
}

func (ms *MapSync[T, M]) LoadAndDelete(key T) (value M, loaded bool) {
	v, loaded := ms.m.LoadAndDelete(key)
	value = v.(M)
	if loaded {
		ms.l.Add(-1)
	}
	return
}

func (ms *MapSync[T, M]) Range(f func(key T, value M) bool) {
	ms.m.Range(func(key, value any) bool {
		return f(key.(T), value.(M))
	})
}

func (ms *MapSync[T, M]) Store(key T, value M) {
	ms.m.Store(key, value)
	ms.l.Add(1)
}

func (ms *MapSync[T, M]) Swap(key T, value M) (previous M, loaded bool) {
	v, loaded := ms.m.Swap(key, value)
	previous = v.(M)
	return
}

func (ms *MapSync[T, M]) Map() map[T]M {
	m := make(map[T]M, ms.Len())
	ms.Range(func(key T, value M) bool {
		m[key] = value
		return true
	})
	return m
}

func (ms *MapSync[T, M]) Len() int {
	return int(ms.l.Load())
}
