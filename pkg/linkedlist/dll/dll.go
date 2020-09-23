package dll

import "fmt"

type node struct {
	payload      interface{}
	previousNode *node
	nextNode     *node
}

// Dll : A doubly linked list which hols the elements
type Dll struct {
	head      *node
	lastNode  *node
	lastIndex int
}

// Append : Append a node to the end of the list
func (list *Dll) Append(payload interface{}) {
	newNode := new(node)
	newNode.payload = payload
	if list.lastNode == nil {
		list.head, list.lastNode = newNode, newNode
	} else {
		list.lastNode.nextNode = newNode
		newNode.previousNode = list.lastNode
		list.lastNode = newNode
	}
	list.lastIndex++
}

// Prepend : Prepend a node to the begining of the list
func (list *Dll) Prepend(payload interface{}) {
	newNode := new(node)
	newNode.payload = payload
	newNode.nextNode = list.head
	list.head.previousNode = newNode
	list.head = newNode
	list.lastIndex++
}

// Remove : Remove node from specified index
func (list *Dll) Remove(index int) (interface{}, bool) {
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
		list.head.previousNode = nil
		return temp.payload, true
	}

	if index == list.lastIndex {
		temp := list.lastNode
		list.lastIndex--
		temp.previousNode.nextNode = nil
		return temp.payload, true
	}

	adjacentNode := list.head
	var temp *node
	if index > list.lastIndex/2 {
		adjacentNode = list.lastNode
		for i := list.lastIndex - 1; i > index; i-- {
			adjacentNode = adjacentNode.previousNode
		}
		temp = adjacentNode.previousNode
		temp.previousNode.nextNode = adjacentNode
		adjacentNode.previousNode = temp.previousNode
	} else {
		for i := 1; i < index; i++ {
			adjacentNode = adjacentNode.nextNode
		}
		temp = adjacentNode.nextNode
		adjacentNode.nextNode = temp.nextNode
		temp.nextNode.previousNode = adjacentNode
	}

	list.lastIndex--
	return temp.payload, true
}

// Replace : Replace a node a gven index with a new node
func (list *Dll) Replace(index int, payload interface{}) (interface{}, bool) {
	if index > list.lastIndex || index < 0 || list.head == nil {
		return nil, false
	}

	newNode := new(node)
	newNode.payload = payload

	if index == 0 {
		oldNode := list.head
		newNode.nextNode = oldNode.nextNode
		oldNode.nextNode.previousNode = newNode
		list.head = newNode
		return oldNode.payload, true
	}

	if index == list.lastIndex {
		oldNode := list.lastNode
		oldNode.previousNode.nextNode = newNode
		newNode.previousNode = oldNode.previousNode
		list.lastNode = newNode
		return oldNode.payload, true
	}

	adjacentNode := list.head
	var oldNode *node
	if index > list.lastIndex/2 {
		adjacentNode = list.lastNode
		for i := list.lastIndex - 1; i > index; i-- {
			adjacentNode = adjacentNode.previousNode
		}
		oldNode = adjacentNode.previousNode
		adjacentNode.previousNode = newNode
		newNode.nextNode = adjacentNode
		newNode.previousNode = oldNode.previousNode
		oldNode.previousNode.nextNode = newNode
	} else {
		for i := 1; i < index; i++ {
			adjacentNode = adjacentNode.nextNode
		}
		oldNode = adjacentNode.nextNode
		adjacentNode.nextNode = newNode
		newNode.previousNode = adjacentNode
		newNode.nextNode = oldNode.nextNode
		oldNode.nextNode.previousNode = newNode
	}

	return oldNode.payload, true
}

// Get : Get a node from specified index
func (list *Dll) Get(index int) (interface{}, bool) {
	if index > list.lastIndex || index < 0 {
		return nil, false
	}

	if index == 0 {
		return list.head.payload, true
	}

	if index == list.lastIndex {
		return list.lastNode.payload, true
	}

	adjacentNode := list.head
	if index > list.lastIndex/2 {
		adjacentNode = list.lastNode
		for i := list.lastIndex - 1; i > index; i-- {
			adjacentNode = adjacentNode.previousNode
		}
		return adjacentNode.previousNode.payload, true
	}
	for i := 1; i < index; i++ {
		adjacentNode = adjacentNode.nextNode
	}
	return adjacentNode.nextNode.payload, true
}

// IndexOf : Finds first occurance given payload in the list, if found returns its index else -1
func (list *Dll) IndexOf(payload interface{}) (int, bool) {
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

// Print : Print the list using fmt.Println, set inReverse to true for reverse order, can be used for debugging
func (list *Dll) Print(inReverse bool) {
	currentNode := list.head
	if inReverse {
		currentNode = list.lastNode
	}

	if currentNode == nil {
		fmt.Println("Empty list")
		return
	}

	for {
		fmt.Println(currentNode)
		if inReverse {
			if currentNode.previousNode == nil {
				break
			}
			currentNode = currentNode.previousNode
		} else {
			if currentNode.nextNode == nil {
				break
			}
			currentNode = currentNode.nextNode
		}
	}
}

// Length : Get the current length of the list
func (list *Dll) Length() int {
	return list.lastIndex + 1
}

// ToSlice : Returns the elements present in the list as an array, set inReverse to true for reverse order
func (list *Dll) ToSlice(inReverse bool) []interface{} {
	elementsArray := []interface{}{}
	if list.head == nil {
		return elementsArray
	}

	if inReverse {
		currentNode := list.lastNode
		for {
			elementsArray = append(elementsArray, currentNode.payload)
			if currentNode.previousNode == nil {
				break
			}
			currentNode = currentNode.previousNode
		}
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
func (list *Dll) Clear() {
	list.head = nil
	list.lastNode = nil
}

// InitList : Instantiate a new linked list
func InitList(payload interface{}) *Dll {
	list := new(Dll)
	list.head = &node{
		payload: payload,
	}
	list.lastNode = list.head
	return list
}
