package methods

import (
	"reflect"
	"testing"
)

func TestArrayFlip(t *testing.T) {
	for name, tt := range map[string]struct {
		args []int
		want map[int]int
	}{
		"#1": {
			args: []int{5, 4, 3, 2, 1},
			want: map[int]int{
				1: 4,
				2: 3,
				3: 2,
				4: 1,
				5: 0,
			},
		},
		"#2": {
			args: []int{5, 4, 3, 3, 3},
			want: map[int]int{
				3: 4,
				4: 1,
				5: 0,
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			if got := ArrayFlip[[]int](tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayFlip() = %v, want %v", got, tt.want)
			}
		})
	}
}
