package patterns

import (
	"context"
	"errors"
)

type (
	Semaphore struct {
		C chan struct{}
	}
)

var (
	ErrInvalidCount = errors.New("invalid count")
)

func (s *Semaphore) AcquireCtx(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case s.C <- struct{}{}:
		return nil
	}
}

func (s *Semaphore) Acquire() {
	s.C <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.C
}

func NewSemaphore(n int) (*Semaphore, error) {
	if n < 1 {
		return nil, ErrInvalidCount
	}

	return &Semaphore{C: make(chan struct{}, n)}, nil
}
