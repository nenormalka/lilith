package patterns

import (
	"context"
	"sync"
)

type (
	WorkersCurator[T any] struct {
		f       WorkFunc[T]
		wg      *sync.WaitGroup
		ctx     context.Context
		source  <-chan T
		errCh   chan error
		panicCh chan struct{}
	}
)

func NewWorkersCurator[T any](
	ctx context.Context,
	f WorkFunc[T],
	source <-chan T,
	workerCount int,
) *WorkersCurator[T] {
	c := &WorkersCurator[T]{
		f:       f,
		wg:      &sync.WaitGroup{},
		ctx:     ctx,
		source:  source,
		errCh:   make(chan error),
		panicCh: make(chan struct{}),
	}

	Workers[T](c.ctx, c.f, workerCount, c.wg, c.source, c.errCh, c.panicCh)

	c.process()

	return c
}

func (c *WorkersCurator[T]) process() {
	go func() {
		for {
			select {
			case _, ok := <-c.panicCh:
				if !ok {
					return
				}

				Workers[T](c.ctx, c.f, 1, c.wg, c.source, c.errCh, c.panicCh)
			}
		}
	}()
}

func (c *WorkersCurator[T]) Wait() {
	c.wg.Wait()
	close(c.panicCh)
	close(c.errCh)
}

func (c *WorkersCurator[T]) GetErrCh() chan error {
	return c.errCh
}
