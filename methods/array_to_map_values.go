package methods

func ArrayToMapValues[S ~[]T, T comparable](s S) map[T]struct{} {
	if len(s) == 0 {
		return nil
	}

	m := make(map[T]struct{}, len(s))

	for i := range s {
		m[s[i]] = struct{}{}
	}

	return m
}
