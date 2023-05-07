package methods

import (
	"reflect"
	"testing"
)

func TestArrayCopyInt(t *testing.T) {
	for name, tt := range map[string]struct {
		args []int
		want []int
	}{
		"#1": {
			args: nil,
			want: []int{},
		},
		"#2": {
			args: []int{},
			want: []int{},
		},
		"#3": {
			args: []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
	} {
		t.Run(name, func(t *testing.T) {
			got := ArrayCopy[[]int](tt.args)
			if got != nil && &got == &tt.args {
				t.Errorf("equal pointers ArrayCopy() = %v, want %v", &got, &tt.want)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayCopy() = %v, want %v", tt.args, tt.want)
			}
		})
	}
}

func TestArrayCopyStr(t *testing.T) {
	for name, tt := range map[string]struct {
		args []string
		want []string
	}{
		"#1": {
			args: nil,
			want: []string{},
		},
		"#2": {
			args: []string{},
			want: []string{},
		},
		"#3": {
			args: []string{"t", "e", "s", "t"},
			want: []string{"t", "e", "s", "t"},
		},
	} {
		t.Run(name, func(t *testing.T) {
			got := ArrayCopy[[]string](tt.args)
			if got != nil && &got == &tt.args {
				t.Errorf("equal pointers ArrayCopy() = %v, want %v", &got, &tt.want)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayCopy() = %v, want %v", tt.args, tt.want)
			}
		})
	}
}

func TestArrayCopyStruct(t *testing.T) {
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
			want: []test{},
		},
		"#2": {
			args: []test{},
			want: []test{},
		},
		"#3": {
			args: []test{
				{
					v: 0,
					s: "test",
				},
			},
			want: []test{
				{
					v: 0,
					s: "test",
				},
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			got := ArrayCopy[[]test](tt.args)
			if got != nil && &got == &tt.args {
				t.Errorf("equal pointers ArrayCopy() = %v, want %v", &got, &tt.want)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayCopy() = %v, want %v", tt.args, tt.want)
			}
		})
	}
}
