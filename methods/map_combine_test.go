package methods

import (
	"reflect"
	"testing"
)

func TestMapCombineInt(t *testing.T) {
	for name, tt := range map[string]struct {
		args1 []int
		args2 []int
		err   error
		want  map[int]int
	}{
		"#1": {
			args1: nil,
			args2: nil,
			want:  map[int]int{},
			err:   nil,
		},
		"#2": {
			args1: []int{1, 2},
			args2: nil,
			want:  nil,
			err:   ErrMapCombine,
		},
		"#3": {
			args1: nil,
			args2: []int{1, 2},
			want:  nil,
			err:   ErrMapCombine,
		},
		"#4": {
			args1: []int{11, 22},
			args2: []int{1, 2},
			want: map[int]int{
				11: 1,
				22: 2,
			},
			err: nil,
		},
	} {
		t.Run(name, func(t *testing.T) {
			got, err := MapCombine[map[int]int](tt.args1, tt.args2)
			if tt.err != err {
				t.Errorf("MapCombine() err = %v, tt.err %v", err, tt.err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapCombine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapCombineStructStr(t *testing.T) {
	type test struct {
		v int
		s string
	}
	for name, tt := range map[string]struct {
		args1 []test
		args2 []string
		err   error
		want  map[test]string
	}{
		"#1": {
			args1: nil,
			args2: nil,
			want:  map[test]string{},
			err:   nil,
		},
		"#2": {
			args1: []test{
				{
					v: 1,
					s: "test",
				},
			},
			args2: nil,
			want:  nil,
			err:   ErrMapCombine,
		},
		"#3": {
			args1: nil,
			args2: []string{"test"},
			want:  nil,
			err:   ErrMapCombine,
		},
		"#4": {
			args1: []test{
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
			args2: []string{"m", "z", "l"},
			want: map[test]string{
				{
					v: 1,
					s: "t",
				}: "m",
				{
					v: 2,
					s: "e",
				}: "z",
				{
					v: 3,
					s: "s",
				}: "l",
			},
			err: nil,
		},
	} {
		t.Run(name, func(t *testing.T) {
			got, err := MapCombine[map[test]string](tt.args1, tt.args2)
			if tt.err != err {
				t.Errorf("MapCombine() err = %v, tt.err %v", err, tt.err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapCombine() = %v, want %v", got, tt.want)
			}
		})
	}
}
