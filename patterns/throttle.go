package patterns

import (
	"context"
	"errors"
	"sync"
	"time"
)

var (
	ErrToManyCalls = errors.New("too many calls")
)

func Throttle[T any](
	ctx context.Context,
	e Effector[T],
	max, refill uint,
	d time.Duration,
) Effector[T] {
	var (
		tokens = max
		mu     sync.Mutex
	)

	go func() {
		ticker := time.NewTicker(d)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				mu.Lock()

				t := tokens + refill
				if t > max {
					t = max
				}
				tokens = t

				mu.Unlock()
			}
		}
	}()

	return func(ctx context.Context) (T, error) {
		var response T
		if ctx.Err() != nil {
			return response, ctx.Err()
		}

		mu.Lock()
		defer mu.Unlock()

		if tokens <= 0 {
			return response, ErrToManyCalls
		}

		tokens--

		return e(ctx)
	}
}
