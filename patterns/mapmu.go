package patterns

import (
	"sync"
)

type (
	MapMu[T comparable, M any] struct {
		mu sync.RWMutex
		m  map[T]M
	}
)

func NewMapMu[T comparable, M any]() *MapMu[T, M] {
	return &MapMu[T, M]{
		m:  make(map[T]M),
		mu: sync.RWMutex{},
	}
}

func (mm *MapMu[T, M]) Clear() {
	mm.mu.Lock()
	defer mm.mu.Unlock()
	mm.m = make(map[T]M)
}

func (mm *MapMu[T, M]) CompareAndDelete(key T, value M) (deleted bool) {
	mm.mu.Lock()
	defer mm.mu.Unlock()

	v, ok := mm.m[key]
	if !ok {
		return
	}

	if any(v) != any(value) {
		return
	}

	delete(mm.m, key)
	deleted = true

	return
}

func (mm *MapMu[T, M]) CompareAndSwap(key T, old, new M) (swapped bool) {
	mm.mu.Lock()
	defer mm.mu.Unlock()

	v, ok := mm.m[key]
	if !ok {
		return
	}

	if any(v) != any(old) {
		return
	}

	mm.m[key] = new
	swapped = true

	return
}

func (mm *MapMu[T, M]) Delete(key T) {
	mm.mu.Lock()
	defer mm.mu.Unlock()

	delete(mm.m, key)
}

func (mm *MapMu[T, M]) Load(key T) (value M, ok bool) {
	mm.mu.RLock()
	defer mm.mu.RUnlock()

	v, ok := mm.m[key]
	return v, ok
}

func (mm *MapMu[T, M]) LoadOrStore(key T, value M) (actual M, loaded bool) {
	mm.mu.Lock()
	defer mm.mu.Unlock()

	v, ok := mm.m[key]
	if ok {
		return v, true
	}

	mm.m[key] = value
	return value, false
}

func (mm *MapMu[T, M]) LoadAndDelete(key T) (value M, loaded bool) {
	mm.mu.Lock()
	defer mm.mu.Unlock()

	v, ok := mm.m[key]
	if !ok {
		return
	}

	delete(mm.m, key)
	return v, true
}

func (mm *MapMu[T, M]) Range(f func(key T, value M) bool) {
	mm.mu.RLock()
	defer mm.mu.RUnlock()

	for k, v := range mm.m {
		if !f(k, v) {
			break
		}
	}
}

func (mm *MapMu[T, M]) Store(key T, value M) {
	mm.mu.Lock()
	defer mm.mu.Unlock()

	mm.m[key] = value
}

func (mm *MapMu[T, M]) Swap(key T, value M) (previous M, loaded bool) {
	mm.mu.Lock()
	defer mm.mu.Unlock()

	v, ok := mm.m[key]
	if !ok {
		return
	}

	mm.m[key] = value
	return v, true
}

func (mm *MapMu[T, M]) Map() map[T]M {
	mm.mu.RLock()
	defer mm.mu.RUnlock()

	m := make(map[T]M, len(mm.m))
	for k, v := range mm.m {
		m[k] = v
	}
	return m
}

func (mm *MapMu[T, M]) Len() int {
	mm.mu.RLock()
	defer mm.mu.RUnlock()

	return len(mm.m)
}
