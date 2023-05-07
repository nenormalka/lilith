package methods

import "errors"

var (
	ErrMapCombine = errors.New("slices must has equal length")
)

func MapCombine[M ~map[K]V, K, V comparable](s1 []K, s2 []V) (M, error) {
	if len(s1) != len(s2) {
		return nil, ErrMapCombine
	}

	m := make(map[K]V, len(s1))
	for i, v := range s1 {
		m[v] = s2[i]
	}

	return m, nil
}
