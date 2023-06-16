package patterns

import (
	"sync"
	"time"
)

type (
	ReloadFunc[K comparable, T any] func(oldCache map[K]T) map[K]T

	CacheMap[K comparable, T any] struct {
		stopCh chan struct{}
		mu     sync.RWMutex
		cache  map[K]T
	}
)

func NewCacheMap[K comparable, T any](f ReloadFunc[K, T], clearTick time.Duration) *CacheMap[K, T] {
	cm := &CacheMap[K, T]{
		mu:     sync.RWMutex{},
		cache:  map[K]T{},
		stopCh: make(chan struct{}),
	}

	cm.reload(f, clearTick)

	return cm
}

func (cm *CacheMap[K, T]) Get(key K) (T, bool) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	var value T

	i, ok := cm.cache[key]

	if ok {
		value = i
	}

	return value, ok
}

func (cm *CacheMap[K, T]) Set(key K, value T) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	cm.cache[key] = value
}

func (cm *CacheMap[K, T]) ClearCache() {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	cm.cache = map[K]T{}
}

func (cm *CacheMap[K, T]) RemoveCache() {
	cm.ClearCache()

	close(cm.stopCh)
}

func (cm *CacheMap[K, T]) reload(f ReloadFunc[K, T], clearTick time.Duration) {
	go func() {
		ticker := time.NewTicker(clearTick)
		defer ticker.Stop()

		for {
			cm.mu.Lock()
			cm.cache = f(cm.cache)
			cm.mu.Unlock()

			select {
			case <-cm.stopCh:
				return
			case <-ticker.C:
			}
		}
	}()
}
