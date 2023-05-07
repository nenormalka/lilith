package patterns

import "sync"

type (
	Job[T any] func(in, out chan T)
)

func Pipeline[T any](jobs ...Job[T]) {
	inCh := make(chan T)
	outCh := make(chan T)

	wg := sync.WaitGroup{}
	wg.Add(len(jobs))

	for _, j := range jobs {
		go func(j Job[T], in, out chan T) {
			j(in, out)
			close(out)
			wg.Done()
		}(j, inCh, outCh)

		inCh = outCh
		outCh = make(chan T)
	}

	wg.Wait()
}
