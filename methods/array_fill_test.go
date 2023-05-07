package methods

import (
	"reflect"
	"testing"
)

func TestArrayFillInt(t *testing.T) {
	for name, tt := range map[string]struct {
		args struct {
			startIndex int
			num        int
			value      int
		}
		want []int
	}{
		"#1": {
			args: struct {
				startIndex int
				num        int
				value      int
			}{
				startIndex: 0,
				num:        0,
				value:      0,
			},
			want: nil,
		},
		"#2": {
			args: struct {
				startIndex int
				num        int
				value      int
			}{
				startIndex: 2,
				num:        3,
				value:      99,
			},
			want: []int{0, 0, 99, 99, 99},
		},
		"#3": {
			args: struct {
				startIndex int
				num        int
				value      int
			}{
				startIndex: 0,
				num:        5,
				value:      7,
			},
			want: []int{7, 7, 7, 7, 7},
		},
	} {
		t.Run(name, func(t *testing.T) {
			if got := ArrayFill[[]int](tt.args.startIndex, tt.args.num, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayFill() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayFillStr(t *testing.T) {
	for name, tt := range map[string]struct {
		args struct {
			startIndex int
			num        int
			value      string
		}
		want []string
	}{
		"#1": {
			args: struct {
				startIndex int
				num        int
				value      string
			}{
				startIndex: 0,
				num:        0,
				value:      "",
			},
			want: nil,
		},
		"#2": {
			args: struct {
				startIndex int
				num        int
				value      string
			}{
				startIndex: 2,
				num:        3,
				value:      "test",
			},
			want: []string{"", "", "test", "test", "test"},
		},
		"#3": {
			args: struct {
				startIndex int
				num        int
				value      string
			}{
				startIndex: 0,
				num:        5,
				value:      "atmta",
			},
			want: []string{"atmta", "atmta", "atmta", "atmta", "atmta"},
		},
	} {
		t.Run(name, func(t *testing.T) {
			if got := ArrayFill[[]string](tt.args.startIndex, tt.args.num, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayFill() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayFillStruct(t *testing.T) {
	type test struct {
		t int
		v int
		s string
	}

	for name, tt := range map[string]struct {
		args struct {
			startIndex int
			num        int
			value      test
		}
		want []test
	}{
		"#1": {
			args: struct {
				startIndex int
				num        int
				value      test
			}{
				startIndex: 0,
				num:        0,
				value: test{
					t: 1,
					v: 2,
					s: "test",
				},
			},
			want: nil,
		},
		"#2": {
			args: struct {
				startIndex int
				num        int
				value      test
			}{
				startIndex: 2,
				num:        3,
				value: test{
					t: 1,
					v: 2,
					s: "test",
				},
			},
			want: []test{
				{
					t: 0,
					v: 0,
					s: "",
				},
				{
					t: 0,
					v: 0,
					s: "",
				},
				{
					t: 1,
					v: 2,
					s: "test",
				},
				{
					t: 1,
					v: 2,
					s: "test",
				},
				{
					t: 1,
					v: 2,
					s: "test",
				},
			},
		},
		"#3": {
			args: struct {
				startIndex int
				num        int
				value      test
			}{
				startIndex: 0,
				num:        5,
				value: test{
					t: 1,
					v: 2,
					s: "test",
				},
			},
			want: []test{
				{
					t: 1,
					v: 2,
					s: "test",
				},
				{
					t: 1,
					v: 2,
					s: "test",
				},
				{
					t: 1,
					v: 2,
					s: "test",
				},
				{
					t: 1,
					v: 2,
					s: "test",
				},
				{
					t: 1,
					v: 2,
					s: "test",
				},
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			if got := ArrayFill[[]test](tt.args.startIndex, tt.args.num, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayFill() = %v, want %v", got, tt.want)
			}
		})
	}
}
