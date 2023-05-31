package structures

import (
	"golang.org/x/exp/constraints"
)

type (
	BinaryTreeNode[T constraints.Ordered] struct {
		data  T
		left  *BinaryTreeNode[T]
		right *BinaryTreeNode[T]
	}

	BinaryTree[T constraints.Ordered] struct {
		root *BinaryTreeNode[T]
	}
)

func NewBinaryTree[T constraints.Ordered]() *BinaryTree[T] {
	return &BinaryTree[T]{}
}

func (bt *BinaryTree[T]) Insert(val T) {
	bt.insertRec(bt.root, val)
}

func (bt *BinaryTree[T]) Search(val T) bool {
	return bt.searchRec(bt.root, val) != nil
}

func (bt *BinaryTree[T]) Get(val T) *BinaryTreeNode[T] {
	return bt.searchRec(bt.root, val)
}

func (bt *BinaryTree[T]) InOrderRootValue() []T {
	return bt.InOrderValue(bt.root)
}

func (bt *BinaryTree[T]) InOrderValue(node *BinaryTreeNode[T]) []T {
	if node == nil {
		return nil
	} else {
		var (
			nodes  []*BinaryTreeNode[T]
			values []T
		)

		bt.inOrderRec(node, nodes)

		for i := range nodes {
			values = append(values, nodes[i].data)
		}

		return values
	}
}

func (bt *BinaryTree[T]) InOrderRootNodes() []*BinaryTreeNode[T] {
	return bt.InOrderNodes(bt.root)
}

func (bt *BinaryTree[T]) InOrderNodes(node *BinaryTreeNode[T]) []*BinaryTreeNode[T] {
	if node == nil {
		return nil
	} else {
		var nodes []*BinaryTreeNode[T]

		bt.inOrderRec(node, nodes)

		return nodes
	}
}

func (bt *BinaryTree[T]) LevelOrderRootValue() []T {
	return bt.LevelOrderValue(bt.root)
}

func (bt *BinaryTree[T]) LevelOrderValue(node *BinaryTreeNode[T]) []T {
	if node == nil {
		return nil
	} else {
		var values []T

		nodes := bt.levelOrderCommon(node)

		for i := range nodes {
			values = append(values, nodes[i].data)
		}

		return values
	}
}

func (bt *BinaryTree[T]) LevelOrderRootNodes() []*BinaryTreeNode[T] {
	return bt.LevelOrderNodes(bt.root)
}

func (bt *BinaryTree[T]) LevelOrderNodes(node *BinaryTreeNode[T]) []*BinaryTreeNode[T] {
	if node == nil {
		return nil
	} else {
		return bt.levelOrderCommon(node)
	}
}

func (bt *BinaryTree[T]) insertRec(node *BinaryTreeNode[T], val T) *BinaryTreeNode[T] {
	if bt.root == nil {
		bt.root = &BinaryTreeNode[T]{val, nil, nil}

		return bt.root
	}

	if node == nil {
		return &BinaryTreeNode[T]{val, nil, nil}
	}

	if val <= node.data {
		node.left = bt.insertRec(node.left, val)
	}

	if val > node.data {
		node.right = bt.insertRec(node.right, val)
	}

	return node
}

func (bt *BinaryTree[T]) searchRec(node *BinaryTreeNode[T], val T) *BinaryTreeNode[T] {
	if node == nil {
		return nil
	}

	if node.data == val {
		return node
	}

	if val < node.data {
		return bt.searchRec(node.left, val)
	}

	if val > node.data {
		return bt.searchRec(node.right, val)
	}

	return nil
}

func (bt *BinaryTree[T]) inOrderRec(node *BinaryTreeNode[T], nodes []*BinaryTreeNode[T]) {
	if node == nil {
		return
	}

	bt.inOrderRec(node.left, nodes)

	nodes = append(nodes, node)

	bt.inOrderRec(node.right, nodes)
}

func (bt *BinaryTree[T]) levelOrderCommon(node *BinaryTreeNode[T]) []*BinaryTreeNode[T] {
	if node == nil {
		return nil
	}

	var nodes []*BinaryTreeNode[T]

	nodeList := append([]*BinaryTreeNode[T]{}, node)

	for !(len(nodeList) == 0) {
		current := nodeList[0]
		nodes = append(nodes, current)

		if current.left != nil {
			nodeList = append(nodeList, current.left)
		}

		if current.right != nil {
			nodeList = append(nodeList, current.right)
		}

		nodeList = nodeList[1:]
	}

	return nodes
}
