package methods

import (
	"testing"
)

func TestArrayToMapValuesBool(t *testing.T) {
	for name, tt := range map[string]struct {
		args []int
		want map[int]bool
	}{
		"#1": {
			args: []int{5, 4, 3, 2, 1},
			want: map[int]bool{
				5: true,
				4: true,
				3: true,
				2: true,
				1: true,
			},
		},
		"#2": {
			args: nil,
			want: nil,
		},
	} {
		t.Run(name, func(t *testing.T) {
			got := ArrayToMapValuesBool[[]int](tt.args)

			if len(got) != len(tt.want) {
				t.Errorf("size ArrayToMapValuesBool() = %v, want %v", got, tt.want)
			}

			for key := range got {
				if _, ok := tt.want[key]; !ok {
					t.Errorf("not found key %v", key)
				}

				delete(tt.want, key)
			}

			if len(tt.want) != 0 {
				t.Errorf("invalid keys %v", tt.want)
			}
		})
	}
}
