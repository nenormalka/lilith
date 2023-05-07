package methods

func ArrayFill[S ~[]T, T any](startIndex int, num int, value T) S {
	if num == 0 {
		return nil
	}

	l := startIndex + num
	m := make([]T, l)

	for i := startIndex; i < l; i++ {
		m[i] = value
	}

	return m
}
