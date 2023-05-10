package methods

func ArrayInsert[S ~[]T, T any](s *S, indx int, elems ...T) {
	if len(*s) < indx {
		ArrayPush[S, T](s, elems...)

		return
	}

	elems = append(elems, (*s)[indx:]...)

	*s = append((*s)[:indx], elems...)
}
