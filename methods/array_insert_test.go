package methods

import (
	"reflect"
	"testing"
)

func TestArrayInsert(t *testing.T) {
	for name, tt := range map[string]struct {
		args  []int
		indx  int
		elems []int
		want  []int
	}{
		"#1": {
			args:  []int{5, 4, 3, 2, 1},
			indx:  2,
			elems: []int{1},
			want:  []int{5, 4, 1, 3, 2, 1},
		},
		"#2": {
			args:  []int{5},
			indx:  2,
			elems: []int{1},
			want:  []int{5, 1},
		},
		"#3": {
			args:  []int{5},
			indx:  2,
			elems: nil,
			want:  []int{5},
		},
		"#4": {
			args:  []int{5, 4, 3, 2, 1},
			indx:  0,
			elems: []int{1, 0},
			want:  []int{1, 0, 5, 4, 3, 2, 1},
		},
		"#5": {
			args:  nil,
			indx:  0,
			elems: []int{1},
			want:  []int{1},
		},
	} {
		t.Run(name, func(t *testing.T) {
			ArrayInsert[[]int](&tt.args, tt.indx, tt.elems...)

			if !reflect.DeepEqual(tt.args, tt.want) {
				t.Errorf("ArrayInsert() = %v, want %v", tt.args, tt.want)
			}
		})
	}
}
