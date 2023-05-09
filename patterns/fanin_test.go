package patterns

import (
	"reflect"
	"testing"
)

func TestFanIn(t *testing.T) {
	for name, tt := range map[string]struct {
		times int
		want  map[int]int
	}{
		"#1": {
			times: 3,
			want: map[int]int{
				1: 3,
				2: 3,
				3: 3,
			},
		},
		"#2": {
			times: 5,
			want: map[int]int{
				1: 5,
				2: 5,
				3: 5,
				4: 5,
				5: 5,
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			sources := make([]<-chan int, 0)

			for i := 0; i < tt.times; i++ {
				ch := make(chan int)
				sources = append(sources, ch)

				go func() {
					defer close(ch)
					for j := 1; j <= tt.times; j++ {
						ch <- j
					}
				}()
			}

			dest := FanIn[int](sources...)
			got := make(map[int]int)

			for d := range dest {
				got[d]++
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FanIn() = %v, want %v", got, tt.want)
			}
		})
	}
}
