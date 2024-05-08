package methods

func ArrayToMapValuesBool[S ~[]T, T comparable](s S) map[T]bool {
	if len(s) == 0 {
		return nil
	}

	m := make(map[T]bool, len(s))

	for i := range s {
		m[s[i]] = true
	}

	return m
}
