package methods

import "testing"

func TestInMapInt(t *testing.T) {
	for name, tt := range map[string]struct {
		args     map[string]int
		wantElem int
		want     bool
	}{
		"#1": {
			args:     nil,
			wantElem: 0,
			want:     false,
		},
		"#2": {
			args: map[string]int{
				"t": 1,
				"m": 2,
			},
			wantElem: 0,
			want:     false,
		},
		"#3": {
			args: map[string]int{
				"t": 1,
				"m": 2,
			},
			wantElem: 1,
			want:     true,
		},
	} {
		t.Run(name, func(t *testing.T) {
			if InMap[map[string]int](tt.args, tt.wantElem) != tt.want {
				t.Errorf("InMap() = %v, want %v", tt.args, tt.wantElem)
			}
		})
	}
}

func TestInMapStruct(t *testing.T) {
	type test struct {
		v int
		s string
	}
	for name, tt := range map[string]struct {
		args     map[int]test
		wantElem test
		want     bool
	}{
		"#1": {
			args:     nil,
			wantElem: test{},
			want:     false,
		},
		"#2": {
			args: map[int]test{
				1: {
					v: 1,
					s: "t",
				},
				2: {
					v: 2,
					s: "e",
				},
			},
			wantElem: test{
				v: 2,
				s: "t",
			},
			want: false,
		},
		"#3": {
			args: map[int]test{
				1: {
					v: 1,
					s: "t",
				},
				2: {
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
			if InMap[map[int]test](tt.args, tt.wantElem) != tt.want {
				t.Errorf("InMap() = %v, want %v", tt.args, tt.wantElem)
			}
		})
	}
}
