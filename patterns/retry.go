package patterns

import (
	"context"
	"log"
	"time"
)

type (
	Effector[T any] func(ctx context.Context) (T, error)
)

func Retry[T any](effector Effector[T], retries int, delay time.Duration) Effector[T] {
	return func(ctx context.Context) (T, error) {
		for r := 0; ; r++ {
			response, err := effector(ctx)
			if err == nil || r >= retries {
				return response, err
			}

			log.Printf("Attempt %d failed; retrying in %v", r+1, delay)

			select {
			case <-time.After(delay):
			case <-ctx.Done():
				return response, ctx.Err()
			}
		}
	}
}
