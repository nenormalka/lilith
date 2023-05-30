package patterns

import "context"

type (
	Broadcast[T any] struct {
		source    <-chan T
		listeners []chan T
	}
)

func NewBroadcast[T any](ctx context.Context, source <-chan T) *Broadcast[T] {
	b := &Broadcast[T]{
		source: source,
	}

	go b.serve(ctx)

	return b
}

func (b *Broadcast[T]) Subscribe() <-chan T {
	newListener := make(chan T)
	b.listeners = append(b.listeners, newListener)

	return newListener
}

func (b *Broadcast[T]) Unsubscribe(channel <-chan T) {
	for i, ch := range b.listeners {
		if ch == channel {
			b.listeners = append(b.listeners[:i], b.listeners[i+1:]...)
			close(ch)
			break
		}
	}
}

func (b *Broadcast[T]) SubscribersCount() int {
	return len(b.listeners)
}

func (b *Broadcast[T]) serve(ctx context.Context) {
	defer func() {
		for _, listener := range b.listeners {
			if listener != nil {
				close(listener)
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case val, ok := <-b.source:
			if !ok {
				return
			}

			for _, listener := range b.listeners {
				if listener == nil {
					continue
				}

				select {
				case listener <- val:
				case <-ctx.Done():
					return
				default:

				}
			}
		}
	}
}
