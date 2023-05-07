package methods

func ArrayPush[S ~[]T, T any](s *S, elements ...T) {
	*s = append(*s, elements...)
}
