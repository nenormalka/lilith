package methods

import "testing"

func TestMapValuesInt(t *testing.T) {
	for name, tt := range map[string]struct {
		args map[int]int
		want map[int]struct{}
	}{
		"#1": {
			args: map[int]int{},
			want: nil,
		},
		"#2": {
			args: map[int]int{1: 12, 44: 13, -123: -14},
			want: map[int]struct{}{
				12:  {},
				13:  {},
				-14: {},
			},
		},
		"#3": {
			args: map[int]int{1: 111},
			want: map[int]struct{}{
				111: {},
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			got := MapValues(tt.args)

			if len(got) != len(tt.want) {
				t.Errorf("size MapValues() = %v, want %v", got, tt.want)
			}

			for _, k := range got {
				_, ok := tt.want[k]
				if !ok {
					t.Errorf("not found value %v", k)
				}

				delete(tt.want, k)
			}

			if len(tt.want) != 0 {
				t.Errorf("invalid values %v", tt.want)
			}
		})
	}
}

func TestMapValuesStr(t *testing.T) {
	for name, tt := range map[string]struct {
		args map[int]string
		want map[string]struct{}
	}{
		"#1": {
			args: map[int]string{},
			want: nil,
		},
		"#2": {
			args: map[int]string{1: "12", 44: "13", -123: "-14"},
			want: map[string]struct{}{
				"12":  {},
				"13":  {},
				"-14": {},
			},
		},
		"#3": {
			args: map[int]string{1: "111"},
			want: map[string]struct{}{
				"111": {},
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			got := MapValues(tt.args)

			if len(got) != len(tt.want) {
				t.Errorf("size MapValues() = %v, want %v", got, tt.want)
			}

			for _, k := range got {
				_, ok := tt.want[k]
				if !ok {
					t.Errorf("not found value %v", k)
				}

				delete(tt.want, k)
			}

			if len(tt.want) != 0 {
				t.Errorf("invalid values %v", tt.want)
			}
		})
	}
}

func TestMapValuesStruct(t *testing.T) {
	type test struct {
		t int
		s string
	}

	for name, tt := range map[string]struct {
		args map[int]test
		want map[test]struct{}
	}{
		"#1": {
			args: map[int]test{},
			want: nil,
		},
		"#2": {
			args: map[int]test{
				1: {
					t: 12,
					s: "test",
				},
				2: {
					t: 13,
					s: "test1",
				},
			},
			want: map[test]struct{}{
				{
					t: 12,
					s: "test",
				}: {},
				{
					t: 13,
					s: "test1",
				}: {},
			},
		},
		"#3": {
			args: map[int]test{
				1: {
					t: 12,
					s: "test",
				},
			},
			want: map[test]struct{}{
				{
					t: 12,
					s: "test",
				}: {},
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			got := MapValues(tt.args)

			if len(got) != len(tt.want) {
				t.Errorf("size MapValues() = %v, want %v", got, tt.want)
			}

			for _, k := range got {
				_, ok := tt.want[k]
				if !ok {
					t.Errorf("not found value %v", k)
				}

				delete(tt.want, k)
			}

			if len(tt.want) != 0 {
				t.Errorf("invalid values %v", tt.want)
			}
		})
	}
}
