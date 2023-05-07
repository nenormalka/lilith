package methods

import (
	"reflect"
	"testing"
)

func TestArrayChunkInt(t *testing.T) {
	for name, tt := range map[string]struct {
		args []int
		size int
		err  error
		want [][]int
	}{
		"#1": {
			args: nil,
			size: 0,
			want: nil,
			err:  ErrArrayChunkSize,
		},
		"#2": {
			args: []int{1, 2, 3, 4, 5},
			size: 2,
			want: [][]int{{1, 2}, {3, 4}, {5}},
		},
		"#3": {
			args: []int{1, 2, 3},
			size: 4,
			want: [][]int{{1, 2, 3}},
		},
	} {
		t.Run(name, func(t *testing.T) {
			got, err := ArrayChunk[[]int](tt.args, tt.size)
			if tt.err != err {
				t.Errorf("ArrayChunk() err = %v, tt.err %v", err, tt.err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayChunk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayChunkStr(t *testing.T) {
	for name, tt := range map[string]struct {
		args []string
		size int
		err  error
		want [][]string
	}{
		"#1": {
			args: nil,
			size: 0,
			want: nil,
			err:  ErrArrayChunkSize,
		},
		"#2": {
			args: []string{"t", "e", "s", "t"},
			size: 3,
			want: [][]string{{"t", "e", "s"}, {"t"}},
		},
		"#3": {
			args: []string{"t", "e", "s", "t"},
			size: 4,
			want: [][]string{{"t", "e", "s", "t"}},
		},
	} {
		t.Run(name, func(t *testing.T) {
			got, err := ArrayChunk[[]string](tt.args, tt.size)
			if tt.err != err {
				t.Errorf("ArrayChunk() err = %v, tt.err %v", err, tt.err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayChunk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayChunkStruct(t *testing.T) {
	type test struct {
		v int
		s string
	}

	for name, tt := range map[string]struct {
		args []test
		size int
		err  error
		want [][]test
	}{
		"#1": {
			args: nil,
			size: 0,
			want: nil,
			err:  ErrArrayChunkSize,
		},
		"#2": {
			args: []test{
				{
					v: 1,
					s: "a",
				},
				{
					v: 2,
					s: "b",
				},
				{
					v: 3,
					s: "c",
				},
			},
			size: 2,
			want: [][]test{
				{
					{
						v: 1,
						s: "a",
					},
					{
						v: 2,
						s: "b",
					},
				},
				{
					{
						v: 3,
						s: "c",
					},
				},
			},
		},
		"#3": {
			args: []test{
				{
					v: 1,
					s: "a",
				},
				{
					v: 2,
					s: "b",
				},
				{
					v: 3,
					s: "c",
				},
			},
			size: 4,
			want: [][]test{
				{
					{
						v: 1,
						s: "a",
					},
					{
						v: 2,
						s: "b",
					},
					{
						v: 3,
						s: "c",
					},
				},
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			got, err := ArrayChunk[[]test](tt.args, tt.size)
			if tt.err != err {
				t.Errorf("ArrayChunk() err = %v, tt.err %v", err, tt.err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayChunk() = %v, want %v", got, tt.want)
			}
		})
	}
}
