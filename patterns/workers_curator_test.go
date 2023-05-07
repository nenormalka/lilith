package patterns

import (
	"context"
	"reflect"
	"testing"
)

func TestWorkersCurator(t *testing.T) {
	want := map[int]int{
		1: 2,
		2: 4,
		3: 6,
		4: 8,
		5: 10,
	}

	source := make(chan int)

	got := make(map[int]int, len(want))

	curator := NewWorkersCurator[int](context.Background(), func(data int) error {
		got[data] = data * 2
		return nil
	}, source, 2)

	go func() {
		defer close(source)

		for key := range want {
			source <- key
		}
	}()

	curator.Wait()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("NewWorkersCurator() = %v, want %v", got, want)
	}
}
