package methods

func ArrayFlip[S ~[]T, T comparable](s S) map[T]int {
	if len(s) == 0 {
		return nil
	}

	m := make(map[T]int, len(s))

	for i := range s {
		m[s[i]] = i
	}

	return m
}
