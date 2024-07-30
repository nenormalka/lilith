package patterns

import (
	"context"
	"sync"
)

type (
	WorkFunc[T any] func(data T)
)

func Workers[T any](
	ctx context.Context,
	f WorkFunc[T],
	n uint,
	source <-chan T,
) <-chan struct{} {
	wg := &sync.WaitGroup{}
	wg.Add(int(n))

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	for i := 0; i < int(n); i++ {
		go func() {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case data, ok := <-source:
					if !ok {
						return
					}

					f(data)
				}
			}
		}()
	}

	return done
}
