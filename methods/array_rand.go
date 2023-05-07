package methods

import (
	"math/rand"
	"time"
)

func ArrayRand[S ~[]T, T any](elements S) S {
	l := len(elements)

	if l == 0 {
		return nil
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := make([]T, l)

	for i, v := range r.Perm(l) {
		n[i] = elements[v]
	}

	return n
}
