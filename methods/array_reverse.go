package methods

func ArrayReverse[S ~[]T, T any](s S) S {
	if len(s) == 0 {
		return nil
	}

	m := make([]T, len(s))
	copy(m, s)

	for i, j := 0, len(m)-1; i < j; i, j = i+1, j-1 {
		m[i], m[j] = m[j], m[i]
	}

	return m
}
