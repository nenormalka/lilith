package patterns

import (
	"context"
	"errors"
	"time"
)

type (
	detachedContext struct{ parent context.Context }
)

var (
	ErrEmptyParentContext = errors.New("empty parent context")
)

func Detach(ctx context.Context) (context.Context, error) {
	if ctx == nil {
		return nil, ErrEmptyParentContext
	}

	return detachedContext{ctx}, nil
}

func (v detachedContext) Deadline() (time.Time, bool)       { return time.Time{}, false }
func (v detachedContext) Done() <-chan struct{}             { return nil }
func (v detachedContext) Err() error                        { return nil }
func (v detachedContext) Value(key interface{}) interface{} { return v.parent.Value(key) }
