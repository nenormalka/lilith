package methods

import (
	"testing"
)

func TestArrayIntersect(t *testing.T) {
	for name, tt := range map[string]struct {
		first  []int
		second []int
		want   []int
	}{
		"#1": {
			first:  []int{5, 4, 3, 2, 1},
			second: []int{},
			want:   nil,
		},
		"#2": {
			first:  []int{},
			second: []int{5, 4, 3, 2, 1},
			want:   nil,
		},
		"#3": {
			first:  nil,
			second: nil,
			want:   nil,
		},
		"#4": {
			first:  []int{5, 4, 3, 2, 1},
			second: []int{5, 4, 3, 1},
			want:   []int{5, 4, 3, 1},
		},
		"#5": {
			first:  []int{5, 4, 3, 1},
			second: []int{5, 4, 3, 2, 1},
			want:   []int{5, 4, 3, 1},
		},
		"#6": {
			first:  []int{5, 4, 3, 1, 0},
			second: []int{5, 4, 3, 2, 1},
			want:   []int{5, 4, 3, 1},
		},
	} {
		t.Run(name, func(t *testing.T) {
			got := ArrayIntersect[[]int](tt.first, tt.second)

			if len(got) != len(tt.want) {
				t.Errorf("size ArrayIntersect() = %v, want %v", got, tt.want)
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
