package patterns

import (
	"errors"
	"math"
	"sync"
	"time"
)

const (
	MaxPriority = math.MaxInt64
	MinPriority = math.MinInt64
)

var (
	ErrKeyDoesNotExist = errors.New("key does not exist")
	ErrEmptyQueue      = errors.New("queue is empty")
	ErrEmptyChannel    = errors.New("channel is empty")
	ErrEmptyReqTime    = errors.New("request time is empty")
)

type (
	NodeInfo[T any, M comparable] struct {
		Key      M
		Value    T
		Priority int64
	}

	Queue[T any, M comparable] struct {
		head    *node[T, M]
		m       map[M]*node[T, M]
		mu      sync.Mutex
		ch      NodeInfoChannel[T, M]
		closeCh chan struct{}
	}

	NodeInfoChannel[T any, M comparable] chan NodeInfo[T, M]

	QueueOption[T any, M comparable] func(q *Queue[T, M])

	node[T any, M comparable] struct {
		info NodeInfo[T, M]
		next *node[T, M]
		prev *node[T, M]
	}
)

func WithChannel[T any, M comparable]() QueueOption[T, M] {
	return func(q *Queue[T, M]) {
		q.ch = make(NodeInfoChannel[T, M])
		q.closeCh = make(chan struct{})
	}
}

// NewQueue creates a new queue with the given options.
// Use one of the methods Pop/PopByKey or PopByChannelWithReqTime to get the value from the queue.
func NewQueue[T any, M comparable](opts ...QueueOption[T, M]) *Queue[T, M] {
	q := &Queue[T, M]{
		m:  make(map[M]*node[T, M]),
		mu: sync.Mutex{},

		head:    nil,
		ch:      nil,
		closeCh: nil,
	}

	for _, opt := range opts {
		opt(q)
	}

	return q
}

func (q *Queue[T, M]) Push(
	priority int64,
	key M,
	value T,
) {
	q.mu.Lock()
	defer q.mu.Unlock()

	existsNode, ok := q.m[key]
	if ok {
		q.removeNode(existsNode)
	}

	q.setNode(&node[T, M]{
		info: NodeInfo[T, M]{
			Key:      key,
			Value:    value,
			Priority: priority,
		},
	})
}

func (q *Queue[T, M]) Pop() (NodeInfo[T, M], error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.head == nil {
		return NodeInfo[T, M]{}, ErrEmptyQueue
	}

	n := q.head
	q.removeNode(n)

	return n.info, nil
}

func (q *Queue[T, M]) PopByKey(key M) (NodeInfo[T, M], error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	existsNode, ok := q.m[key]
	if ok {
		q.removeNode(existsNode)
		return existsNode.info, nil
	}

	return NodeInfo[T, M]{}, ErrKeyDoesNotExist
}

func (q *Queue[T, M]) PopByChannelWithReqTime(reqTime time.Duration) (<-chan NodeInfo[T, M], error) {
	if q.ch == nil {
		return nil, ErrEmptyChannel
	}

	if reqTime == 0 {
		return nil, ErrEmptyReqTime
	}

	go func() {
		ticker := time.NewTicker(reqTime)
		defer ticker.Stop()

		for {
			<-ticker.C

			info, err := q.Pop()
			if err != nil {
				continue
			}

			select {
			case q.ch <- info:
			case <-q.closeCh:
				return
			default:
				q.Push(info.Priority, info.Key, info.Value)
			}
		}
	}()

	return q.ch, nil
}

func (q *Queue[T, M]) Size() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	return len(q.m)
}

func (q *Queue[T, M]) Close() {
	if q.closeCh != nil {
		close(q.closeCh)
	}

	if q.ch != nil {
		close(q.ch)
	}
}

func (q *Queue[T, M]) removeNode(n *node[T, M]) {
	if n.prev == nil {
		q.head = n.next
	} else {
		n.prev.next = n.next
	}

	if n.next != nil {
		n.next.prev = n.prev
	}

	delete(q.m, n.info.Key)
}

func (q *Queue[T, M]) setNode(n *node[T, M]) {
	if q.head == nil {
		q.head = n
		q.m[n.info.Key] = n
		return
	}

	if q.head.info.Priority < n.info.Priority {
		n.next = q.head
		q.head.prev = n
		q.head = n
		q.m[n.info.Key] = n
		return
	}

	current := q.head
	for current.next != nil && current.next.info.Priority > n.info.Priority {
		current = current.next
	}

	n.next = current.next
	if current.next != nil {
		current.next.prev = n
	}

	current.next = n
	n.prev = current
	q.m[n.info.Key] = n
}
