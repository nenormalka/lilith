package methods

func ArrayMerge[S ~[]T, T any](arrs ...S) S {
	if len(arrs) == 0 {
		return nil
	}

	n := 0
	for _, v := range arrs {
		n += len(v)
	}

	res := make([]T, 0, n)

	for _, v := range arrs {
		res = append(res, v...)
	}

	return res
}
