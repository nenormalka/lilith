package patterns

import (
	"context"
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestPubSub(t *testing.T) {
	ctx := context.Background()

	broker := NewBroker[int]()
	ch, err := broker.Subscribe("test")
	if err != ErrTopicNotFound {
		t.Errorf("pubsub Subscribe() error = %v, want %v", err, ErrTopicNotFound)
	}

	if ch != nil {
		t.Errorf("pubsub Subscribe() ch = %v, want %v", ch, nil)
	}

	err = broker.Publish(ctx, "test", 1)
	if err != ErrTopicNotFound {
		t.Errorf("pubsub Publish() error = %v, want %v", err, ErrTopicNotFound)
	}

	err = broker.CreateTopic(ctx, "test 1")
	if err != nil {
		t.Errorf("pubsub CreateTopic() error = %v, want %v", err, nil)
	}

	err = broker.CreateTopic(ctx, "test 2")
	if err != nil {
		t.Errorf("pubsub CreateTopic() error = %v, want %v", err, nil)
	}

	err = broker.CreateTopic(ctx, "test 1")
	if err != ErrTopicAlreadyExists {
		t.Errorf("pubsub CreateTopic() error = %v, want %v", err, ErrTopicAlreadyExists)
	}

	mu := sync.Mutex{}
	m := make(map[int]int)
	wg := sync.WaitGroup{}
	wg.Add(5)

	lastCh := ch

	for i := 0; i < 3; i++ {
		ch, err = broker.Subscribe("test 1")
		if err != nil {
			t.Errorf("pubsub Subscribe() error = %v, want %v", err, nil)
		}

		go func(ch <-chan int) {
			defer wg.Done()

			for data := range ch {
				mu.Lock()
				m[data]++
				mu.Unlock()
			}
		}(ch)

		lastCh = ch
	}

	for i := 0; i < 2; i++ {
		ch, err = broker.Subscribe("test 2")
		if err != nil {
			t.Errorf("pubsub Subscribe() error = %v, want %v", err, nil)
		}

		go func(ch <-chan int) {
			defer wg.Done()

			for data := range ch {
				mu.Lock()
				m[data]++
				mu.Unlock()
			}
		}(ch)
	}

	err = broker.Publish(ctx, "test 1", 1)
	if err != nil {
		t.Errorf("pubsub Publish() error = %v, want %v", err, nil)
	}

	err = broker.Publish(ctx, "test 1", 2)
	if err != nil {
		t.Errorf("pubsub Publish() error = %v, want %v", err, nil)
	}

	err = broker.Publish(ctx, "test 2", 3)
	if err != nil {
		t.Errorf("pubsub Publish() error = %v, want %v", err, nil)
	}

	count, err := broker.SubscribersCountAll()
	if err != nil {
		t.Errorf("pubsub SubscribersCountAll() error = %v, want %v", err, nil)
	}

	if count != 5 {
		t.Errorf("pubsub SubscribersCountAll() = %v, want %v", count, 5)
	}

	count, err = broker.SubscribersCountByTopic("test 1")
	if err != nil {
		t.Errorf("pubsub SubscribersCountByTopic() error = %v, want %v", err, nil)
	}

	if count != 3 {
		t.Errorf("pubsub SubscribersCountByTopic() = %v, want %v", count, 3)
	}

	time.Sleep(100 * time.Millisecond)

	err = broker.Unsubscribe("test 1", lastCh)
	if err != nil {
		t.Errorf("pubsub Unsubscribe() error = %v, want %v", err, nil)
	}

	count, err = broker.SubscribersCountByTopic("test 1")
	if err != nil {
		t.Errorf("pubsub SubscribersCountByTopic() error = %v, want %v", err, nil)
	}

	if count != 2 {
		t.Errorf("pubsub SubscribersCountByTopic() = %v, want %v", count, 2)
	}

	err = broker.CloseTopic("test 1")
	if err != nil {
		t.Errorf("pubsub CloseTopic() error = %v, want %v", err, nil)
	}

	count, err = broker.SubscribersCountAll()
	if err != nil {
		t.Errorf("pubsub SubscribersCountAll() error = %v, want %v", err, nil)
	}

	if count != 2 {
		t.Errorf("pubsub SubscribersCountAll() = %v, want %v", count, 2)
	}

	err = broker.CloseAll()
	if err != nil {
		t.Errorf("pubsub CloseAll() error = %v, want %v", err, nil)
	}

	wg.Wait()
	want := map[int]int{
		1: 3,
		2: 3,
		3: 2,
	}

	if !reflect.DeepEqual(m, want) {
		t.Errorf("broadcast = %v, want %v", m, want)
	}

	err = broker.CloseAll()
	if err != ErrBrokerClosed {
		t.Errorf("pubsub CloseAll() error = %v, want %v", err, nil)
	}
}
