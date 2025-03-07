package binarytree

import (
	"errors"
)

var ErrNotEmptyForInsertion = errors.New("Node child poistions not empty for insertion")
var ErrEmptyForDeletion = errors.New("Node child poistions is already empty")

type Payload interface{}

type Node struct {
	Payload    Payload
	leftNode   *Node
	rightNode  *Node
	parentNode *Node
	nodesCount int
}

type BinaryTree struct {
	Root *Node
}

func traverseInOrder(node *Node, result *[]Payload) {
	if node != nil {
		traverseInOrder(node.leftNode, result)
		*result = append(*result, node.Payload)
		traverseInOrder(node.rightNode, result)
	}
}

func traversePreOrder(node *Node, result *[]Payload) {
	if node != nil {
		*result = append(*result, node.Payload)
		traversePreOrder(node.leftNode, result)
		traversePreOrder(node.rightNode, result)
	}
}

func traversePostOrder(node *Node, result *[]Payload) {
	if node != nil {
		traversePostOrder(node.leftNode, result)
		traversePostOrder(node.rightNode, result)
		*result = append(*result, node.Payload)
	}
}

// InsertLeft : Inserts a node to the left of the current node
func (n *Node) InsertLeft(payload Payload) (*Node, error) {
	if n.leftNode != nil {
		return nil, ErrNotEmptyForInsertion
	}

	n.leftNode = &Node{Payload: payload, parentNode: n}
	return n.leftNode, nil
}

// InsertRight : Inserts a node to the right of the current node
func (n *Node) InsertRight(payload Payload) (*Node, error) {
	if n.rightNode != nil {
		return nil, ErrNotEmptyForInsertion
	}

	n.rightNode = &Node{Payload: payload, parentNode: n}
	return n.rightNode, nil
}

// DeleteLeft : Deletes the left child of the node
func (n *Node) DeleteLeft() (Payload, error) {
	if n.leftNode == nil {
		return nil, ErrEmptyForDeletion
	}

	payload := n.leftNode.Payload
	n.leftNode = nil
	return payload, nil
}

// DeleteRight : Deletes the right child of the node
func (n *Node) DeleteRight() (Payload, error) {
	if n.rightNode == nil {
		return nil, ErrEmptyForDeletion
	}

	payload := n.rightNode.Payload
	n.rightNode = nil
	return payload, nil
}

func (n *Node) Height() int {
	if n == nil {
		return 0
	}

	leftDepth := n.leftNode.Height()
	rightDepth := n.rightNode.Height()

	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

// ToSliceInOrder : Returns a slice of the tree in in-order traversal
func (bt *BinaryTree) ToSliceInOrder() *[]Payload {
	if bt.Root == nil {
		return nil
	}

	s := make([]Payload, bt.Root.nodesCount)
	traverseInOrder(bt.Root, &s)
	return &s
}

// ToSlicePreOrder : Returns a slice of the tree in pre-order
func (bt *BinaryTree) ToSlicePreOrder() *[]Payload {
	if bt.Root == nil {
		return nil
	}

	s := make([]Payload, bt.Root.nodesCount)
	traversePreOrder(bt.Root, &s)
	return &s
}

// ToSlicePostOrder : Returns a slice of the tree in post order
func (bt *BinaryTree) ToSlicePostOrder() *[]Payload {
	if bt.Root == nil {
		return nil
	}

	s := make([]Payload, bt.Root.nodesCount)
	traversePostOrder(bt.Root, &s)
	return &s
}

// Depth : Returns the depth of the tree
func (tr *BinaryTree) Depth() int {
	return tr.Root.Height()
}

// Init : Initializes a new binary tree with the given payload
func Init(payload interface{}) *BinaryTree {
	tree := new(BinaryTree)
	tree.Root = new(Node)
	tree.Root.Payload = payload

	return tree
}
