package methods

import "testing"

func TestInMapKeyInt(t *testing.T) {
	for name, tt := range map[string]struct {
		args    map[int]struct{}
		wantKey int
		want    bool
	}{
		"#1": {
			args:    nil,
			wantKey: 0,
			want:    false,
		},
		"#2": {
			args: map[int]struct{}{
				1: {},
			},
			wantKey: 2,
			want:    false,
		},
		"#3": {
			args: map[int]struct{}{
				1: {},
				2: {},
			},
			wantKey: 1,
			want:    true,
		},
	} {
		t.Run(name, func(t *testing.T) {
			if InMapKey[map[int]struct{}](tt.args, tt.wantKey) != tt.want {
				t.Errorf("InMapKey() = %v, want %v", tt.args, tt.wantKey)
			}
		})
	}
}

func TestInMapKeyStruct(t *testing.T) {
	type test struct {
		v int
		s string
	}
	for name, tt := range map[string]struct {
		args     map[test]struct{}
		wantElem test
		want     bool
	}{
		"#1": {
			args:     nil,
			wantElem: test{},
			want:     false,
		},
		"#2": {
			args: map[test]struct{}{
				test{
					v: 1,
					s: "t",
				}: {},
			},
			wantElem: test{
				v: 2,
				s: "t",
			},
			want: false,
		},
		"#3": {
			args: map[test]struct{}{
				test{
					v: 1,
					s: "t",
				}: {},
				test{
					v: 2,
					s: "e",
				}: {},
			},
			wantElem: test{
				v: 2,
				s: "e",
			},
			want: true,
		},
	} {
		t.Run(name, func(t *testing.T) {
			if InMapKey[map[test]struct{}](tt.args, tt.wantElem) != tt.want {
				t.Errorf("InMapKey() = %v, want %v", tt.args, tt.wantElem)
			}
		})
	}
}
