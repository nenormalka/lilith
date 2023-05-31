package structures

import "errors"

type (
	ItemTree[T, M comparable] struct {
		Value    T
		ID       M
		ParentID M
	}

	NodeTree[T, M comparable] struct {
		Data     *ItemTree[T, M]
		Children []*NodeTree[T, M]
	}

	Tree[T, M comparable] struct {
		Root *NodeTree[T, M]
		m    map[M]*NodeTree[T, M]
	}
)

var (
	ErrManyRootNodes  = errors.New("many root nodes")
	ErrEmptyItems     = errors.New("empty items")
	ErrEmptyRootNodes = errors.New("empty root nodes")
	ErrNodeNotFound   = errors.New("node not found")
	ErrInvalidNode    = errors.New("invalid node")
)

func NewTree[T, M comparable](items []*ItemTree[T, M]) (*Tree[T, M], error) {
	if len(items) == 0 {
		return nil, ErrEmptyItems
	}

	mapItems := make(map[M]*NodeTree[T, M], len(items))

	for i := range items {
		mapItems[items[i].ID] = &NodeTree[T, M]{
			Data: items[i],
		}
	}

	var (
		rootNodes []*NodeTree[T, M]
		emptyElem M
	)

	for _, node := range mapItems {
		if node.Data.ParentID == emptyElem {
			rootNodes = append(rootNodes, node)

			continue
		}

		parentItem, ok := mapItems[node.Data.ParentID]
		if !ok {
			continue
		}

		parentItem.Children = append(parentItem.Children, node)
	}

	if len(rootNodes) == 0 {
		return nil, ErrEmptyRootNodes
	}

	if len(rootNodes) > 1 {
		return nil, ErrManyRootNodes
	}

	return &Tree[T, M]{
		Root: rootNodes[0],
		m:    mapItems,
	}, nil
}

func (t *Tree[T, M]) GetChildren(parentID M) ([]*NodeTree[T, M], error) {
	parent, ok := t.m[parentID]
	if !ok {
		return nil, ErrNodeNotFound
	}

	return parent.Children, nil
}

func (t *Tree[T, M]) InsertWithParent(node *NodeTree[T, M]) error {
	if node == nil {
		return ErrInvalidNode
	}

	parent, ok := t.m[node.Data.ParentID]
	if !ok {
		return ErrNodeNotFound
	}

	parent.Children = append(parent.Children, node)
	t.m[node.Data.ID] = node

	return nil
}

func (t *Tree[T, M]) SearchNodeByID(id M) (*NodeTree[T, M], error) {
	node, ok := t.m[id]
	if !ok {
		return nil, ErrNodeNotFound
	}

	return node, nil
}

func (t *Tree[T, M]) SearchNodeByValue(value T) (*NodeTree[T, M], error) {
	res, found := t.searchNode(t.Root, value)

	if !found {
		return nil, ErrNodeNotFound
	}

	return res, nil
}

func (t *Tree[T, M]) searchNode(node *NodeTree[T, M], value T) (*NodeTree[T, M], bool) {
	if node.Data.Value == value {
		return node, true
	}

	for i := range node.Children {
		if res, found := t.searchNode(node.Children[i], value); found {
			return res, found
		}
	}

	return nil, false
}
