package patterns

import (
	"context"
	"fmt"
	"sync"
)

type (
	WorkFunc[T any] func(data T) error
)

func Workers[T any](
	ctx context.Context,
	f WorkFunc[T],
	n int,
	wg *sync.WaitGroup,
	source <-chan T,
	errCh chan error,
	panicCh chan struct{},
) {
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("worker panic %v", r)

					if panicCh != nil {
						panicCh <- struct{}{}
					}
				}

				wg.Done()
			}()

			for {
				select {
				case <-ctx.Done():
					return
				case data, ok := <-source:
					if !ok {
						return
					}

					if err := f(data); err != nil && errCh != nil {
						select {
						case errCh <- err:
						default:
						}
					}
				}
			}
		}()
	}
}
