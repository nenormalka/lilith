package methods

func InMapKey[M ~map[K]V, K, V comparable](m M, needle K) bool {
	_, ok := m[needle]
	return ok
}
