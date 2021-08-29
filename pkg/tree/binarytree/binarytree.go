package binarytree

import "errors"

var ErrNotEmptyForInsertion = errors.New("Node child poistions not empty for insertion")
var ErrEmptyForDeletion = errors.New("Node child poistions is already empty")

type Node struct {
	Payload    interface{}
	leftNode   *Node
	rightNode  *Node
	parentNode *Node
}

type BinaryTree struct {
	Root *Node
}

func (tree *BinaryTree) InsertLeft(node *Node, payload interface{}) (*Node, error) {
	if node.leftNode != nil {
		return nil, ErrNotEmptyForInsertion
	}

	newNode := new(Node)
	newNode.Payload = payload
	newNode.parentNode = node
	node.leftNode = newNode
	return newNode, nil
}

func (tree *BinaryTree) InserRight(node *Node, payload interface{}) (*Node, error) {
	if node.rightNode != nil {
		return nil, ErrNotEmptyForInsertion
	}

	newNode := new(Node)
	newNode.Payload = payload
	newNode.parentNode = node
	node.rightNode = newNode
	return newNode, nil
}

func (tree *BinaryTree) DeleteLeft(node *Node) (*Node, error) {
	if node.leftNode == nil {
		return nil, ErrEmptyForDeletion
	}

	deletedNode := node.leftNode
	deletedNode.parentNode = nil
	node.leftNode = nil
	return deletedNode, nil
}

func (tree *BinaryTree) DeleteRight(node *Node) (*Node, error) {
	if node.rightNode == nil {
		return nil, ErrEmptyForDeletion
	}

	deletedNode := node.rightNode
	deletedNode.parentNode = nil
	node.rightNode = nil
	return deletedNode, nil
}

func Init(payload interface{}) *BinaryTree {
	tree := new(BinaryTree)
	tree.Root.Payload = payload

	return tree
}
