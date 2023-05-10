package collections

import (
	"reflect"
	"testing"
)

func TestArrayList_Add(t *testing.T) {
	al := NewArrayList[int]()

	if !al.Add(1) {
		t.Errorf("add() invalid")
	}

	if al.Size() != 1 {
		t.Errorf("add() invalid size")
	}

	elem, err := al.Get(0)

	if err != nil {
		t.Errorf("add() err to get elem")
	}

	if elem != 1 {
		t.Errorf("add() invalid elem")
	}
}

func TestArrayList_AddAll(t *testing.T) {
	m := map[int]struct{}{
		1: {},
		3: {},
		4: {},
	}

	al1 := NewArrayList[int]()

	for key := range m {
		al1.Add(key)
	}

	al2 := NewArrayList[int]()
	al2.AddAll(al1)

	if al2.Size() != 3 {
		t.Errorf("AddAll() invalid size")
	}

	for i := range al2.ToArray() {
		if _, ok := m[al2.ToArray()[i]]; ok {
			delete(m, al2.ToArray()[i])
		} else {
			t.Errorf("AddAll() invalid value")
		}
	}

	if len(m) != 0 {
		t.Errorf("AddAll() invalid values in map")
	}
}

func TestArrayList_Clear(t *testing.T) {
	al := createAL[int](1, 2, 3)

	if al.Size() != 3 {
		t.Errorf("Clear() invalid size before clear")
	}

	al.Clear()

	if al.Size() != 0 {
		t.Errorf("Clear() invalid size after clear")
	}
}

func TestArrayList_Contains(t *testing.T) {
	al := createAL[int](1, 2, 3)

	if al.Contains(0) {
		t.Errorf("Contains() invalid value 0")
	}

	if !al.Contains(2) {
		t.Errorf("Contains() invalid value 2")
	}
}

func TestArrayList_ContainsAll(t *testing.T) {
	al1 := createAL[int](1, 2, 3)
	al2 := createAL[int](1, 2, 3)
	al3 := createAL[int](1, 2)
	al4 := createAL[int](1, 2, 3, 4)
	al5 := createAL[int](1, 2, 3, 3)

	if !al1.ContainsAll(al2) {
		t.Errorf("ContainsAll() invalid, al1 %v, al2 %v", al1.ToArray(), al2.ToArray())
	}

	if !al1.ContainsAll(al3) {
		t.Errorf("ContainsAll() invalid, al1 %v, al3 %v", al1.ToArray(), al3.ToArray())
	}

	if al1.ContainsAll(al4) {
		t.Errorf("ContainsAll() invalid, al1 %v, al4 %v", al1.ToArray(), al4.ToArray())
	}

	if !al1.ContainsAll(al5) {
		t.Errorf("ContainsAll() invalid, al1 %v, al5 %v", al1.ToArray(), al5.ToArray())
	}
}

func TestArrayList_Equals(t *testing.T) {
	al1 := createAL[int](1, 2, 3)
	al2 := createAL[int](1, 2)
	al3 := createAL[int](4, 5, 6)
	al4 := createAL[int](1, 2, 3)

	if al1.Equals(1) {
		t.Errorf("Equals() invalid al1 and 1")
	}

	if al1.Equals(al2) {
		t.Errorf("Equals() invalid, al1 %v, al2 %v", al1.ToArray(), al2.ToArray())
	}

	if al1.Equals(al3) {
		t.Errorf("Equals() invalid, al1 %v, al3 %v", al1.ToArray(), al3.ToArray())
	}

	if !al1.Equals(al4) {
		t.Errorf("Equals() invalid, al1 %v, al4 %v", al1.ToArray(), al4.ToArray())
	}
}

func TestArrayList_IsEmpty(t *testing.T) {
	al := createAL[int]()

	if !al.IsEmpty() {
		t.Errorf("IsEmpty() invalid empty before add")
	}

	al.Add(1)

	if al.IsEmpty() {
		t.Errorf("IsEmpty() invalid empty after add")
	}
}

func TestArrayList_Remove(t *testing.T) {
	al := createAL[int](1, 2, 3)

	al.Remove(2)

	if !reflect.DeepEqual(al.ToArray(), []int{1, 3}) {
		t.Errorf("Remove() invalid, got %v, want %v", al.ToArray(), []int{1, 3})
	}
}

func TestArrayList_RemoveAll(t *testing.T) {
	al := createAL[int](1, 2, 3)

	al.RemoveAll(createAL[int](1, 3))

	if !reflect.DeepEqual(al.ToArray(), []int{2}) {
		t.Errorf("RemoveAll() invalid, got %v, want %v", al.ToArray(), []int{2})
	}
}

func TestArrayList_RetainAll(t *testing.T) {
	al := createAL[int](1, 2, 3, 4)

	al.RetainAll(createAL[int](1, 3))

	if !reflect.DeepEqual(al.ToArray(), []int{1, 3}) {
		t.Errorf("RetainAll() invalid, got %v, want %v", al.ToArray(), []int{1, 3})
	}
}

func TestArrayList_Size(t *testing.T) {
	al := createAL[int](1, 2, 3, 4)

	if al.Size() != 4 {
		t.Errorf("Size() invalid")
	}
}

func TestArrayList_ToArray(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	al := createAL[int](arr...)

	if !reflect.DeepEqual(al.ToArray(), arr) {
		t.Errorf("ToArray() invalid, got %v, want %v", al.ToArray(), []int{1, 3})
	}
}

func TestArrayList_RemoveIf(t *testing.T) {
	al := createAL[int](1, 2, 3, 4)

	al.RemoveIf(func(elem int) bool {
		return elem%2 == 0
	})

	if !reflect.DeepEqual(al.ToArray(), []int{1, 3}) {
		t.Errorf("RemoveIf() invalid, got %v, want %v", al.ToArray(), []int{1, 3})
	}
}

func TestArrayList_AddList(t *testing.T) {
	al := createAL[int](1, 2, 3, 4)

	if !al.AddList(2, 5) {
		t.Errorf("AddList() invalid")
	}

	if !reflect.DeepEqual(al.ToArray(), []int{1, 2, 5, 3, 4}) {
		t.Errorf("AddList() invalid, got %v, want %v", al.ToArray(), []int{1, 2, 5, 3, 4})
	}
}

func TestArrayList_AddAllList(t *testing.T) {
	al := createAL[int](1, 2, 3, 4)

	if !al.AddAllList(2, createAL[int](1, 2, 3)) {
		t.Errorf("AddAllList() invalid")
	}

	if !reflect.DeepEqual(al.ToArray(), []int{1, 2, 1, 2, 3, 3, 4}) {
		t.Errorf("AddAllList() invalid, got %v, want %v", al.ToArray(), []int{1, 2, 1, 2, 3, 3, 4})
	}
}

func TestArrayList_Get(t *testing.T) {
	al := createAL[int](1, 2, 3, 4)

	got, err := al.Get(1)
	if err != nil {
		t.Errorf("Get() invalid err exists index")
	}

	if got != 2 {
		t.Errorf("Get() invalid value exists index")
	}

	got, err = al.Get(15)
	if err == nil {
		t.Errorf("Get() invalid err not exists index")
	}

	if err != ErrInvalidIndex {
		t.Errorf("Get() undefined err not exists index")
	}

	if got != 0 {
		t.Errorf("Get() invalid value not exists index")
	}
}

func TestArrayList_IndexOf(t *testing.T) {
	al := createAL[int](1, 2, 3, 4, 2)

	if al.IndexOf(2) != 1 {
		t.Errorf("IndexOf() invalid index exists elem")
	}

	if al.IndexOf(-100) != UndefinedElem {
		t.Errorf("IndexOf() invalid index not exists elem")
	}
}

func TestArrayList_LastIndexOf(t *testing.T) {
	al := createAL[int](1, 2, 3, 4, 2)

	if al.LastIndexOf(2) != 4 {
		t.Errorf("LastIndexOf() invalid index exists elem")
	}

	if al.LastIndexOf(-100) != UndefinedElem {
		t.Errorf("LastIndexOf() invalid index not exists elem")
	}
}

func TestArrayList_ListIterator(t *testing.T) {
	m := map[int]struct{}{
		1: {},
		3: {},
		4: {},
	}

	al := NewArrayList[int]()

	for key := range m {
		al.Add(key)
	}

	iter := al.ListIterator()

	for iter.HasNext() {
		delete(m, iter.GetNext())
	}

	if len(m) != 0 {
		t.Errorf("ListIterator() invalid values in map")
	}
}

func TestArrayList_ListIteratorIndx(t *testing.T) {
	m := map[int]struct{}{
		3: {},
		4: {},
	}

	al := createAL[int](1, 2, 3, 4)

	iter := al.ListIteratorIndx(2)

	for iter.HasNext() {
		delete(m, iter.GetNext())
	}

	if len(m) != 0 {
		t.Errorf("ListIteratorIndx() invalid values in map")
	}
}

func TestArrayList_RemoveList(t *testing.T) {
	al := createAL[int](1, 2, 3, 4)

	got, err := al.RemoveList(2)

	if err != nil {
		t.Errorf("RemoveList() invalid err exists index")
	}

	if got != 3 {
		t.Errorf("RemoveList() invalid value exists index")
	}

	got, err = al.Get(15)
	if err == nil {
		t.Errorf("RemoveList() invalid err not exists index")
	}

	if err != ErrInvalidIndex {
		t.Errorf("RemoveList() undefined err not exists index")
	}

	if got != 0 {
		t.Errorf("RemoveList() invalid value not exists index")
	}
}

func TestArrayList_Set(t *testing.T) {
	al := createAL[int](1, 2, 3, 4)

	got, err := al.Set(2, 8)

	if err != nil {
		t.Errorf("Set() invalid err exists index")
	}

	if got != 3 {
		t.Errorf("Set() invalid value exists index")
	}

	got, err = al.Get(2)

	if err != nil {
		t.Errorf("Set() get invalid err exists index")
	}

	if got != 8 {
		t.Errorf("Set() get invalid value exists index")
	}

	got, err = al.Get(15)
	if err == nil {
		t.Errorf("Set() invalid err not exists index")
	}

	if err != ErrInvalidIndex {
		t.Errorf("Set() undefined err not exists index")
	}

	if got != 0 {
		t.Errorf("Set() invalid value not exists index")
	}
}

func TestArrayList_Sort(t *testing.T) {
	al := createAL[int](1, 2, 3, 4)

	al.Sort(func(first, second int) bool {
		return first > second
	})

	if !reflect.DeepEqual(al.ToArray(), []int{4, 3, 2, 1}) {
		t.Errorf("Sort() invalid, got %v, want %v", al.ToArray(), []int{4, 3, 2, 1})
	}
}

func TestArrayList_SubList(t *testing.T) {
	al := createAL[int](1, 2, 3, 4)

	arr, err := al.SubList(2, 4)

	if err != nil {
		t.Errorf("SubList() invalid err exists index")
	}

	if !reflect.DeepEqual(arr, []int{3, 4}) {
		t.Errorf("SubList() invalid value exists index")
	}

	arr, err = al.SubList(15, 200)
	if err == nil {
		t.Errorf("SubList() invalid err not exists index")
	}

	if err != ErrInvalidIndex {
		t.Errorf("SubList() undefined err not exists index")
	}

	arr, err = al.SubList(15, 1)
	if err == nil {
		t.Errorf("SubList() invalid err not exists index")
	}

	if err != ErrInvalidIndex {
		t.Errorf("SubList() undefined err not exists index")
	}
}

func createAL[T comparable](args ...T) *ArrayList[T] {
	al := NewArrayList[T]()

	for i := range args {
		al.Add(args[i])
	}

	return al
}
