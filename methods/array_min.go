package methods

import "golang.org/x/exp/constraints"

func ArrayMin[S ~[]T, T constraints.Ordered](s S) T {
	var min T

	if len(s) == 0 {
		return min
	}

	if len(s) == 1 {
		return s[0]
	}

	min = s[0]

	for _, val := range s[1:] {
		if val < min {
			min = val
		}
	}

	return min
}
