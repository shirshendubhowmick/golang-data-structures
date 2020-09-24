package sll

import (
	"fmt"
)

// Node : A struct describing the node has payload and a pointer to the next node in the list
type Node struct {
	Payload  interface{}
	nextNode *Node
}

// Sll : A singly linked list which hols the elements
type Sll struct {
	Head      *Node
	LastNode  *Node
	lastIndex int
}

// Append : Append a node to the end of the list
func (list *Sll) Append(payload interface{}) *Node {
	newNode := new(Node)
	newNode.Payload = payload
	if list.LastNode == nil {
		list.Head, list.LastNode = newNode, newNode
	} else {
		list.LastNode.nextNode = newNode
		list.LastNode = newNode
	}
	list.lastIndex++
	return newNode
}

// Prepend : Prepend a node to the begining of the list
func (list *Sll) Prepend(payload interface{}) *Node {
	newNode := new(Node)
	newNode.Payload = payload
	if list.LastNode == nil {
		list.Head, list.LastNode = newNode, newNode
	} else {
		newNode.nextNode = list.Head
		list.Head = newNode
	}
	list.lastIndex++
	return newNode
}

// RemoveByIndex : Remove node from specified index
func (list *Sll) RemoveByIndex(index int) (interface{}, bool) {
	if index > list.lastIndex || index < 0 || list.Head == nil {
		return nil, false
	}

	if index == 0 {
		temp := list.Head
		list.lastIndex--
		if list.Head == list.LastNode {
			list.Head = nil
			list.LastNode = nil
			return temp.Payload, true
		}
		list.Head = list.Head.nextNode
		return temp.Payload, true
	}

	previousNode := list.Head

	for i := 1; i < index-1; i++ {
		previousNode = previousNode.nextNode
	}

	temp := previousNode.nextNode
	previousNode.nextNode = temp.nextNode
	list.lastIndex--
	return temp.Payload, true
}

// ReplaceByIndex : Replace a node with a new node created with a given payload dpending upon the index. Returns the old node payload, newly created node, ok
func (list *Sll) ReplaceByIndex(index int, payload interface{}) (interface{}, *Node, bool) {
	if index > list.lastIndex || index < 0 || list.Head == nil {
		return nil, nil, false
	}

	newNode := new(Node)
	newNode.Payload = payload

	if index == 0 {
		temp := list.Head
		newNode.nextNode = temp.nextNode
		list.Head = newNode
		return temp.Payload, newNode, true
	}

	previousNode := list.Head

	for i := 1; i < index; i++ {
		previousNode = previousNode.nextNode
	}

	temp := previousNode.nextNode
	if previousNode.nextNode == list.LastNode {
		list.LastNode = newNode
	}
	previousNode.nextNode = newNode
	newNode.nextNode = temp.nextNode
	return temp.Payload, newNode, true
}

// ReplaceByNode : Replace a node with a new node created with a given payload dpending upon a node reference. Returns the old node payload, newly created node, ok
func (list *Sll) ReplaceByNode(nodeToBeReplaced *Node, payload interface{}) (interface{}, *Node, bool) {
	var currentNode *Node
	var previousNode *Node
	for currentNode = list.Head; currentNode != nil; previousNode, currentNode = currentNode, currentNode.nextNode {
		if currentNode == nodeToBeReplaced {
			break
		}
	}

	if previousNode == nil && currentNode == nodeToBeReplaced {
		newNode := new(Node)
		newNode.Payload = payload
		newNode.nextNode = currentNode.nextNode
		list.Head = newNode
		return currentNode.Payload, newNode, true
	}

	if currentNode == nodeToBeReplaced {
		newNode := new(Node)
		newNode.Payload = payload
		if currentNode == list.LastNode {
			list.LastNode = newNode
		} else {
			newNode.nextNode = currentNode.nextNode.nextNode
		}
		previousNode.nextNode = newNode

		return currentNode.Payload, newNode, true
	}

	return nil, nil, false
}

// Get : Get a node from specified index
func (list *Sll) Get(index int) (*Node, bool) {
	if index > list.lastIndex || index < 0 {
		return nil, false
	}

	if index == 0 {
		return list.Head, true
	}

	if index == list.lastIndex {
		return list.LastNode, true
	}

	previousNode := list.Head

	for i := 1; i < index; i++ {
		previousNode = previousNode.nextNode
	}

	return previousNode.nextNode, true
}

// IndexOf : Finds first occurance given payload in the list, if found returns its index else -1
func (list *Sll) IndexOf(payload interface{}) (int, bool) {
	currentNode := list.Head
	if currentNode == nil {
		fmt.Println("Empty list")
		return -1, false
	}
	i := 0
	for {
		if currentNode.Payload == payload {
			return i, true
		}
		if currentNode.nextNode == nil {
			break
		}
		currentNode = currentNode.nextNode
		i++
	}
	return -1, false
}

//Print : Print the list in using fmt.Println canbe used for debugging
func (list *Sll) Print() {
	currentNode := list.Head
	if currentNode == nil {
		fmt.Println("Empty list")
		return
	}
	for {
		fmt.Println(currentNode)
		if currentNode.nextNode == nil {
			break
		}
		currentNode = currentNode.nextNode
	}
}

// Length : Get the current length of the list
func (list *Sll) Length() int {
	return list.lastIndex + 1
}

// ToSlice : Returns the elements present in the list as an array, orders are same as of the list
func (list *Sll) ToSlice() []interface{} {
	elementsArray := []interface{}{}
	if list.Head == nil {
		return elementsArray
	}

	currentNode := list.Head
	for {
		elementsArray = append(elementsArray, currentNode.Payload)
		if currentNode.nextNode == nil {
			break
		}
		currentNode = currentNode.nextNode
	}
	return elementsArray
}

// Clear : Clears the list
func (list *Sll) Clear() {
	list.Head = nil
	list.LastNode = nil
}

// InitList : Instantiate a new linked list
func InitList(payload interface{}) *Sll {
	list := new(Sll)
	list.Head = &Node{
		Payload: payload,
	}
	list.LastNode = list.Head
	return list
}
