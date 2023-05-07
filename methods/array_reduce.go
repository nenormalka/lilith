package methods

type (
	ArrayReduceFunc[T any] func(first, second T) T
)

func ArrayReduce[S ~[]T, T any](f ArrayReduceFunc[T], s S, init T) T {
	accumulator := init

	for i := range s {
		accumulator = f(accumulator, s[i])
	}

	return accumulator
}
