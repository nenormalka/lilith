package patterns

import (
	"context"
	"reflect"
	"sync"
	"testing"
)

func TestWorkers(t *testing.T) {
	mu := sync.Mutex{}
	m := make(map[int]struct{})

	for name, tt := range map[string]struct {
		args  []int
		want  map[int]struct{}
		count int
		f     WorkFunc[int]
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
			f: func(data int) error {
				mu.Lock()
				m[data] = struct{}{}
				mu.Unlock()

				return nil
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			source := make(chan int)
			wg := &sync.WaitGroup{}
			Workers(context.Background(), tt.f, tt.count, wg, source, nil, nil)

			go func() {
				defer close(source)

				for _, arg := range tt.args {
					source <- arg
				}
			}()

			wg.Wait()

			if !reflect.DeepEqual(m, tt.want) {
				t.Errorf("Workers() = %v, want %v", m, tt.want)
			}

			m = make(map[int]struct{})
		})
	}
}
