package structures

import "testing"

func TestNewTree(t *testing.T) {
	tr, err := NewTree([]*ItemTree[int, int]{
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

	res, err := tr.SearchNodeByValue(600)
	if err != nil {
		t.Error(err)
	}

	if res == nil {
		t.Error("node not found")
	}

	if res.Data.Value != 600 {
		t.Error("wrong node")
	}

	res, err = tr.SearchNodeByValue(111)
	if err != ErrNodeNotFound {
		t.Error(err)
	}

	res, err = tr.SearchNodeByID(6)
	if err != nil {
		t.Error(err)
	}

	if res == nil {
		t.Error("node not found")
	}

	if res.Data.Value != 600 {
		t.Error("wrong node")
	}

	if err = tr.InsertWithParent(&NodeTree[int, int]{
		Data:     &ItemTree[int, int]{Value: 700, ID: 7, ParentID: 6},
		Children: nil,
	}); err != nil {
		t.Error(err)
	}

	if err = tr.InsertWithParent(nil); err != ErrInvalidNode {
		t.Error(err)
	}

	if err = tr.InsertWithParent(&NodeTree[int, int]{
		Data:     &ItemTree[int, int]{Value: 800, ID: 8, ParentID: 666},
		Children: nil,
	}); err != ErrNodeNotFound {
		t.Error(err)
	}

	c, err := tr.GetChildren(6)
	if err != nil {
		t.Error(err)
	}

	if len(c) == 0 {
		t.Error("children not found")
	}

	if c[0].Data.Value != 700 {
		t.Error("wrong children")
	}

	c, err = tr.GetChildren(666)
	if err != ErrNodeNotFound {
		t.Error(err)
	}

	tr, err = NewTree([]*ItemTree[int, int]{})
	if err != ErrEmptyItems {
		t.Errorf("expected error %v, got %v", ErrEmptyItems, err)
	}

	if tr != nil {
		t.Error("expected nil, got not nil")
	}

	tr, err = NewTree([]*ItemTree[int, int]{
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

	tr, err = NewTree([]*ItemTree[int, int]{
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
