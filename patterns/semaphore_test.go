package patterns

import (
	"context"
	"testing"
	"time"
)

func TestNewSemaphore(t *testing.T) {
	s, err := NewSemaphore(0)
	if err != ErrInvalidCount {
		t.Errorf("expected ErrInvalidCount, got %v", err)
	}

	s, err = NewSemaphore(1)
	s.Acquire()

	ctx, _ := context.WithTimeout(context.Background(), time.Second)

	if err = s.AcquireCtx(ctx); err != context.DeadlineExceeded {
		t.Errorf("expected context.DeadlineExceeded, got %v", err)
	}

	s.Release()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err = s.AcquireCtx(ctx); err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	s.Release()
}
