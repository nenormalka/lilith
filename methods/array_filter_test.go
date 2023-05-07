package methods

import (
	"reflect"
	"testing"
)

func TestArrayFilter(t *testing.T) {
	for name, tt := range map[string]struct {
		args []int
		f    ArrayFilterFunc[int]
		want []int
	}{
		"#1": {
			args: []int{5, 4, 3, 2, 1},
			f: func(elem int) bool {
				return elem%2 == 0
			},
			want: []int{4, 2},
		},
		"#2": {
			args: []int{5, 4, 3, 2, 1},
			f: func(elem int) bool {
				return elem%2 == 1
			},
			want: []int{5, 3, 1},
		},
		"#3": {
			args: []int{},
			want: nil,
		},
	} {
		t.Run(name, func(t *testing.T) {
			if got := ArrayFilter[[]int](tt.f, tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}
