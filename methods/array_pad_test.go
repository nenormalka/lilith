package methods

import (
	"reflect"
	"testing"
)

func TestArrayPadInt(t *testing.T) {
	for name, tt := range map[string]struct {
		args  []int
		size  int
		value int
		want  []int
	}{
		"#1": {
			args:  nil,
			size:  0,
			value: 0,
			want:  nil,
		},
		"#2": {
			args:  []int{1, 2, 3, 4, 5},
			size:  2,
			value: 1,
			want:  []int{1, 2, 3, 4, 5},
		},
		"#3": {
			args:  []int{1, 2, 3},
			size:  5,
			value: 99,
			want:  []int{1, 2, 3, 99, 99},
		},
		"#4": {
			args:  nil,
			size:  2,
			value: 99,
			want:  []int{99, 99},
		},
	} {
		t.Run(name, func(t *testing.T) {
			if got := ArrayPad[[]int](tt.args, tt.size, tt.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayPad() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayPadStr(t *testing.T) {
	for name, tt := range map[string]struct {
		args  []string
		size  int
		value string
		want  []string
	}{
		"#1": {
			args:  nil,
			size:  0,
			value: "",
			want:  nil,
		},
		"#2": {
			args:  []string{"t", "e", "s", "t"},
			size:  2,
			value: "!!!",
			want:  []string{"t", "e", "s", "t"},
		},
		"#3": {
			args:  []string{"t", "e", "s", "t"},
			size:  5,
			value: " A!",
			want:  []string{"t", "e", "s", "t", " A!"},
		},
		"#4": {
			args:  nil,
			size:  1,
			value: " A!",
			want:  []string{" A!"},
		},
	} {
		t.Run(name, func(t *testing.T) {
			if got := ArrayPad[[]string](tt.args, tt.size, tt.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayPad() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayPadStruct(t *testing.T) {
	type test struct {
		v int
		s string
	}
	for name, tt := range map[string]struct {
		args  []test
		size  int
		value test
		want  []test
	}{
		"#1": {
			args:  nil,
			size:  0,
			value: test{},
			want:  nil,
		},
		"#2": {
			args: []test{
				{
					v: 1,
					s: "t",
				},
			},
			size:  1,
			value: test{},
			want: []test{
				{
					v: 1,
					s: "t",
				},
			},
		},
		"#3": {
			args: []test{
				{
					v: 1,
					s: "t",
				},
			},
			size: 3,
			value: test{
				v: 2,
				s: "e",
			},
			want: []test{
				{
					v: 1,
					s: "t",
				},
				{
					v: 2,
					s: "e",
				},
				{
					v: 2,
					s: "e",
				},
			},
		},
		"#4": {
			args: nil,
			size: 2,
			value: test{
				v: 34,
				s: "test",
			},
			want: []test{
				{
					v: 34,
					s: "test",
				},
				{
					v: 34,
					s: "test",
				},
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			if got := ArrayPad[[]test](tt.args, tt.size, tt.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayPad() = %v, want %v", got, tt.want)
			}
		})
	}
}
