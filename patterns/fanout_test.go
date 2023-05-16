package patterns

import (
	"reflect"
	"sync"
	"testing"
)

func TestFanOut(t *testing.T) {
	for name, tt := range map[string]struct {
		args  []int
		want  map[int]struct{}
		count int
	}{
		"#1": {
			args: []int{1, 2, 3, 4, 5},
			want: map[int]struct{}{
				1: {},
				2: {},
				3: {},
				4: {},
				5: {},
			},
			count: 3,
		},
	} {
		t.Run(name, func(t *testing.T) {
			source := make(chan int)
			dests := FanOut(source, tt.count)

			got := make(map[int]struct{})
			m := sync.Mutex{}
			wg := sync.WaitGroup{}
			wg.Add(tt.count)

			go func() {
				for _, arg := range tt.args {
					source <- arg
				}
				close(source)
			}()

			for _, dest := range dests {
				go func(d <-chan int) {
					defer wg.Done()
					for n := range d {
						m.Lock()
						got[n] = struct{}{}
						m.Unlock()
					}
				}(dest)
			}

			wg.Wait()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FanOut() = %v, want %v", got, tt.want)
			}
		})
	}
}
