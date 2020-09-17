package sll

import (
	"fmt"
)

type node struct {
	payload  interface{}
	nextNode *node
}

// Sll : A singly linked list which hols the elements
type Sll struct {
	head      *node
	lastNode  *node
	lastIndex int
}

//Append : Append a node to the end of the list
func (list *Sll) Append(payload interface{}) {
	newNode := new(node)
	newNode.payload = payload
	if list.lastNode == nil {
		list.head, list.lastNode = newNode, newNode
	} else {
		list.lastNode.nextNode = newNode
		list.lastNode = newNode
	}
	list.lastIndex++
}

//Prepend : Prepend a node to the begining of the list
func (list *Sll) Prepend(payload interface{}) {
	newNode := new(node)
	newNode.payload = payload
	newNode.nextNode = list.head
	list.head = newNode
	list.lastIndex++
}

// Remove : Remove node from specified index
func (list *Sll) Remove(index int) (interface{}, bool) {
	if index > list.lastIndex || index < 0 || list.head == nil {
		return nil, false
	}

	if index == 0 {
		temp := list.head
		list.lastIndex--
		if list.head == list.lastNode {
			list.head = nil
			list.lastNode = nil
			return temp, true
		}
		list.head = list.head.nextNode
		return temp, true
	}

	previousNode := list.head

	for i := 1; i < index-1; i++ {
		previousNode = previousNode.nextNode
	}

	temp := previousNode.nextNode
	previousNode.nextNode = temp.nextNode
	list.lastIndex--
	return temp, true
}

// Replace : Replace a node a gven index with a new node
func (list *Sll) Replace(index int, payload interface{}) (interface{}, bool) {
	if index > list.lastIndex || index < 0 || list.head == nil {
		return nil, false
	}

	newNode := new(node)
	newNode.payload = payload

	if index == 0 {
		temp := list.head
		newNode.nextNode = temp.nextNode
		list.head = newNode
		return temp, true
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
	return temp, true
}

// Get : Get a node from specified index
func (list *Sll) Get(index int) (interface{}, bool) {
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

// InitList : Instantiate a new linked list
func InitList(payload interface{}) *Sll {
	list := new(Sll)
	list.head = &node{
		payload: payload,
	}
	list.lastNode = list.head
	return list
}
