package methods

import "golang.org/x/exp/constraints"

func ArrayMax[S ~[]T, T constraints.Ordered](s S) T {
	var max T

	if len(s) == 0 {
		return max
	}

	if len(s) == 1 {
		return s[0]
	}

	max = s[0]

	for _, val := range s[1:] {
		if val > max {
			max = val
		}
	}

	return max
}
