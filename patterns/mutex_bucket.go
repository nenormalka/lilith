package patterns

import "sync"

type (
	MutexBucket[T any] struct {
		mu     [SegmentCount]sync.RWMutex
		bucket [SegmentCount]T
	}

	MutexBucketFunc[T any] func(p T) error
)

const (
	SegmentCount    = 256
	segmentAndOpVal = SegmentCount - 1
)

func NewMutexBucket[T any](elems [SegmentCount]T) *MutexBucket[T] {
	return &MutexBucket[T]{
		mu:     [SegmentCount]sync.RWMutex{},
		bucket: elems,
	}
}

func (mb *MutexBucket[T]) UnderLock(identifier int64, f MutexBucketFunc[T]) error {
	segment := mb.getSegment(identifier)

	mb.mu[segment].Lock()
	defer mb.mu[segment].Unlock()

	return f(mb.bucket[segment])
}

func (mb *MutexBucket[T]) UnderReadLock(identifier int64, f MutexBucketFunc[T]) error {
	segment := mb.getSegment(identifier)

	mb.mu[segment].RLock()
	defer mb.mu[segment].RUnlock()

	return f(mb.bucket[segment])
}

func (mb *MutexBucket[T]) getSegment(identifier int64) int64 {
	return identifier & segmentAndOpVal
}
