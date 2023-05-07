package methods

import (
	"reflect"
	"testing"
)

func TestTernary(t *testing.T) {
	for name, tt := range map[string]struct {
		cond   bool
		first  int
		second int
		want   int
	}{
		"#1": {
			cond:   func() bool { return true }(),
			first:  1,
			second: 2,
			want:   1,
		},
		"#2": {
			cond:   func() bool { return false }(),
			first:  1,
			second: 2,
			want:   2,
		},
	} {
		t.Run(name, func(t *testing.T) {
			if got := Ternary[int](tt.cond, tt.first, tt.second); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ternary() = %v, want %v", got, tt.want)
			}
		})
	}
}
