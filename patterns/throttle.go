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

func Throttle[T any](e Effector[T], max, refill uint, d time.Duration) Effector[T] {
	var (
		tokens = max
		once   sync.Once
	)

	return func(ctx context.Context) (T, error) {
		var response T
		if ctx.Err() != nil {
			return response, ctx.Err()
		}

		once.Do(func() {
			ticker := time.NewTicker(d)
			go func() {
				defer ticker.Stop()
				for {
					select {
					case <-ctx.Done():
						return
					case <-ticker.C:
						t := tokens + refill
						if t > max {
							t = max
						}
						tokens = t
					}
				}
			}()
		})

		if tokens <= 0 {
			return response, ErrToManyCalls
		}

		tokens--

		return e(ctx)
	}
}
