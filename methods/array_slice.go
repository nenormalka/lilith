package methods

import "errors"

var (
	ErrArraySliceOffset = errors.New("offset must be less than length array")
)

func ArraySlice[S ~[]T, T any](s S, offset, length int) (S, error) {
	l := len(s)

	if l == 0 {
		return nil, nil
	}

	if offset > l {
		return nil, ErrArraySliceOffset
	}

	end := offset + length

	if end < l {
		return s[offset:end], nil
	}

	return s[offset:], nil
}
