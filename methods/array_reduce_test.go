package methods

import (
	"reflect"
	"testing"
)

func TestArrayReduce(t *testing.T) {
	for name, tt := range map[string]struct {
		args []int
		f    ArrayReduceFunc[int]
		init int
		want int
	}{
		"#1": {
			args: []int{5, 4, 3, 2, 1},
			f: func(first, second int) int {
				return first + second
			},
			init: 0,
			want: 15,
		},
		"#2": {
			args: []int{},
			f: func(first, second int) int {
				return first + second
			},
			init: 1,
			want: 1,
		},
	} {
		t.Run(name, func(t *testing.T) {
			if got := ArrayReduce[[]int](tt.f, tt.args, tt.init); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
