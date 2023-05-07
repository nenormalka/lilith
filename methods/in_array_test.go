package methods

import (
	"testing"
)

func TestInArrayInt(t *testing.T) {
	for name, tt := range map[string]struct {
		args     []int
		wantElem int
		want     bool
	}{
		"#1": {
			args:     nil,
			wantElem: 0,
			want:     false,
		},
		"#2": {
			args:     []int{1, 2},
			wantElem: 0,
			want:     false,
		},
		"#3": {
			args:     []int{1, 2},
			wantElem: 1,
			want:     true,
		},
	} {
		t.Run(name, func(t *testing.T) {
			if InArray[[]int](tt.args, tt.wantElem) != tt.want {
				t.Errorf("InArray() = %v, want %v", tt.args, tt.wantElem)
			}
		})
	}
}

func TestInArrayStr(t *testing.T) {
	for name, tt := range map[string]struct {
		args     []string
		wantElem string
		want     bool
	}{
		"#1": {
			args:     nil,
			wantElem: "",
			want:     false,
		},
		"#2": {
			args:     []string{"t", "e"},
			wantElem: "s",
			want:     false,
		},
		"#3": {
			args:     []string{"t", "e"},
			wantElem: "e",
			want:     true,
		},
	} {
		t.Run(name, func(t *testing.T) {
			if InArray[[]string](tt.args, tt.wantElem) != tt.want {
				t.Errorf("InArray() = %v, want %v", tt.args, tt.wantElem)
			}
		})
	}
}

func TestInArrayStruct(t *testing.T) {
	type test struct {
		v int
		s string
	}
	for name, tt := range map[string]struct {
		args     []test
		wantElem test
		want     bool
	}{
		"#1": {
			args:     nil,
			wantElem: test{},
			want:     false,
		},
		"#2": {
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
			wantElem: test{
				v: 1,
				s: "e",
			},
			want: false,
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
			wantElem: test{
				v: 2,
				s: "e",
			},
			want: true,
		},
	} {
		t.Run(name, func(t *testing.T) {
			if InArray[[]test](tt.args, tt.wantElem) != tt.want {
				t.Errorf("InArray() = %v, want %v", tt.args, tt.wantElem)
			}
		})
	}
}
