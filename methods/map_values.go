package methods

func MapValues[M ~map[K]V, K comparable, V any](m M) []V {
	vals := make([]V, 0, len(m))

	for _, val := range m {
		vals = append(vals, val)
	}

	return vals
}
