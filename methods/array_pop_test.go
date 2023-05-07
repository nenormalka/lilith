package methods

import (
	"reflect"
	"testing"
)

func TestArrayPopInt(t *testing.T) {
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
			want:     []int{1},
			wantElem: 2,
		},
	} {
		t.Run(name, func(t *testing.T) {
			elem := ArrayPop[[]int](&tt.args)
			if elem != tt.wantElem {
				t.Errorf("elem ArrayPop() = %v, want %v", elem, tt.wantElem)
			}

			if !reflect.DeepEqual(tt.args, tt.want) {
				t.Errorf("ArrayPop() = %v, want %v", tt.args, tt.want)
			}
		})
	}
}

func TestArrayPopStr(t *testing.T) {
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
			args:     []string{"t", "e", "s", "t"},
			want:     []string{"t", "e", "s"},
			wantElem: "t",
		},
	} {
		t.Run(name, func(t *testing.T) {
			elem := ArrayPop[[]string](&tt.args)
			if elem != tt.wantElem {
				t.Errorf("elem ArrayPop() = %v, want %v", elem, tt.wantElem)
			}

			if !reflect.DeepEqual(tt.args, tt.want) {
				t.Errorf("ArrayPop() = %v, want %v", tt.args, tt.want)
			}
		})
	}
}

func TestArrayPopStruct(t *testing.T) {
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
					v: 1,
					s: "t",
				},
				{
					v: 2,
					s: "e",
				},
			},
			wantElem: test{
				v: 3,
				s: "s",
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			elem := ArrayPop[[]test](&tt.args)
			if elem != tt.wantElem {
				t.Errorf("elem ArrayPop() = %v, want %v", elem, tt.wantElem)
			}

			if !reflect.DeepEqual(tt.args, tt.want) {
				t.Errorf("ArrayPop() = %v, want %v", tt.args, tt.want)
			}
		})
	}
}
