package patterns

import (
	"context"
	"reflect"
	"testing"
)

func TestAggregateChannel(t *testing.T) {
	for name, tt := range map[string]struct {
		args []int
		want []int
		f    AggregateFunc[int]
	}{
		"#1": {
			args: []int{1, 2, 3, 4},
			want: []int{2, 4, 6, 8},
			f: func(data int) int {
				return data * 2
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			source := make(chan int)
			stream := AggregateChannel(context.Background(), source, tt.f)

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
				t.Errorf("AggregateChannel() = %v, want %v", got, tt.want)
			}
		})
	}
}
