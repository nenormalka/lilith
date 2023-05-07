package patterns

import (
	"reflect"
	"testing"
)

func TestPipeline(t *testing.T) {
	var z int

	jobs := []Job[int]{
		func(_, out chan int) {
			out <- 1
		},
		func(in, out chan int) {
			out <- <-in + 1
		},
		func(in, out chan int) {
			out <- <-in * 2
		},
		func(in, _ chan int) {
			z = <-in - 1
		},
	}

	Pipeline[int](jobs...)

	if !reflect.DeepEqual(z, 3) {
		t.Errorf("pipeline = %v, want %v", z, 3)
	}
}
