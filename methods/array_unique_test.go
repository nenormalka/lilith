package methods

import (
	"reflect"
	"testing"
)

func TestArrayUnique(t *testing.T) {
	for name, tt := range map[string]struct {
		args []int
		want []int
	}{
		"#1": {
			args: []int{5, 4, 3, 2, 1},
			want: []int{5, 4, 3, 2, 1},
		},
		"#2": {
			args: nil,
			want: nil,
		},
		"#3": {
			args: []int{5, 5, 3, 4, 7, 8, 3},
			want: []int{5, 3, 4, 7, 8},
		},
	} {
		t.Run(name, func(t *testing.T) {
			if got := ArrayUnique[[]int](tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestArrayUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
