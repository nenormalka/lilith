package methods

import (
	"reflect"
	"testing"
)

func TestMapFlipInt(t *testing.T) {
	for name, tt := range map[string]struct {
		args map[string]int
		want map[int]string
	}{
		"#1": {
			args: nil,
			want: map[int]string{},
		},
		"#2": {
			args: map[string]int{
				"t": 11,
				"e": 22,
				"s": 33,
			},
			want: map[int]string{
				11: "t",
				22: "e",
				33: "s",
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			if got := MapFlip[map[string]int, map[int]string](tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayUnshift() = %v, want %v", tt.args, tt.want)
			}
		})
	}
}

func TestArrayFlipInt(t *testing.T) {
	type test struct {
		v int
		s string
	}
	for name, tt := range map[string]struct {
		args map[test]int
		want map[int]test
	}{
		"#1": {
			args: nil,
			want: map[int]test{},
		},
		"#2": {
			args: map[test]int{
				{
					v: 1,
					s: "t",
				}: 11,
				{
					v: 2,
					s: "e",
				}: 22,
				{
					v: 3,
					s: "s",
				}: 33,
			},
			want: map[int]test{
				11: {
					v: 1,
					s: "t",
				},
				22: {
					v: 2,
					s: "e",
				},
				33: {
					v: 3,
					s: "s",
				},
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			if got := MapFlip[map[test]int, map[int]test](tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayUnshift() = %v, want %v", tt.args, tt.want)
			}
		})
	}
}
