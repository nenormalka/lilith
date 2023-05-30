package patterns

import (
	"context"
	"reflect"
	"sync"
	"testing"
)

const (
	countBroadcastTest = 3
)

func TestBroadcast(t *testing.T) {
	source := make(chan int)
	b := NewBroadcast[int](context.Background(), source)
	mu := sync.Mutex{}
	m := make(map[int]int)
	wg := sync.WaitGroup{}

	wg.Add(countBroadcastTest)

	for i := 0; i < countBroadcastTest; i++ {
		ch := b.Subscribe()

		go func(ch <-chan int) {
			defer wg.Done()
			for data := range ch {
				mu.Lock()
				m[data]++
				mu.Unlock()
			}
		}(ch)
	}

	for _, arg := range []int{1, 2, 3, 4, 5} {
		source <- arg
	}

	close(source)

	wg.Wait()

	want := map[int]int{
		1: countBroadcastTest,
		2: countBroadcastTest,
		3: countBroadcastTest,
		4: countBroadcastTest,
		5: countBroadcastTest,
	}

	if !reflect.DeepEqual(m, want) {
		t.Errorf("broadcast = %v, want %v", m, want)
	}
}
