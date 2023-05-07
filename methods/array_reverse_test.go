package methods

import (
	"reflect"
	"testing"
)

func TestArrayReverseInt(t *testing.T) {
	for name, tt := range map[string]struct {
		args []int
		want []int
	}{
		"#1": {
			args: nil,
			want: nil,
		},
		"#2": {
			args: []int{1, 2, 3},
			want: []int{3, 2, 1},
		},
	} {
		t.Run(name, func(t *testing.T) {
			got := ArrayReverse[[]int](tt.args)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayReverse() = %v, want %v", tt.args, tt.want)
			}
		})
	}
}

func TestArrayReverseStr(t *testing.T) {
	for name, tt := range map[string]struct {
		args []string
		want []string
	}{
		"#1": {
			args: nil,
			want: nil,
		},
		"#2": {
			args: []string{"t", "e", "s", "t"},
			want: []string{"t", "s", "e", "t"},
		},
	} {
		t.Run(name, func(t *testing.T) {
			got := ArrayReverse[[]string](tt.args)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayReverse() = %v, want %v", tt.args, tt.want)
			}
		})
	}
}

func TestArrayReverseStruct(t *testing.T) {
	type test struct {
		v int
		s string
	}
	for name, tt := range map[string]struct {
		args []test
		want []test
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
					s: "t",
				},
			},
			want: []test{
				{
					v: 4,
					s: "t",
				},
				{
					v: 3,
					s: "s",
				},
				{
					v: 2,
					s: "e",
				},
				{
					v: 1,
					s: "t",
				},
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			got := ArrayReverse[[]test](tt.args)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayReverse() = %v, want %v", tt.args, tt.want)
			}
		})
	}
}
