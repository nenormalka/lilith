package methods

import (
	"reflect"
	"testing"
)

func TestArrayMax(t *testing.T) {
	for name, tt := range map[string]struct {
		args []int
		want int
	}{
		"#1": {
			args: []int{5, 4, 3, 2, 1},
			want: 5,
		},
		"#2": {
			args: []int{},
			want: 0,
		},
		"#3": {
			args: []int{1, 2, 3, 4, 5},
			want: 5,
		},
	} {
		t.Run(name, func(t *testing.T) {
			if got := ArrayMax[[]int](tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
