package dll

import "fmt"

// Node : A struct describing the node, has Payload and pointers to the next node and previous node in the list
type Node struct {
	Payload      interface{}
	previousNode *Node
	nextNode     *Node
}

// Dll : A doubly linked list which hols the elements
type Dll struct {
	Head      *Node
	LastNode  *Node
	lastIndex int
}

// Append : Append a node to the end of the list
func (list *Dll) Append(payload interface{}) {
	newNode := new(Node)
	newNode.Payload = payload
	if list.LastNode == nil {
		list.Head, list.LastNode = newNode, newNode
	} else {
		list.LastNode.nextNode = newNode
		newNode.previousNode = list.LastNode
		list.LastNode = newNode
	}
	list.lastIndex++
}

// Prepend : Prepend a node to the begining of the list
func (list *Dll) Prepend(payload interface{}) {
	newNode := new(Node)
	newNode.Payload = payload
	if list.Head == nil {
		list.Head, list.LastNode = newNode, newNode
	} else {
		newNode.nextNode = list.Head
		list.Head.previousNode = newNode
		list.Head = newNode
	}
	list.lastIndex++
}

// Remove : Remove node from specified index
func (list *Dll) Remove(index int) (interface{}, bool) {
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
		list.Head.previousNode = nil
		return temp.Payload, true
	}

	if index == list.lastIndex {
		temp := list.LastNode
		list.lastIndex--
		temp.previousNode.nextNode = nil
		return temp.Payload, true
	}

	adjacentNode := list.Head
	var temp *Node
	if index > list.lastIndex/2 {
		adjacentNode = list.LastNode
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
	return temp.Payload, true
}

// Replace : Replace a node a gven index with a new node
func (list *Dll) Replace(index int, payload interface{}) (interface{}, bool) {
	if index > list.lastIndex || index < 0 || list.Head == nil {
		return nil, false
	}

	newNode := new(Node)
	newNode.Payload = payload

	if index == 0 {
		oldNode := list.Head
		newNode.nextNode = oldNode.nextNode
		oldNode.nextNode.previousNode = newNode
		list.Head = newNode
		return oldNode.Payload, true
	}

	if index == list.lastIndex {
		oldNode := list.LastNode
		oldNode.previousNode.nextNode = newNode
		newNode.previousNode = oldNode.previousNode
		list.LastNode = newNode
		return oldNode.Payload, true
	}

	adjacentNode := list.Head
	var oldNode *Node
	if index > list.lastIndex/2 {
		adjacentNode = list.LastNode
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

	return oldNode.Payload, true
}

// InsertAfter : Insert a node after a given node. Returns pointer to newly inserted node, ok
func (list *Dll) InsertAfter(referenceNode *Node, payload interface{}) (*Node, bool) {
	newNode := new(Node)
	newNode.Payload = payload
	newNode.nextNode = referenceNode.nextNode
	newNode.previousNode = referenceNode
	referenceNode.nextNode = newNode
	if referenceNode == list.LastNode {
		list.LastNode = newNode
	}
	list.lastIndex++
	return newNode, true
}

// InsertBefore : Insert a node before a given node. Returns pointer to newly inserted node, ok
func (list *Dll) InsertBefore(referenceNode *Node, payload interface{}) (*Node, bool) {
	newNode := new(Node)
	newNode.Payload = payload
	newNode.nextNode = referenceNode
	newNode.previousNode = referenceNode.previousNode
	referenceNode.previousNode = newNode
	newNode.previousNode.nextNode = newNode

	return newNode, true
}

// Get : Get a node from specified index
func (list *Dll) Get(index int) (*Node, bool) {
	if index > list.lastIndex || index < 0 {
		return nil, false
	}

	if index == 0 {
		return list.Head, true
	}

	if index == list.lastIndex {
		return list.LastNode, true
	}

	adjacentNode := list.Head
	if index > list.lastIndex/2 {
		adjacentNode = list.LastNode
		for i := list.lastIndex - 1; i > index; i-- {
			adjacentNode = adjacentNode.previousNode
		}
		return adjacentNode.previousNode, true
	}
	for i := 1; i < index; i++ {
		adjacentNode = adjacentNode.nextNode
	}
	return adjacentNode.nextNode, true
}

// IndexOf : Finds first occurance given payload in the list, if found returns its index else -1
func (list *Dll) IndexOf(payload interface{}) (int, bool) {
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

// Print : Print the list using fmt.Println, set inReverse to true for reverse order, can be used for debugging
func (list *Dll) Print(inReverse bool) {
	currentNode := list.Head
	if inReverse {
		currentNode = list.LastNode
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
	if list.Head == nil {
		return elementsArray
	}

	if inReverse {
		currentNode := list.LastNode
		for {
			elementsArray = append(elementsArray, currentNode.Payload)
			if currentNode.previousNode == nil {
				break
			}
			currentNode = currentNode.previousNode
		}
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
func (list *Dll) Clear() {
	list.Head = nil
	list.LastNode = nil
}

// InitList : Instantiate a new linked list
func InitList(payload interface{}) *Dll {
	list := new(Dll)
	list.Head = &Node{
		Payload: payload,
	}
	list.LastNode = list.Head
	return list
}
