package patterns

import (
	"context"
	"reflect"
	"testing"
	"time"
)

func TestDebounce(t *testing.T) {
	type test struct {
		res int
		err error
	}

	for name, tt := range map[string]struct {
		times int
		c     Circuit[int]
		want  []test
		d     time.Duration
	}{
		"#1": {
			times: 2,
			c: func(ctx context.Context) (int, error) {
				return 1, nil
			},
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
			d: 0,
		},
		"#2": {
			times: 2,
			c: func(ctx context.Context) (int, error) {
				return 1, nil
			},
			want: []test{
				{
					res: 1,
					err: nil,
				},
				{
					res: 0,
					err: ErrCircuitWait,
				},
			},
			d: time.Second,
		},
	} {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			deb := Debounce[int](tt.c, tt.d)
			for i := 0; i < tt.times; i++ {
				got, err := deb(ctx)
				if !reflect.DeepEqual(got, tt.want[i].res) {
					t.Errorf("Debounce() = %v, want %v", got, tt.want[i].res)
				}
				if !reflect.DeepEqual(err, tt.want[i].err) {
					t.Errorf("Debounce() = %v, want %v", err, tt.want[i].err)
				}
			}
		})
	}
}
