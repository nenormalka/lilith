package patterns

func Split[T any](source <-chan T, n int) []<-chan T {
	dests := make([]<-chan T, n)

	for i := 0; i < n; i++ {
		ch := make(chan T)
		dests = append(dests, ch)

		go func() {
			defer close(ch)
			for val := range source {
				ch <- val
			}
		}()
	}

	return dests
}
