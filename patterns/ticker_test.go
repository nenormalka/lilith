package patterns

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	i := 1

	wg.Add(1)

	Ticker(ctx, 1*time.Second, func() {
		fmt.Println("tick", i)

		i++

		if i == 4 {
			wg.Done()
		}
	})

	wg.Wait()
	cancel()

	time.Sleep(2 * time.Second)

	if i != 4 {
		t.Errorf("Ticker() = %v, want %v", i, 4)
	}
}

func TestTickerV2(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	i := 1

	wg.Add(1)

	TickerV2(ctx, 1*time.Second, func() {
		fmt.Println("tick", i)

		i++

		if i == 4 {
			wg.Done()
		}
	})

	wg.Wait()
	cancel()

	time.Sleep(2 * time.Second)

	if i != 4 {
		t.Errorf("Ticker() = %v, want %v", i, 4)
	}
}
