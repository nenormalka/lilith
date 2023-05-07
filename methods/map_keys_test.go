package methods

import (
	"testing"
)

func TestMapKeysInt(t *testing.T) {
	for name, tt := range map[string]struct {
		args map[int]string
		want map[int]struct{}
	}{
		"#1": {
			args: map[int]string{},
			want: nil,
		},
		"#2": {
			args: map[int]string{1: "", 44: "", -123: ""},
			want: map[int]struct{}{
				1:    {},
				44:   {},
				-123: {},
			},
		},
		"#3": {
			args: map[int]string{1: ""},
			want: map[int]struct{}{
				1: {},
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			got := MapKeys(tt.args)

			if len(got) != len(tt.want) {
				t.Errorf("size MapKeys() = %v, want %v", got, tt.want)
			}

			for _, k := range got {
				if _, ok := tt.want[k]; !ok {
					t.Errorf("not found key %v", k)
				}

				delete(tt.want, k)
			}

			if len(tt.want) != 0 {
				t.Errorf("invalid keys %v", tt.want)
			}
		})
	}
}
