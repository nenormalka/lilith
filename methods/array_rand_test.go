package methods

import (
	"testing"
)

func TestArrayRandInt(t *testing.T) {
	for name, tt := range map[string]struct {
		args []int
		want map[int]struct{}
	}{
		"#1": {
			args: nil,
			want: nil,
		},
		"#2": {
			args: []int{1, 44, -123, 88, 99},
			want: map[int]struct{}{
				1:    {},
				44:   {},
				-123: {},
				88:   {},
				99:   {},
			},
		},
		"#3": {
			args: []int{1},
			want: map[int]struct{}{
				1: {},
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			got := ArrayRand[[]int](tt.args)

			if len(got) != len(tt.want) {
				t.Errorf("size ArrayRand() = %v, want %v", got, tt.want)
			}

			sim := 0

			for inx := range got {
				if got[inx] == tt.args[inx] {
					sim++
				}
			}

			if len(tt.args) == sim && len(tt.args) > 1 {
				t.Errorf("similar array ArrayRand %v, want %v", got, tt.want)
			}

			for _, value := range got {
				_, ok := tt.want[value]
				if !ok {
					t.Errorf("not found key %v", value)
				}

				delete(tt.want, value)
			}

			if len(tt.want) != 0 {
				t.Errorf("invalid keys %v", tt.want)
			}
		})
	}
}

func TestArrayRandStr(t *testing.T) {
	for name, tt := range map[string]struct {
		args []string
		want map[string]struct{}
	}{
		"#1": {
			args: nil,
			want: nil,
		},
		"#2": {
			args: []string{"t", "e", "s", "m"},
			want: map[string]struct{}{
				"t": {},
				"e": {},
				"s": {},
				"m": {},
			},
		},
		"#3": {
			args: []string{"t"},
			want: map[string]struct{}{
				"t": {},
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			got := ArrayRand[[]string](tt.args)

			if len(got) != len(tt.want) {
				t.Errorf("size ArrayRand() = %v, want %v", got, tt.want)
			}

			sim := 0

			for inx := range got {
				if got[inx] == tt.args[inx] {
					sim++
				}
			}

			if len(tt.args) == sim && len(tt.args) > 1 {
				t.Errorf("similar array ArrayRand %v, want %v", got, tt.args)
			}

			for _, value := range got {
				_, ok := tt.want[value]
				if !ok {
					t.Errorf("not found key %v", value)
				}

				delete(tt.want, value)
			}

			if len(tt.want) != 0 {
				t.Errorf("invalid keys %v", tt.want)
			}
		})
	}
}

func TestArrayRandStruct(t *testing.T) {
	type test struct {
		v int
		s string
	}
	for name, tt := range map[string]struct {
		args []test
		want map[test]struct{}
	}{
		"#1": {
			args: nil,
			want: nil,
		},
		"#2": {
			args: []test{
				{
					v: 1,
					s: "t",
				},
				{
					v: 2,
					s: "e",
				},
				{
					v: 3,
					s: "s",
				},
				{
					v: 4,
					s: "m",
				},
			},
			want: map[test]struct{}{
				{
					v: 1,
					s: "t",
				}: {},
				{
					v: 2,
					s: "e",
				}: {},
				{
					v: 3,
					s: "s",
				}: {},
				{
					v: 4,
					s: "m",
				}: {},
			},
		},
		"#3": {
			args: []test{
				{
					v: 2,
					s: "e",
				},
			},
			want: map[test]struct{}{
				{
					v: 2,
					s: "e",
				}: {},
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			got := ArrayRand[[]test](tt.args)

			if len(got) != len(tt.want) {
				t.Errorf("size ArrayRand() = %v, want %v", got, tt.want)
			}

			sim := 0

			for inx := range got {
				if got[inx] == tt.args[inx] {
					sim++
				}
			}

			if len(tt.args) == sim && len(tt.args) > 1 {
				t.Errorf("similar array ArrayRand %v, want %v", got, tt.args)
			}

			for _, value := range got {
				_, ok := tt.want[value]
				if !ok {
					t.Errorf("not found key %v", value)
				}

				delete(tt.want, value)
			}

			if len(tt.want) != 0 {
				t.Errorf("invalid keys %v", tt.want)
			}
		})
	}
}
