package methods

func ArrayInsert[S ~[]T, T any](s *S, indx int, elems ...T) {
	if len(*s) < indx {
		ArrayPush[S, T](s, elems...)

		return
	}

	*s = append((*s)[:indx], append(elems, (*s)[indx:]...)...)
}
