package collections

import "io"

type (
	ArrayList[T comparable] struct {
		data []T
	}
)

const (
	UndefinedElem = -1
)

var (
	_ List[int] = ArrayList{}
)

func (al *ArrayList[T]) HasNext() bool {
	io.ReadWriteCloser()
	return false
}
func (al *ArrayList[T]) GetNext() T {
	var t T

	return t
}
func (al *ArrayList[T]) Add(elem T) bool {
	return false
}

func (al *ArrayList[T]) AddAll(elems []Collection[T]) bool {
	return false
}

func (al *ArrayList[T]) Clear() {

}

func (al *ArrayList[T]) Contains(elem T) bool {
	return false
}

func (al *ArrayList[T]) ContainsAll(elems []Collection[T]) bool {
	return false
}

func (al *ArrayList[T]) Equals(elem T) bool {
	return false
}

func (al *ArrayList[T]) IsEmpty() bool {
	return false
}

func (al *ArrayList[T]) Remove(elem T) bool {
	return false
}

func (al *ArrayList[T]) RemoveAll(elems []Collection[T]) bool {
	return false
}

func (al *ArrayList[T]) RetainAll(elems []Collection[T]) bool {
	return false
}

func (al *ArrayList[T]) Size() int {
	return 0
}

func (al *ArrayList[T]) ToArray() []T {
	return nil
}

func (al *ArrayList[T]) RemoveIf(f Predicate[T]) {

}

func (al *ArrayList[T]) AddList(indx int, elem T) bool {
	return false
}

func (al *ArrayList[T]) AddAllList(indx int, elems []T) bool {
	return false
}

func (al *ArrayList[T]) Get(indx int) T {
	var t T

	return t
}

func (al *ArrayList[T]) IndexOf(elem T) int {
	return UndefinedElem
}

func (al *ArrayList[T]) LastIndexOf(elem T) int {
	return UndefinedElem
}

func (al *ArrayList[T]) ListIterator() Iterator[T] {
	return al
}

func (al *ArrayList[T]) ListIteratorIndx(indx int) Iterator[T] {
	return al
}

func (al *ArrayList[T]) RemoveList(indx int) T {
	var t T

	return t
}

func (al *ArrayList[T]) Set(indx int, elem T) T {
	var t T

	return t
}

func (al *ArrayList[T]) Sort(f Comparator[T]) {

}

func (al *ArrayList[T]) SubList(start, end int) []T {
	return nil
}
