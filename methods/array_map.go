package methods

type (
	ArrayMapFunc[T any] func(elem T) T
)

func ArrayMap[S ~[]T, T any](f ArrayMapFunc[T], s S) S {
	if len(s) == 0 {
		return nil
	}

	var arr S

	for i := range s {
		arr = append(arr, f(s[i]))
	}

	return arr
}
