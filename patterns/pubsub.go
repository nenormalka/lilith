package patterns

import (
	"context"
	"errors"
	"sync"
)

type (
	Broker[T any] struct {
		mu     sync.Mutex
		topics map[string]*Topic[T]
		closed bool
	}

	Topic[T any] struct {
		source chan T
		bc     *Broadcast[T]
	}
)

var (
	ErrBrokerClosed       = errors.New("broker closed")
	ErrTopicNotFound      = errors.New("topic not found")
	ErrPublishFailed      = errors.New("publish failed")
	ErrTopicAlreadyExists = errors.New("topic already exists")
)

func NewBroker[T any]() *Broker[T] {
	return &Broker[T]{
		mu:     sync.Mutex{},
		topics: map[string]*Topic[T]{},
		closed: false,
	}
}

func (b *Broker[T]) CreateTopic(ctx context.Context, topic string) error {
	return b.process(func() error {
		if _, ok := b.topics[topic]; ok {
			return ErrTopicAlreadyExists
		}

		source := make(chan T)

		b.topics[topic] = &Topic[T]{
			source: source,
			bc:     NewBroadcast[T](ctx, source),
		}

		return nil
	})
}

func (b *Broker[T]) Publish(ctx context.Context, topic string, message T) error {
	return b.processTopic(topic, func(t *Topic[T]) error {
		select {
		case <-ctx.Done():
			return ErrPublishFailed
		case t.source <- message:
			return nil
		}
	})
}

func (b *Broker[T]) Subscribe(topic string) (<-chan T, error) {
	var ch <-chan T

	err := b.processTopic(topic, func(t *Topic[T]) error {
		ch = t.bc.Subscribe()

		return nil
	})

	return ch, err
}

func (b *Broker[T]) Unsubscribe(topic string, ch <-chan T) error {
	return b.processTopic(topic, func(t *Topic[T]) error {
		t.bc.Unsubscribe(ch)

		return nil
	})
}

func (b *Broker[T]) CloseTopic(topic string) error {
	return b.processTopic(topic, func(t *Topic[T]) error {
		close(t.source)
		delete(b.topics, topic)

		return nil
	})
}

func (b *Broker[T]) CloseAll() error {
	return b.process(func() error {
		b.closed = true

		for _, t := range b.topics {
			close(t.source)
		}

		b.topics = nil

		return nil
	})
}

func (b *Broker[T]) SubscribersCountByTopic(topic string) (int, error) {
	var count int

	err := b.processTopic(topic, func(t *Topic[T]) error {
		count = t.bc.SubscribersCount()

		return nil
	})

	return count, err
}

func (b *Broker[T]) SubscribersCountAll() (int, error) {
	var count int

	err := b.process(func() error {
		for _, t := range b.topics {
			count += t.bc.SubscribersCount()
		}

		return nil
	})

	return count, err
}

func (b *Broker[T]) process(f func() error) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.closed {
		return ErrBrokerClosed
	}

	return f()
}

func (b *Broker[T]) processTopic(topic string, f func(t *Topic[T]) error) error {
	return b.process(func() error {
		t, ok := b.topics[topic]
		if !ok {
			return ErrTopicNotFound
		}

		return f(t)
	})
}
