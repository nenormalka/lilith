package structures

import "testing"

func TestBinaryTree(t *testing.T) {
	bt := NewBinaryTree[int]()

	values := []int{8, 3, 10, 1, 6, 14, 4, 7, 13}

	for i := range values {
		bt.Insert(values[i])
	}

	if bt.Search(6) != true {
		t.Errorf("Search(6) should be true")
	}

	node := bt.Get(6)
	if node == nil {
		t.Errorf("Get(6) should not be nil")
	}

	if node.data != 6 {
		t.Errorf("Get(6) should be 6")
	}

	levValues := bt.LevelOrderRootValue()

	for i := range levValues {
		if levValues[i] != values[i] {
			t.Errorf("LevelOrderRootValue() should be %v", values)
		}
	}
}
