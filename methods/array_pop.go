package methods

func ArrayPop[S ~[]T, T any](s *S) T {
	var e T
	if len(*s) == 0 {
		return e
	}

	ep := len(*s) - 1
	e = (*s)[ep]
	*s = (*s)[:ep]

	return e
}
