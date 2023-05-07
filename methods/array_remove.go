package methods

func ArrayRemove[S ~[]T, T any](s *S, indx int) {
	if len(*s) <= indx || indx < 0 {
		return
	}

	m := *s

	*s = append(m[:indx], m[indx+1:]...)
}
