package methods

import (
	"testing"
)

func TestArrayDiff(t *testing.T) {
	for name, tt := range map[string]struct {
		first  []int
		second []int
		want   []int
	}{
		"#1": {
			first:  []int{5, 4, 3, 2, 1},
			second: []int{},
			want:   []int{5, 4, 3, 2, 1},
		},
		"#2": {
			first:  []int{},
			second: []int{5, 4, 3, 2, 1},
			want:   []int{5, 4, 3, 2, 1},
		},
		"#3": {
			first:  nil,
			second: nil,
			want:   nil,
		},
		"#4": {
			first:  []int{5, 4, 3, 2, 1},
			second: []int{5, 4, 3, 1},
			want:   []int{2},
		},
		"#5": {
			first:  []int{5, 4, 3, 1},
			second: []int{5, 4, 3, 2, 1},
			want:   []int{2},
		},
		"#6": {
			first:  []int{5, 4, 3, 1, 0},
			second: []int{5, 4, 3, 2, 1},
			want:   []int{0, 2},
		},
	} {
		t.Run(name, func(t *testing.T) {
			got := ArrayDiff[[]int](tt.first, tt.second)

			if len(got) != len(tt.want) {
				t.Errorf("size ArrayDiff() = %v, want %v", got, tt.want)
			}

			l := len(got)

			for _, val1 := range got {
				for _, val2 := range tt.want {
					if val1 != val2 {
						continue
					}

					l--
					break
				}
			}

			if l != 0 {
				t.Errorf("invalid values %v", tt.want)
			}
		})
	}
}
