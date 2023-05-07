package methods

import (
	"reflect"
	"testing"
)

func TestArraySliceInt(t *testing.T) {
	for name, tt := range map[string]struct {
		args   []int
		offset int
		length int
		err    error
		want   []int
	}{
		"#1": {
			args: nil,
			want: nil,
		},
		"#2": {
			args:   []int{1},
			offset: 2,
			length: 0,
			err:    ErrArraySliceOffset,
			want:   nil,
		},
		"#3": {
			args:   []int{1, 2},
			offset: 2,
			length: 3,
			err:    nil,
			want:   []int{},
		},
		"#4": {
			args:   []int{1, 2, 3},
			offset: 2,
			length: 1,
			err:    nil,
			want:   []int{3},
		},
		"#5": {
			args:   []int{1, 2, 3, 4},
			offset: 2,
			length: 1,
			err:    nil,
			want:   []int{3},
		},
		"#6": {
			args:   []int{1, 2, 3, 4},
			offset: 2,
			length: 4,
			err:    nil,
			want:   []int{3, 4},
		},
		"#7": {
			args:   []int{1, 2, 3, 4},
			offset: 0,
			length: 4,
			err:    nil,
			want:   []int{1, 2, 3, 4},
		},
	} {
		t.Run(name, func(t *testing.T) {
			got, err := ArraySlice[[]int](tt.args, tt.offset, tt.length)
			if tt.err != err {
				t.Errorf("ArraySlice() err = %v, tt.err %v", err, tt.err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArraySlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArraySliceStr(t *testing.T) {
	for name, tt := range map[string]struct {
		args   []string
		offset int
		length int
		err    error
		want   []string
	}{
		"#1": {
			args: nil,
			want: nil,
		},
		"#2": {
			args:   []string{"t"},
			offset: 2,
			length: 0,
			err:    ErrArraySliceOffset,
			want:   nil,
		},
		"#3": {
			args:   []string{"t", "e"},
			offset: 2,
			length: 3,
			err:    nil,
			want:   []string{},
		},
		"#4": {
			args:   []string{"t", "e", "s"},
			offset: 2,
			length: 1,
			err:    nil,
			want:   []string{"s"},
		},
		"#5": {
			args:   []string{"t", "e", "s", "t"},
			offset: 2,
			length: 1,
			err:    nil,
			want:   []string{"s"},
		},
		"#6": {
			args:   []string{"t", "e", "s", "t"},
			offset: 2,
			length: 4,
			err:    nil,
			want:   []string{"s", "t"},
		},
		"#7": {
			args:   []string{"t", "e", "s", "t"},
			offset: 0,
			length: 4,
			err:    nil,
			want:   []string{"t", "e", "s", "t"},
		},
	} {
		t.Run(name, func(t *testing.T) {
			got, err := ArraySlice[[]string](tt.args, tt.offset, tt.length)
			if tt.err != err {
				t.Errorf("ArraySlice() err = %v, tt.err %v", err, tt.err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArraySlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArraySliceStruct(t *testing.T) {
	type test struct {
		v int
		s string
	}
	for name, tt := range map[string]struct {
		args   []test
		offset int
		length int
		err    error
		want   []test
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
			},
			offset: 2,
			length: 0,
			err:    ErrArraySliceOffset,
			want:   nil,
		},
		"#3": {
			args: []test{
				{
					v: 1,
					s: "t",
				},
				{
					v: 2,
					s: "e",
				},
			},
			offset: 2,
			length: 3,
			err:    nil,
			want:   []test{},
		},
		"#4": {
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
			},
			offset: 2,
			length: 1,
			err:    nil,
			want: []test{
				{
					v: 3,
					s: "s",
				},
			},
		},
		"#5": {
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
			},
			offset: 2,
			length: 1,
			err:    nil,
			want: []test{
				{
					v: 3,
					s: "s",
				},
			},
		},
		"#6": {
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
			offset: 2,
			length: 4,
			err:    nil,
			want: []test{
				{
					v: 3,
					s: "s",
				},
				{
					v: 4,
					s: "t",
				},
			},
		},
		"#7": {
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
			offset: 0,
			length: 4,
			err:    nil,
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
					v: 3,
					s: "s",
				},
				{
					v: 4,
					s: "t",
				},
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			got, err := ArraySlice[[]test](tt.args, tt.offset, tt.length)
			if tt.err != err {
				t.Errorf("ArraySlice() err = %v, tt.err %v", err, tt.err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArraySlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
