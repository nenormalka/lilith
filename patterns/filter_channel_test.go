package patterns

import (
	"context"
	"reflect"
	"testing"
)

func TestFilterChannel(t *testing.T) {
	for name, tt := range map[string]struct {
		args []int
		want []int
		f    FilterFunc[int]
	}{
		"#1": {
			args: []int{1, 2, 3, 4},
			want: []int{1, 3},
			f: func(data int) bool {
				if data%2 == 0 {
					return false
				}

				return true
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			source := make(chan int)
			stream := FilterChannel(context.Background(), source, tt.f)

			go func() {
				for _, arg := range tt.args {
					source <- arg
				}
				close(source)
			}()

			var got []int

			for val := range stream {
				got = append(got, val)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterChannel() = %v, want %v", got, tt.want)
			}
		})
	}
}
