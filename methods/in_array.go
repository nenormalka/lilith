package methods

func InArray[S ~[]T, T comparable](s S, needle T) bool {
	for i := range s {
		if s[i] == needle {
			return true
		}
	}

	return false
}
