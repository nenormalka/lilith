package patterns

import (
	"context"
	"errors"
)

type (
	Semaphore struct {
		c chan struct{}
	}
)

var (
	ErrInvalidCount = errors.New("invalid count")
)

func (s *Semaphore) AcquireCtx(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case s.c <- struct{}{}:
		return nil
	}
}

func (s *Semaphore) Acquire() {
	s.c <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.c
}

func NewSemaphore(n int) (*Semaphore, error) {
	if n < 1 {
		return nil, ErrInvalidCount
	}

	return &Semaphore{c: make(chan struct{}, n)}, nil
}
