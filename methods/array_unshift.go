package methods

func ArrayUnshift[S ~[]T, T any](s *S, elements ...T) {
	*s = append(elements, *s...)
}
