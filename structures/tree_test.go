package structures

import "testing"

func TestNewTree(t *testing.T) {
	tr, err := NewTree([]*ItemTree[int]{
		{
			Value:    100,
			ID:       1,
			ParentID: 0,
		},
		{
			Value:    200,
			ID:       2,
			ParentID: 1,
		},
		{
			Value:    300,
			ID:       3,
			ParentID: 1,
		},
		{
			Value:    400,
			ID:       4,
			ParentID: 2,
		},
		{
			Value:    500,
			ID:       5,
			ParentID: 2,
		},
		{
			Value:    600,
			ID:       6,
			ParentID: 4,
		},
	})

	if err != nil {
		t.Error(err)
	}

	res := tr.SearchNode(600)
	if res == nil {
		t.Error("node not found")
	}

	if res.Data.Value != 600 {
		t.Error("wrong node")
	}

	tr, err = NewTree([]*ItemTree[int]{})
	if err != ErrEmptyItems {
		t.Errorf("expected error %v, got %v", ErrEmptyItems, err)
	}

	if tr != nil {
		t.Error("expected nil, got not nil")
	}

	tr, err = NewTree([]*ItemTree[int]{
		{
			Value:    1,
			ID:       1,
			ParentID: 0,
		},
		{
			Value:    2,
			ID:       2,
			ParentID: 0,
		},
	})

	if err != ErrManyRootNodes {
		t.Errorf("expected error %v, got %v", ErrEmptyItems, err)
	}

	if tr != nil {
		t.Error("expected nil, got not nil")
	}

	tr, err = NewTree([]*ItemTree[int]{
		{
			Value:    1,
			ID:       1,
			ParentID: 3,
		},
		{
			Value:    2,
			ID:       2,
			ParentID: 3,
		},
	})

	if err != ErrEmptyRootNodes {
		t.Errorf("expected error %v, got %v", ErrEmptyItems, err)
	}

	if tr != nil {
		t.Error("expected nil, got not nil")
	}
}
