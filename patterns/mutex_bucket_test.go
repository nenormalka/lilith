package patterns

import (
	"sync"
	"testing"
)

func TestMutexBucket(t *testing.T) {
	type test struct {
		count int
	}

	mb := NewMutexBucket[*test](func() [SegmentCount]*test {
		var elems [SegmentCount]*test
		for i := 0; i < SegmentCount; i++ {
			elems[i] = &test{count: 0}
		}
		return elems
	}())

	wg := sync.WaitGroup{}
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			if err := mb.UnderLock(5, func(p *test) error {
				p.count++
				return nil
			}); err != nil {
				t.Errorf("pubsub UnderLock() error = %v, want %v", err, nil)
			}

			wg.Done()
		}()
	}

	wg.Wait()

	count := 0

	if err := mb.UnderReadLock(5, func(p *test) error {
		count = p.count
		return nil
	}); err != nil {
		t.Errorf("pubsub UnderReadLock() error = %v, want %v", err, nil)
	}

	if count != 100 {
		t.Errorf("pubsub UnderReadLock() count = %v, want %v", count, 100)
	}
}
