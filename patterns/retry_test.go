package patterns

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"
)

func TestRetry(t *testing.T) {
	err := errors.New("test")
	count := 0

	for name, tt := range map[string]struct {
		effector Effector[int]
		retries  int
		delay    time.Duration
		ctx      context.Context
		want     int
		err      error
	}{
		"#1": {
			effector: func(ctx context.Context) (int, error) {
				count++
				if count == 2 {
					return 1, nil
				}

				return 0, err
			},
			retries: 3,
			delay:   0,
			ctx:     context.Background(),
			want:    1,
			err:     nil,
		},
		"#2": {
			effector: func(ctx context.Context) (int, error) {
				return 0, err
			},
			retries: 3,
			delay:   0,
			ctx:     context.Background(),
			want:    0,
			err:     err,
		},
		"#3": {
			effector: func(ctx context.Context) (int, error) {
				return 0, err
			},
			retries: 3,
			delay:   time.Second,
			ctx: func() context.Context {
				ctx, _ := context.WithTimeout(context.Background(), time.Second)
				return ctx
			}(),
			want: 0,
			err:  context.DeadlineExceeded,
		},
	} {
		t.Run(name, func(t *testing.T) {
			r := Retry(tt.effector, tt.retries, tt.delay)

			got, err := r(tt.ctx)
			count = 0

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Retry() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("Retry() = %v, want %v", got, tt.want)
			}
		})
	}
}
