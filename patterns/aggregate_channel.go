package patterns

import "context"

type (
	AggregateFunc[T any] func(data T) T
)

func AggregateChannel[T any](ctx context.Context, inputStream <-chan T, f AggregateFunc[T]) <-chan T {
	aggregateStream := make(chan T)

	go func() {
		defer close(aggregateStream)
		for {
			select {
			case <-ctx.Done():
				return
			case i, ok := <-inputStream:
				if !ok {
					return
				}

				select {
				case <-ctx.Done():
					return
				case aggregateStream <- f(i):
				}
			}
		}
	}()

	return aggregateStream
}
