package collections

type (
	Predicate[T comparable] func(elem T) bool

	Iterator[T comparable] interface {
		HasNext() bool
		GetNext() T
	}

	Collection[T comparable] interface {
		Add(elem T) bool
		AddAll(elems Collection[T]) bool
		Clear()
		Contains(elem T) bool
		ContainsAll(elems Collection[T]) bool
		Equals(elem Collection[T]) bool
		IsEmpty() bool
		Remove(elem T) bool
		RemoveAll(elems Collection[T]) bool
		RetainAll(elems Collection[T]) bool
		Size() int
		ToArray() []T
		RemoveIf(f Predicate[T])
	}
)
