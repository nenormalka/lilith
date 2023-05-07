package patterns

import (
	"context"
	"sync"
	"time"
)

func Debounce[T any](circuit Circuit[T], d time.Duration) Circuit[T] {
	var (
		threshold time.Time
		m         sync.Mutex
		result    T
	)

	return func(ctx context.Context) (T, error) {
		m.Lock()
		defer func() {
			threshold = time.Now().Add(d)
			m.Unlock()
		}()

		if time.Now().Before(threshold) {
			return result, ErrCircuitWait
		}

		return circuit(ctx)
	}
}
