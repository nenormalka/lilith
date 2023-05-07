package patterns

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"
)

func TestBreaker(t *testing.T) {
	type test struct {
		resp int
		err  error
	}

	err := errors.New("test")

	for name, tt := range map[string]struct {
		waitDuration     time.Duration
		failureThreshold uint
		timesExec        int
		want             []test
		ctxs             []context.Context
		circuit          Circuit[int]
	}{
		"#1": {
			waitDuration:     time.Second,
			failureThreshold: 1,
			timesExec:        3,
			want: []test{
				{
					resp: 1,
					err:  nil,
				},
				{
					resp: 1,
					err:  nil,
				},
				{
					resp: 1,
					err:  nil,
				},
			},
			ctxs: []context.Context{
				context.Background(),
				context.Background(),
				context.Background(),
			},
			circuit: func(ctx context.Context) (int, error) {
				return 1, nil
			},
		},
		"#2": {
			waitDuration:     time.Second,
			failureThreshold: 1,
			timesExec:        3,
			want: []test{
				{
					resp: 1,
					err:  nil,
				},
				{
					resp: 0,
					err:  err,
				},
				{
					resp: 0,
					err:  ErrCircuitWait,
				},
			},
			ctxs: []context.Context{
				context.Background(),
				context.WithValue(context.Background(), "err", err),
				context.Background(),
			},
			circuit: func(ctx context.Context) (int, error) {
				err, ok := ctx.Value("err").(error)
				if ok {
					return 0, err
				}

				return 1, nil
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			cb := Breaker(tt.circuit, tt.waitDuration, tt.failureThreshold)

			for i := 0; i < tt.timesExec; i++ {
				got, err := cb(tt.ctxs[i])
				if !reflect.DeepEqual(got, tt.want[i].resp) {
					t.Errorf("Breaker() = %v, want %v", got, tt.want[i].resp)
				}
				if !reflect.DeepEqual(err, tt.want[i].err) {
					t.Errorf("Breaker() = %v, want %v", err, tt.want[i].err)
				}
			}
		})
	}
}
