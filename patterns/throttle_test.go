package patterns

import (
	"context"
	"reflect"
	"testing"
	"time"
)

func TestThrottle(t *testing.T) {
	type test struct {
		res int
		err error
	}
	for name, tt := range map[string]struct {
		e      Effector[int]
		max    uint
		refill uint
		d      time.Duration
		want   []test
		ctx    context.Context
		times  int
		sleep  time.Duration
	}{
		"#1": {
			e: func(ctx context.Context) (int, error) {
				return 1, nil
			},
			max:    1,
			refill: 1,
			d:      time.Second,
			want: []test{
				{
					res: 1,
					err: nil,
				},
				{
					res: 0,
					err: ErrToManyCalls,
				},
				{
					res: 0,
					err: ErrToManyCalls,
				},
			},
			ctx:   context.Background(),
			times: 3,
			sleep: 0,
		},
		"#2": {
			e: func(ctx context.Context) (int, error) {
				return 1, nil
			},
			max:    1,
			refill: 1,
			d:      500 * time.Millisecond,
			want: []test{
				{
					res: 1,
					err: nil,
				},
				{
					res: 1,
					err: nil,
				},
			},
			ctx:   context.Background(),
			times: 2,
			sleep: time.Second,
		},
	} {
		t.Run(name, func(t *testing.T) {
			th := Throttle(tt.ctx, tt.e, tt.max, tt.refill, tt.d)
			for i := 0; i < tt.times; i++ {
				got, err := th(tt.ctx)
				if !reflect.DeepEqual(got, tt.want[i].res) {
					t.Errorf("Throttle() = %v, want %v", got, tt.want[i].res)
				}
				if !reflect.DeepEqual(err, tt.want[i].err) {
					t.Errorf("Throttle() = %v, want %v", err, tt.want[i].err)
				}
				if tt.sleep != 0 {
					time.Sleep(tt.sleep)
				}
			}
		})
	}
}
