package collections

import (
	"errors"
	"sort"

	"github.com/nenormalka/lilith/methods"
	"github.com/nenormalka/lilith/patterns"
)

type (
	ArrayList[T comparable] struct {
		data []T
	}
)

const (
	UndefinedElem = -1
)

var (
	_ List[int] = &ArrayList[int]{}

	ErrInvalidIndex = errors.New("invalid index")
)

func NewArrayList[T comparable]() *ArrayList[T] {
	return &ArrayList[T]{data: make([]T, 0)}
}

func (al *ArrayList[T]) Add(elem T) bool {
	al.data = append(al.data, elem)

	return true
}

func (al *ArrayList[T]) AddAll(elems Collection[T]) bool {
	al.data = append(al.data, elems.ToArray()...)

	return true
}

func (al *ArrayList[T]) Clear() {
	al.data = make([]T, 0)
}

func (al *ArrayList[T]) Contains(elem T) bool {
	return methods.InArray[[]T](al.data, elem)
}

func (al *ArrayList[T]) ContainsAll(elems Collection[T]) bool {
	for j := range elems.ToArray() {
		if !methods.InArray[[]T](al.data, elems.ToArray()[j]) {
			return false
		}
	}

	return true
}

func (al *ArrayList[T]) Equals(elem Collection[T]) bool {
	return Collection[T](al) == elem
}

func (al *ArrayList[T]) IsEmpty() bool {
	return len(al.data) == 0
}

func (al *ArrayList[T]) Remove(elem T) bool {
	for i := range al.data {
		if al.data[i] == elem {
			methods.ArrayRemove[[]T](&al.data, i)

			return true
		}
	}

	return false
}

func (al *ArrayList[T]) RemoveAll(elems Collection[T]) bool {
	l := len(al.data)
	if l == 0 {
		return false
	}

	m := methods.ArrayToMapValues[[]T](elems.ToArray())
	if len(m) == 0 {
		return false
	}

	isRemoved := false

	for i := 0; i < l; i++ {
		if _, ok := m[al.data[i]]; ok {
			methods.ArrayRemove[[]T](&al.data, i)
			i++

			if !isRemoved {
				isRemoved = true
			}
		}
	}

	return isRemoved
}

func (al *ArrayList[T]) RetainAll(elems Collection[T]) bool {
	l := len(al.data)
	if l == 0 {
		return false
	}

	m := methods.ArrayToMapValues[[]T](elems.ToArray())
	if len(m) == 0 {
		return false
	}

	isUpdated := false
	arr := make([]T, 0)

	for i := 0; i < l; i++ {
		if _, ok := m[al.data[i]]; ok {
			arr = append(arr, al.data[i])

			if !isUpdated {
				isUpdated = true
			}
		}
	}

	return isUpdated
}

func (al *ArrayList[T]) Size() int {
	return len(al.data)
}

func (al *ArrayList[T]) ToArray() []T {
	return al.data
}

func (al *ArrayList[T]) RemoveIf(f Predicate[T]) {
	l := len(al.data)

	for i := 0; i < l; i++ {
		if f(al.data[i]) {
			methods.ArrayRemove[[]T](&al.data, i)
			i++
		}
	}
}

func (al *ArrayList[T]) AddList(indx int, elem T) bool {
	arr := al.data[:indx]
	arr = append(arr, elem)
	al.data = append(arr, al.data[indx:]...)

	return true
}

func (al *ArrayList[T]) AddAllList(indx int, elems []T) bool {
	arr := al.data[:indx]
	arr = append(arr, elems...)
	al.data = append(arr, al.data[indx:]...)

	return false
}

func (al *ArrayList[T]) Get(indx int) (T, error) {
	if len(al.data) < indx {
		var t T

		return t, ErrInvalidIndex
	}

	return al.data[indx], nil
}

func (al *ArrayList[T]) IndexOf(elem T) int {
	for i := range al.data {
		if al.data[i] == elem {
			return i
		}
	}

	return UndefinedElem
}

func (al *ArrayList[T]) LastIndexOf(elem T) int {
	for i := len(al.data) - 1; i >= 0; i-- {
		if al.data[i] == elem {
			return i
		}
	}

	return UndefinedElem
}

func (al *ArrayList[T]) ListIterator() Iterator[T] {
	return patterns.NewIterator[T](al.data)
}

func (al *ArrayList[T]) ListIteratorIndx(indx int) Iterator[T] {
	if len(al.data) < indx {
		return patterns.NewIterator[T](nil)
	}

	return patterns.NewIterator[T](al.data[indx:])
}

func (al *ArrayList[T]) RemoveList(indx int) (T, error) {
	var t T

	if len(al.data) < indx {
		return t, ErrInvalidIndex
	}

	t = al.data[indx]

	methods.ArrayRemove[[]T](&al.data, indx)

	return t, nil
}

func (al *ArrayList[T]) Set(indx int, elem T) (T, error) {
	var t T

	if len(al.data) < indx {
		return t, ErrInvalidIndex
	}

	t = al.data[indx]
	al.data[indx] = elem

	return t, nil
}

func (al *ArrayList[T]) Sort(f Comparator[T]) {
	sort.Slice(al.data, func(i, j int) bool {
		return f(al.data[i], al.data[j])
	})
}

func (al *ArrayList[T]) SubList(start, end int) ([]T, error) {
	if len(al.data) < start || end < start {
		return nil, ErrInvalidIndex
	}

	return methods.ArraySlice[[]T](al.data, start, end-start)
}
