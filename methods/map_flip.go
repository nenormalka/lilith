package methods

func MapFlip[M ~map[K]V, R ~map[V]K, K, V comparable](m M) R {
	r := make(R, len(m))

	for key, value := range m {
		r[value] = key
	}

	return r
}
