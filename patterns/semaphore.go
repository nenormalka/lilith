package patterns

import (
	"context"
	"errors"
)

type (
	Semaphore struct {
		с chan struct{}
	}
)

var (
	ErrInvalidCount = errors.New("invalid count")
)

func (s *Semaphore) AcquireCtx(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case s.с <- struct{}{}:
		return nil
	}
}

func (s *Semaphore) Acquire() {
	s.с <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.с
}

func NewSemaphore(n int) (*Semaphore, error) {
	if n < 1 {
		return nil, ErrInvalidCount
	}

	return &Semaphore{с: make(chan struct{}, n)}, nil
}
