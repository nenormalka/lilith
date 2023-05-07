package methods

type (
	ArrayFilterFunc[T any] func(elem T) bool
)

func ArrayFilter[S ~[]T, T any](f ArrayFilterFunc[T], s S) S {
	if len(s) == 0 {
		return nil
	}

	var arr S

	for i := range s {
		if f(s[i]) {
			arr = append(arr, s[i])
		}
	}

	return arr
}
