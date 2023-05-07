package methods

func ArrayShift[S ~[]T, T any](s *S) T {
	var f T
	if len(*s) == 0 {
		return f
	}

	f = (*s)[0]
	*s = (*s)[1:]

	return f
}
