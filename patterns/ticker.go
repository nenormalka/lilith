package patterns

import (
	"context"
	"time"
)

type (
	TickerFunc func()
)

func Ticker(
	ctx context.Context,
	interval time.Duration,
	f TickerFunc,
) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
			case <-ticker.C:
				f()
			}
		}
	}()
}