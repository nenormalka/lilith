package collections

type (
	Comparator[T comparable] func(first, second T) bool

	List[T comparable] interface {
		Collection[T]

		AddList(indx int, elem T) bool
		AddAllList(indx int, elems []T) bool
		Get(indx int) (T, error)
		IndexOf(elem T) int
		LastIndexOf(elem T) int
		ListIterator() Iterator[T]
		ListIteratorIndx(indx int) Iterator[T]
		RemoveList(indx int) (T, error)
		Set(indx int, elem T) (T, error)
		Sort(f Comparator[T])
		SubList(start, end int) ([]T, error)
	}
)
