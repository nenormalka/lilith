package patterns

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"golang.org/x/sync/singleflight"
)

func TestDoChan(t *testing.T) {
	g := &singleflight.Group{}
	block := make(chan struct{})

	resp1, err := DoChan[int](g, "key", func() (any, error) {
		<-block
		return 1, nil
	})

	if err != nil {
		t.Errorf("DoChan() = %v, want %v", err, nil)
	}

	resp2, err := DoChan[int](g, "key", func() (any, error) {
		<-block
		return 2, nil
	})

	if err != nil {
		t.Errorf("DoChan() = %v, want %v", err, nil)
	}

	close(block)

	res1 := <-resp1
	res2 := <-resp2

	if res1.Val != 1 {
		t.Errorf("DoChan() = %v, want %v", res1.Val, 1)
	}

	if res2.Val != 1 {
		t.Errorf("DoChan() = %v, want %v", res2.Val, 1)
	}
}

func TestDoChanCtx(t *testing.T) {
	g := &singleflight.Group{}
	ctx, cancel := context.WithCancel(context.Background())

	resp1, err := DoChanCtx[int](ctx, g, "key", func() (any, error) {
		time.Sleep(5 * time.Second)
		return 1, nil
	})

	if err != nil {
		t.Errorf("DoChanCtx() = %v, want %v", err, nil)
	}

	resp2, err := DoChanCtx[int](ctx, g, "key", func() (any, error) {
		time.Sleep(5 * time.Second)
		return 2, nil
	})

	if err != nil {
		t.Errorf("DoChanCtx() = %v, want %v", err, nil)
	}

	cancel()

	res1 := <-resp1
	res2 := <-resp2

	if !errors.Is(res1.Err, context.Canceled) {
		t.Errorf("DoChanCtx() = %v, want %v", res1.Err, context.Canceled)
	}

	if !errors.Is(res2.Err, context.Canceled) {
		t.Errorf("DoChanCtx() = %v, want %v", res2.Err, context.Canceled)
	}
}

func TestDo(t *testing.T) {
	var (
		g     = &singleflight.Group{}
		block = make(chan struct{})
		wg    = &sync.WaitGroup{}

		resp1   int
		err1    error
		shared1 bool
		resp2   int
		err2    error
		shared2 bool
	)

	wg.Add(2)

	go func() {
		resp1, err1, shared1 = Do[int](g, "key", func() (any, error) {
			<-block
			return 1, nil
		})
		wg.Done()
	}()

	go func() {
		resp2, err2, shared2 = Do[int](g, "key", func() (any, error) {
			<-block
			return 2, nil
		})
		wg.Done()
	}()

	time.Sleep(1 * time.Second)

	close(block)
	wg.Wait()

	if resp1 != 1 {
		t.Errorf("Do() = %v, want %v", resp1, 1)
	}

	if resp2 != 1 {
		t.Errorf("Do() = %v, want %v", resp2, 1)
	}

	if err1 != nil {
		t.Errorf("Do() = %v, want %v", err1, nil)
	}

	if err2 != nil {
		t.Errorf("Do() = %v, want %v", err2, nil)
	}

	if !shared1 {
		t.Errorf("Do() = %v, want %v", shared1, true)
	}

	if !shared2 {
		t.Errorf("Do() = %v, want %v", shared2, true)
	}
}
