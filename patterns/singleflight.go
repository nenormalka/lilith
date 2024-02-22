package patterns

import (
	"context"
	"errors"

	"golang.org/x/sync/singleflight"
)

var (
	ErrEmptySingleFlight = errors.New("empty singleflight")
	ErrFailedTypeAssert  = errors.New("failed type assert")
)

type (
	Result[T any] struct {
		Val    T
		Err    error
		Shared bool
	}
)

func Do[T any](
	sf *singleflight.Group,
	key string,
	fn func() (any, error),
) (v T, err error, shared bool) {
	if sf == nil {
		return v, ErrEmptySingleFlight, shared
	}

	result, err, shared := sf.Do(key, fn)

	if err != nil {
		return v, err, shared
	}

	val, ok := result.(T)
	if !ok {
		return v, ErrFailedTypeAssert, shared
	}

	return val, nil, shared
}

func DoChan[T any](
	sf *singleflight.Group,
	key string,
	fn func() (any, error),
) (<-chan Result[T], error) {
	return doChan(
		sf,
		key,
		fn,
		func(ch chan Result[T], chSF <-chan singleflight.Result) {
			defer close(ch)

			result := <-chSF
			val, ok := result.Val.(T)
			if !ok {
				ch <- Result[T]{Err: ErrFailedTypeAssert, Shared: result.Shared}
			} else {
				ch <- Result[T]{Val: val, Err: result.Err, Shared: result.Shared}
			}
		},
	)
}

func DoChanCtx[T any](
	ctx context.Context,
	sf *singleflight.Group,
	key string,
	fn func() (any, error),
) (<-chan Result[T], error) {
	return doChan(
		sf,
		key,
		fn,
		func(ch chan Result[T], chSF <-chan singleflight.Result) {
			defer close(ch)

			select {
			case <-ctx.Done():
				ch <- Result[T]{Err: ctx.Err()}
				return
			case result := <-chSF:
				val, ok := result.Val.(T)
				if !ok {
					ch <- Result[T]{Err: ErrFailedTypeAssert, Shared: result.Shared}
				} else {
					ch <- Result[T]{Val: val, Err: result.Err, Shared: result.Shared}
				}
			}
		},
	)
}

func doChan[T any](
	sf *singleflight.Group,
	key string,
	fn func() (any, error),
	convertFn func(chan Result[T], <-chan singleflight.Result),
) (<-chan Result[T], error) {
	if sf == nil {
		return nil, ErrEmptySingleFlight
	}

	ch := make(chan Result[T])

	go convertFn(ch, sf.DoChan(key, fn))

	return ch, nil
}

func Forget(sf *singleflight.Group, key string) error {
	if sf == nil {
		return ErrEmptySingleFlight
	}

	sf.Forget(key)

	return nil
}
