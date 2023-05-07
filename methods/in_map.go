package methods

func InMap[M ~map[K]V, K, V comparable](m M, needle V) bool {
	for _, value := range m {
		if value == needle {
			return true
		}
	}

	return false
}
