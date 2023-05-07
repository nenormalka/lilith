package methods

import (
	"testing"
)

func TestArrayToMapValues(t *testing.T) {
	for name, tt := range map[string]struct {
		args []int
		want map[int]struct{}
	}{
		"#1": {
			args: []int{5, 4, 3, 2, 1},
			want: map[int]struct{}{
				5: {},
				4: {},
				3: {},
				2: {},
				1: {},
			},
		},
		"#2": {
			args: nil,
			want: nil,
		},
	} {
		t.Run(name, func(t *testing.T) {
			got := ArrayToMapValues[[]int](tt.args)

			if len(got) != len(tt.want) {
				t.Errorf("size ArrayToMapValues() = %v, want %v", got, tt.want)
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
