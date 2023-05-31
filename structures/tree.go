package structures

import "errors"

type (
	ItemTree[T comparable] struct {
		Value    T
		ID       T
		ParentID T
	}

	NodeTree[T comparable] struct {
		Data     *ItemTree[T]
		Children []*NodeTree[T]
	}

	Tree[T comparable] struct {
		Root *NodeTree[T]
	}
)

var (
	ErrManyRootNodes  = errors.New("many root nodes")
	ErrEmptyItems     = errors.New("empty items")
	ErrEmptyRootNodes = errors.New("empty root nodes")
)

func NewTree[T comparable](items []*ItemTree[T]) (*Tree[T], error) {
	if len(items) == 0 {
		return nil, ErrEmptyItems
	}

	mapItems := make(map[T]*NodeTree[T], len(items))

	for i := range items {
		mapItems[items[i].ID] = &NodeTree[T]{
			Data: items[i],
		}
	}

	var (
		rootNodes []*NodeTree[T]
		emptyElem T
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

	return &Tree[T]{Root: rootNodes[0]}, nil
}

func (t *Tree[T]) SearchNode(value T) *NodeTree[T] {
	res, found := t.searchNode(t.Root, value)

	if !found {
		return nil
	}

	return res
}

func (t *Tree[T]) searchNode(node *NodeTree[T], value T) (*NodeTree[T], bool) {
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
