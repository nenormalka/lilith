package methods

import (
	"reflect"
	"testing"
)

func TestArrayPushInt(t *testing.T) {
	for name, tt := range map[string]struct {
		args  []int
		elems []int
		want  []int
	}{
		"#1": {
			args:  nil,
			elems: nil,
			want:  nil,
		},
		"#2": {
			args:  nil,
			elems: []int{1, 2, 3},
			want:  []int{1, 2, 3},
		},
		"#3": {
			args:  []int{1, 2, 3},
			elems: []int{4, 5, 6},
			want:  []int{1, 2, 3, 4, 5, 6},
		},
		"#4": {
			args:  []int{1, 2, 3},
			elems: nil,
			want:  []int{1, 2, 3},
		},
	} {
		t.Run(name, func(t *testing.T) {
			ArrayPush[[]int](&tt.args, tt.elems...)

			if !reflect.DeepEqual(tt.args, tt.want) {
				t.Errorf("ArrayPush() = %v, want %v", tt.args, tt.want)
			}
		})
	}
}

func TestArrayPushStr(t *testing.T) {
	for name, tt := range map[string]struct {
		args  []string
		elems []string
		want  []string
	}{
		"#1": {
			args:  nil,
			elems: nil,
			want:  nil,
		},
		"#2": {
			args:  nil,
			elems: []string{"t", "e", "s", "t"},
			want:  []string{"t", "e", "s", "t"},
		},
		"#3": {
			args:  []string{"t", "e"},
			elems: []string{"s", "t"},
			want:  []string{"t", "e", "s", "t"},
		},
		"#4": {
			args:  []string{"t", "e", "s", "t"},
			elems: nil,
			want:  []string{"t", "e", "s", "t"},
		},
	} {
		t.Run(name, func(t *testing.T) {
			ArrayPush[[]string](&tt.args, tt.elems...)

			if !reflect.DeepEqual(tt.args, tt.want) {
				t.Errorf("ArrayPush() = %v, want %v", tt.args, tt.want)
			}
		})
	}
}

func TestArrayPushStruct(t *testing.T) {
	type test struct {
		v int
		s string
	}
	for name, tt := range map[string]struct {
		args  []test
		elems []test
		want  []test
	}{
		"#1": {
			args:  nil,
			elems: nil,
			want:  nil,
		},
		"#2": {
			args: nil,
			elems: []test{
				{
					v: 1,
					s: "t",
				},
				{
					v: 2,
					s: "e",
				},
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
			},
		},
		"#3": {
			args: []test{
				{
					v: 1,
					s: "t",
				},
			},
			elems: []test{
				{
					v: 2,
					s: "e",
				},
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
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			ArrayPush[[]test](&tt.args, tt.elems...)

			if !reflect.DeepEqual(tt.args, tt.want) {
				t.Errorf("ArrayPush() = %v, want %v", tt.args, tt.want)
			}
		})
	}
}
