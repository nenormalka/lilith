package collections

import "testing"

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
	al := NewArrayList[int]()

	al.Add(1)
	al.Add(2)
	al.Add(3)

	if al.Size() != 3 {
		t.Errorf("Clear() invalid size before clear")
	}

	al.Clear()

	if al.Size() != 0 {
		t.Errorf("Clear() invalid size after clear")
	}
}
