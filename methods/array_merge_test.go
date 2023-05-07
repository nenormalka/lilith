package methods

import (
	"reflect"
	"testing"
)

func TestArrayMergeInt(t *testing.T) {
	for name, tt := range map[string]struct {
		args [][]int
		want []int
	}{
		"#1": {
			args: nil,
			want: nil,
		},
		"#2": {
			args: [][]int{{0, 1, 2}, {3, 4}, {5, 6, 7}},
			want: []int{0, 1, 2, 3, 4, 5, 6, 7},
		},
		"#3": {
			args: [][]int{{1, 2, 3}},
			want: []int{1, 2, 3},
		},
	} {
		t.Run(name, func(t *testing.T) {
			if got := ArrayMerge[[]int](tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayMergeStr(t *testing.T) {
	for name, tt := range map[string]struct {
		args [][]string
		want []string
	}{
		"#1": {
			args: nil,
			want: nil,
		},
		"#2": {
			args: [][]string{{"M", "A", "M"}, {"A", " "}, {"M", "bl", "JL", "A"}},
			want: []string{"M", "A", "M", "A", " ", "M", "bl", "JL", "A"},
		},
		"#3": {
			args: [][]string{{"M", "A", "M", "A"}},
			want: []string{"M", "A", "M", "A"},
		},
	} {
		t.Run(name, func(t *testing.T) {
			if got := ArrayMerge[[]string](tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayMergeStruct(t *testing.T) {
	type test struct {
		v int
		s string
	}
	for name, tt := range map[string]struct {
		args [][]test
		want []test
	}{
		"#1": {
			args: nil,
			want: nil,
		},
		"#2": {
			args: [][]test{
				{
					{
						v: 1,
						s: "m",
					},
					{
						v: 2,
						s: "a",
					},
				},
				{
					{
						v: 3,
						s: "m",
					},
				},
				{
					{
						v: 4,
						s: "a",
					},
				},
			},
			want: []test{
				{
					v: 1,
					s: "m",
				},
				{
					v: 2,
					s: "a",
				},
				{
					v: 3,
					s: "m",
				},
				{
					v: 4,
					s: "a",
				},
			},
		},
		"#3": {
			args: [][]test{{{v: 1, s: "test"}}},
			want: []test{{v: 1, s: "test"}},
		},
	} {
		t.Run(name, func(t *testing.T) {
			if got := ArrayMerge[[]test](tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}
