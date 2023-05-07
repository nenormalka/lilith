package methods

func ArrayPad[S ~[]T, T any](s S, size int, val T) S {
	if size <= len(s) {
		return s
	}

	n := size - len(s)
	tmp := make([]T, n)

	for i := 0; i < n; i++ {
		tmp[i] = val
	}

	return append(s, tmp...)
}
