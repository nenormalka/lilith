package methods

import (
	"errors"
	"math"
)

var (
	ErrArrayChunkSize       = errors.New("size must be greater than 0")
	ErrArrayChunkEmptySlice = errors.New("input slice must not be empty")
)

func ArrayChunk[S ~[]T, T any](s S, size int) ([]S, error) {
	if size < 1 {
		return nil, ErrArrayChunkSize
	}

	if len(s) == 0 {
		return nil, ErrArrayChunkEmptySlice
	}

	length := len(s)
	if size >= length {
		return []S{s}, nil
	}

	chunks := int(math.Ceil(float64(length) / float64(size)))

	n := make([]S, 0, chunks)
	for i, end := 0, 0; chunks > 0; chunks-- {
		end = (i + 1) * size
		if end > length {
			end = length
		}

		arr := make(S, end-i*size)
		copy(arr, s[i*size:end])
		n = append(n, arr)
		i++
	}

	return n, nil
}
