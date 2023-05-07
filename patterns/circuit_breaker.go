package patterns

import (
	"context"
	"errors"
	"sync"
	"time"
)

type (
	Circuit[T any] func(ctx context.Context) (T, error)
)

var (
	ErrCircuitWait = errors.New("err need more time wait before request")
)

func Breaker[T any](circuit Circuit[T], waitDuration time.Duration, failureThreshold uint) Circuit[T] {
	var (
		consecutiveFailures = 0
		lastAttempt         = time.Now()
		m                   sync.RWMutex
	)

	return func(ctx context.Context) (T, error) {
		var response T

		m.RLock()

		if consecutiveFailures-int(failureThreshold) >= 0 {
			shouldRetryAt := lastAttempt.Add(waitDuration)
			if !time.Now().After(shouldRetryAt) {
				m.RUnlock()
				return response, ErrCircuitWait
			}
		}

		m.RUnlock()

		select {
		case <-ctx.Done():
			return response, ctx.Err()
		default:
		}

		response, err := circuit(ctx)
		m.Lock()
		defer m.Unlock()

		lastAttempt = time.Now()

		if err != nil {
			consecutiveFailures++
			return response, err
		}

		consecutiveFailures = 0

		return response, nil
	}
}
