package patterns

import "context"

type (
	WrappedFunc func() error
)

func Wrap(ctx context.Context, f WrappedFunc) error {
	errCh := make(chan error)

	go func() {
		defer close(errCh)
		errCh <- f()
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errCh:
		return err
	}
}
