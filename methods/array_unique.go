package methods

func ArrayUnique[S ~[]T, T comparable](s S) S {
	if len(s) == 0 {
		return s
	}

	var arr S

	m := make(map[T]struct{})

	for i := range s {
		if _, ok := m[s[i]]; ok {
			continue
		}

		arr = append(arr, s[i])
		m[s[i]] = struct{}{}
	}

	return arr
}
