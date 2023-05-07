package patterns

import "context"

type (
	FilterFunc[T any] func(data T) bool
)

func FilterChannel[T any](ctx context.Context, inputStream <-chan T, f FilterFunc[T]) <-chan T {
	filteredStream := make(chan T)
	go func() {
		defer close(filteredStream)

		for {
			select {
			case <-ctx.Done():
				return
			case i, ok := <-inputStream:
				if !ok {
					return
				}

				if !f(i) {
					break
				}

				select {
				case <-ctx.Done():
					return
				case filteredStream <- i:
				}
			}
		}
	}()

	return filteredStream
}
