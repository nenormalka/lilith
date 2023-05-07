package methods

import (
	"reflect"
	"testing"
)

func TestArrayRemoveInt(t *testing.T) {
	for name, tt := range map[string]struct {
		args []int
		indx int
		want []int
	}{
		"#1": {
			args: nil,
			indx: 0,
			want: nil,
		},
		"#2": {
			args: nil,
			indx: -1,
			want: nil,
		},
		"#3": {
			args: []int{1, 2, 3},
			indx: 2,
			want: []int{1, 2},
		},
	} {
		t.Run(name, func(t *testing.T) {
			ArrayRemove[[]int](&tt.args, tt.indx)

			if !reflect.DeepEqual(tt.args, tt.want) {
				t.Errorf("ArrayRemove() = %v, want %v", tt.args, tt.want)
			}
		})
	}
}

func TestArrayRemoveStr(t *testing.T) {
	for name, tt := range map[string]struct {
		args []string
		indx int
		want []string
	}{
		"#1": {
			args: nil,
			indx: 0,
			want: nil,
		},
		"#2": {
			args: nil,
			indx: -1,
			want: nil,
		},
		"#3": {
			args: []string{"t", "e", "s"},
			indx: 0,
			want: []string{"e", "s"},
		},
	} {
		t.Run(name, func(t *testing.T) {
			ArrayRemove[[]string](&tt.args, tt.indx)

			if !reflect.DeepEqual(tt.args, tt.want) {
				t.Errorf("ArrayRemove() = %v, want %v", tt.args, tt.want)
			}
		})
	}
}

func TestArrayRemoveStruct(t *testing.T) {
	type test struct {
		v int
		s string
	}
	for name, tt := range map[string]struct {
		args []test
		indx int
		want []test
	}{
		"#1": {
			args: nil,
			indx: 0,
			want: nil,
		},
		"#2": {
			args: nil,
			indx: -1,
			want: nil,
		},
		"#3": {
			args: []test{
				{
					v: 0,
					s: "test",
				},
			},
			indx: 0,
			want: []test{},
		},
	} {
		t.Run(name, func(t *testing.T) {
			ArrayRemove[[]test](&tt.args, tt.indx)

			if !reflect.DeepEqual(tt.args, tt.want) {
				t.Errorf("ArrayRemove() = %v, want %v", tt.args, tt.want)
			}
		})
	}
}
