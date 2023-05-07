package methods

func Ternary[T any](cond bool, first, second T) T {
	if cond {
		return first
	}

	return second
}
