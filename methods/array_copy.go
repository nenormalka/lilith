package methods

func ArrayCopy[S ~[]T, T any](s S) S {
	return append([]T{}, s...)
}
