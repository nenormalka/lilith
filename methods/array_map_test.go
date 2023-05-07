package methods

import (
	"reflect"
	"testing"
)

func TestArrayMap(t *testing.T) {
	for name, tt := range map[string]struct {
		args []int
		f    ArrayMapFunc[int]
		want []int
	}{
		"#1": {
			args: []int{5, 4, 3, 2, 1},
			f: func(elem int) int {
				return elem * elem
			},
			want: []int{25, 16, 9, 4, 1},
		},
		"#2": {
			args: []int{5, 4, 3, 2, 1},
			f: func(elem int) int {
				return elem * 2
			},
			want: []int{10, 8, 6, 4, 2},
		},
		"#3": {
			args: []int{},
			want: nil,
		},
	} {
		t.Run(name, func(t *testing.T) {
			if got := ArrayMap[[]int](tt.f, tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
