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

func (node *Node) InsertLeft(payload interface{}) (*Node, error) {
	if node.leftNode != nil {
		return nil, ErrNotEmptyForInsertion
	}

	newNode := new(Node)
	newNode.Payload = payload
	newNode.parentNode = node
	node.leftNode = newNode
	return newNode, nil
}

func (node *Node) InsertRight(payload interface{}) (*Node, error) {
	if node.rightNode != nil {
		return nil, ErrNotEmptyForInsertion
	}

	newNode := new(Node)
	newNode.Payload = payload
	newNode.parentNode = node
	node.rightNode = newNode
	return newNode, nil
}

func (node *Node) DeleteLeft() (*Node, error) {
	if node.leftNode == nil {
		return nil, ErrEmptyForDeletion
	}

	deletedNode := node.leftNode
	deletedNode.parentNode = nil
	node.leftNode = nil
	return deletedNode, nil
}

func (node *Node) DeleteRight() (*Node, error) {
	if node.rightNode == nil {
		return nil, ErrEmptyForDeletion
	}

	deletedNode := node.rightNode
	deletedNode.parentNode = nil
	node.rightNode = nil
	return deletedNode, nil
}

type BinaryTree struct {
	Root *Node
}

func (tree *BinaryTree) ToSliceInOrder() []interface{} {
	payloadSlice := make([]interface{}, 0)
	inOrderTraversal(tree.Root, &payloadSlice)
	return payloadSlice
}

func (tree *BinaryTree) ToSlicePreOrder() []interface{} {
	payloadSlice := make([]interface{}, 0)
	preOrderTraversal(tree.Root, &payloadSlice)
	return payloadSlice
}

func (tree *BinaryTree) ToSlicePostOrder() []interface{} {
	payloadSlice := make([]interface{}, 0)
	postOrderTraversal(tree.Root, &payloadSlice)
	return payloadSlice
}

func Init(payload interface{}) *BinaryTree {
	tree := new(BinaryTree)
	tree.Root = new(Node)

	tree.Root.Payload = payload

	return tree
}
