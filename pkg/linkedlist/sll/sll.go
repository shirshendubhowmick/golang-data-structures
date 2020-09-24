package sll

import (
	"fmt"
)

// Node : A struct describing the node has payload and a pointer to the next node in the list
type Node struct {
	payload  interface{}
	nextNode *Node
}

// Sll : A singly linked list which hols the elements
type Sll struct {
	head      *Node
	lastNode  *Node
	lastIndex int
}

// Append : Append a node to the end of the list
func (list *Sll) Append(payload interface{}) *Node {
	newNode := new(Node)
	newNode.payload = payload
	if list.lastNode == nil {
		list.head, list.lastNode = newNode, newNode
	} else {
		list.lastNode.nextNode = newNode
		list.lastNode = newNode
	}
	list.lastIndex++
	return newNode
}

// Prepend : Prepend a node to the begining of the list
func (list *Sll) Prepend(payload interface{}) *Node {
	newNode := new(Node)
	newNode.payload = payload
	if list.lastNode == nil {
		list.head, list.lastNode = newNode, newNode
	} else {
		newNode.nextNode = list.head
		list.head = newNode
	}
	list.lastIndex++
	return newNode
}

// RemoveByIndex : Remove node from specified index
func (list *Sll) RemoveByIndex(index int) (interface{}, bool) {
	if index > list.lastIndex || index < 0 || list.head == nil {
		return nil, false
	}

	if index == 0 {
		temp := list.head
		list.lastIndex--
		if list.head == list.lastNode {
			list.head = nil
			list.lastNode = nil
			return temp.payload, true
		}
		list.head = list.head.nextNode
		return temp.payload, true
	}

	previousNode := list.head

	for i := 1; i < index-1; i++ {
		previousNode = previousNode.nextNode
	}

	temp := previousNode.nextNode
	previousNode.nextNode = temp.nextNode
	list.lastIndex--
	return temp.payload, true
}

// ReplaceByIndex : Replace a node with a new node created with a given payload dpending upon the index. Returns the old node payload, newly created node, ok
func (list *Sll) ReplaceByIndex(index int, payload interface{}) (interface{}, *Node, bool) {
	if index > list.lastIndex || index < 0 || list.head == nil {
		return nil, nil, false
	}

	newNode := new(Node)
	newNode.payload = payload

	if index == 0 {
		temp := list.head
		newNode.nextNode = temp.nextNode
		list.head = newNode
		return temp.payload, newNode, true
	}

	previousNode := list.head

	for i := 1; i < index; i++ {
		previousNode = previousNode.nextNode
	}

	temp := previousNode.nextNode
	if previousNode.nextNode == list.lastNode {
		list.lastNode = newNode
	}
	previousNode.nextNode = newNode
	newNode.nextNode = temp.nextNode
	return temp.payload, newNode, true
}

// ReplaceByNode : Replace a node with a new node created with a given payload dpending upon a node reference. Returns the old node payload, newly created node, ok
func (list *Sll) ReplaceByNode(nodeToBeReplaced *Node, payload interface{}) (interface{}, *Node, bool) {
	var currentNode *Node
	var previousNode *Node
	for currentNode = list.head; currentNode != nil; previousNode, currentNode = currentNode, currentNode.nextNode {
		if currentNode == nodeToBeReplaced {
			break
		}
	}

	if previousNode == nil && currentNode == nodeToBeReplaced {
		newNode := new(Node)
		newNode.payload = payload
		newNode.nextNode = currentNode.nextNode
		list.head = newNode
		return currentNode.payload, newNode, true
	}

	if currentNode == nodeToBeReplaced {
		newNode := new(Node)
		newNode.payload = payload
		if currentNode == list.lastNode {
			list.lastNode = newNode
		} else {
			newNode.nextNode = currentNode.nextNode.nextNode
		}
		previousNode.nextNode = newNode

		return currentNode.payload, newNode, true
	}

	return nil, nil, false
}

// Get : Get a node from specified index
func (list *Sll) Get(index int) (*Node, bool) {
	if index > list.lastIndex || index < 0 {
		return nil, false
	}

	if index == 0 {
		return list.head, true
	}

	if index == list.lastIndex {
		return list.lastNode, true
	}

	previousNode := list.head

	for i := 1; i < index; i++ {
		previousNode = previousNode.nextNode
	}

	return previousNode.nextNode, true
}

// IndexOf : Finds first occurance given payload in the list, if found returns its index else -1
func (list *Sll) IndexOf(payload interface{}) (int, bool) {
	currentNode := list.head
	if currentNode == nil {
		fmt.Println("Empty list")
		return -1, false
	}
	i := 0
	for {
		if currentNode.payload == payload {
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
	currentNode := list.head
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
	if list.head == nil {
		return elementsArray
	}

	currentNode := list.head
	for {
		elementsArray = append(elementsArray, currentNode.payload)
		if currentNode.nextNode == nil {
			break
		}
		currentNode = currentNode.nextNode
	}
	return elementsArray
}

// Clear : Clears the list
func (list *Sll) Clear() {
	list.head = nil
	list.lastNode = nil
}

// InitList : Instantiate a new linked list
func InitList(payload interface{}) *Sll {
	list := new(Sll)
	list.head = &Node{
		payload: payload,
	}
	list.lastNode = list.head
	return list
}
