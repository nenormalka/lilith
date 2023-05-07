package patterns

type (
	Iterator[T any] struct {
		index int
		data  []T
	}
)

func NewIterator[T any](data []T) *Iterator[T] {
	return &Iterator[T]{
		data:  data,
		index: 0,
	}
}

func (i *Iterator[T]) HasNext() bool {
	if i.index < len(i.data) {
		return true
	}
	return false
}

func (i *Iterator[T]) GetNext() T {
	var data T
	if i.HasNext() {
		data = i.data[i.index]
		i.index++

		return data
	}

	return data
}
