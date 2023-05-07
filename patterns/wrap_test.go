package patterns

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"
)

func TestWrap(t *testing.T) {
	err := errors.New("test")
	a := 0

	for name, tt := range map[string]struct {
		err error
		f   WrappedFunc
	}{
		"#1": {
			err: nil,
			f: func() error {
				a = 1
				return nil
			},
		},
		"#2": {
			err: err,
			f: func() error {
				return err
			},
		},
		"#3": {
			err: context.DeadlineExceeded,
			f: func() error {
				time.Sleep(2 * time.Second)
				return nil
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			errW := Wrap(ctx, tt.f)
			cancel()

			if tt.err != nil {
				if !reflect.DeepEqual(errW, tt.err) {
					t.Errorf("Wrap() = %v, want %v", errW, tt.err)
				}
			} else {
				if a != 1 {
					t.Errorf("Breaker() = %v, want %v", a, 1)
				}

				if errW != nil {
					t.Errorf("Breaker() = %v, want %v", errW, nil)
				}
			}
		})
	}
}
