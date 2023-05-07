package methods

import (
	"reflect"
	"testing"
)

func TestArrayShiftInt(t *testing.T) {
	for name, tt := range map[string]struct {
		args     []int
		want     []int
		wantElem int
	}{
		"#1": {
			args:     nil,
			want:     nil,
			wantElem: 0,
		},
		"#2": {
			args:     []int{1},
			want:     []int{},
			wantElem: 1,
		},
		"#3": {
			args:     []int{1, 2},
			want:     []int{2},
			wantElem: 1,
		},
	} {
		t.Run(name, func(t *testing.T) {
			elem := ArrayShift[[]int](&tt.args)
			if elem != tt.wantElem {
				t.Errorf("elem ArrayShift() = %v, want %v", elem, tt.wantElem)
			}

			if !reflect.DeepEqual(tt.args, tt.want) {
				t.Errorf("ArrayShift() = %v, want %v", tt.args, tt.want)
			}
		})
	}
}

func TestArrayShiftStr(t *testing.T) {
	for name, tt := range map[string]struct {
		args     []string
		want     []string
		wantElem string
	}{
		"#1": {
			args:     nil,
			want:     nil,
			wantElem: "",
		},
		"#2": {
			args:     []string{"test"},
			want:     []string{},
			wantElem: "test",
		},
		"#3": {
			args:     []string{"m", "e", "s", "t"},
			want:     []string{"e", "s", "t"},
			wantElem: "m",
		},
	} {
		t.Run(name, func(t *testing.T) {
			elem := ArrayShift[[]string](&tt.args)
			if elem != tt.wantElem {
				t.Errorf("elem ArrayShift() = %v, want %v", elem, tt.wantElem)
			}

			if !reflect.DeepEqual(tt.args, tt.want) {
				t.Errorf("ArrayShift() = %v, want %v", tt.args, tt.want)
			}
		})
	}
}

func TestArrayShiftStruct(t *testing.T) {
	type test struct {
		v int
		s string
	}
	for name, tt := range map[string]struct {
		args     []test
		want     []test
		wantElem test
	}{
		"#1": {
			args:     nil,
			want:     nil,
			wantElem: test{},
		},
		"#2": {
			args: []test{
				{
					v: 1,
					s: "test",
				},
			},
			want: []test{},
			wantElem: test{
				v: 1,
				s: "test",
			},
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
				{
					v: 3,
					s: "s",
				},
			},
			want: []test{
				{
					v: 2,
					s: "e",
				},
				{
					v: 3,
					s: "s",
				},
			},
			wantElem: test{
				v: 1,
				s: "t",
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			elem := ArrayShift[[]test](&tt.args)
			if elem != tt.wantElem {
				t.Errorf("elem ArrayShift() = %v, want %v", elem, tt.wantElem)
			}

			if !reflect.DeepEqual(tt.args, tt.want) {
				t.Errorf("ArrayShift() = %v, want %v", tt.args, tt.want)
			}
		})
	}
}
