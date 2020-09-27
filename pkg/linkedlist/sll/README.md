# Singly Linked List (sll)

## Contents

* [**Structs**](#Structs)
  * Node
  * Sll
* [**Functions**](#Functions)
  * [InitList](#InitList)
* [**Methods**](#Methods)
  * [Append](#Append)
  * [Prepend](#Prepend)
  * [RemoveByIndex](#RemoveByIndex)
  * [RemoveByNode](#RemoveByNode)
  * [ReplaceByIndex](#ReplaceByIndex)
  * [ReplaceByNode](#ReplaceByNode)
  * [InsertAt](#InsertAt)
  * [InsertAfter](#InsertAfter)
  * [InsertBefore](#InsertBefore)
  * [Get](#Get)
  * [IndexOf](#IndexOf)
  * [Print](#Print)
  * [Length](#Length)
  * [ToSlice](#ToSlice)
  * [Clear](#Clear)

### Functions

#### InitList

```go
func InitList(payload interface{}) *Sll
```

Initialize a singly linked list, accepts a payload of any type for the head node, returns a pointer to the list.

### Methods

#### Append
```go
func (list *Sll) Append(payload interface{}) *Node
```
Append a paylod in the list as a node, returns a pointer to the newly created node

Time complexity:
`O(1)`

<br><br>

#### Prepend
```go
func (list *Sll) Prepend(payload interface{}) *Node
```
Prepend a payload in the list as a node (i.e. change the head node), returns a pointer to the newly created node.

Time complexity:
`O(1)`

<br><br>

#### RemoveByIndex
```go
func (list *Sll) RemoveByIndex(index int) (interface{}, bool)
```
Remove a node at a spcified index, returns the payload of the node removed and an ok value.

0 <= index <= n

Where `n` is the length of the list.

Average Time complexity:
`O(index)`

<br><br>

#### RemoveByNode
```go
func (list *Sll) RemoveByNode(nodeToBeRemoved *Node) (interface{}, bool)
```
Remove a node by its reference, returns the payload of the node removed and an ok value.

Average time complexity:
`O(n)`

Where `n` is the length of the list.

<br><br>

#### ReplaceByIndex
```go
func (list *Sll) ReplaceByIndex(index int, payload interface{}) (interface{}, *Node, bool)
```
Replace a node at the specified with a new node created with the supplied payload. Returns the payload of the old node, a pointer to the new
node and an ok value.

0 <= index <= n

Where `n` is the length of the list.

Average time complexity:
`O(index)`

<br><br>

#### ReplaceByNode
```go
func (list *Sll) ReplaceByNode(nodeToBeReplaced *Node, payload interface{}) (interface{}, *Node, bool)
```
Replace a node specified with a pointer with a new node created with the supplied payload. Returns the payload of the old node, a pointer to the new
node and an ok value

Average time complexity:
`O(n)`

Where `n` is the length of the list.

<br><br>

#### InsertAt
```go
func (list *Sll) InsertAt(index int, payload interface{}) (*Node, bool)
```
Insert a new node created with the supplied payload at a specific index. Any exisiting node at that index become the
next node of the newly created node. Returns a pointer to the newly created node and an ok value.

0 <= index <= n

Average time complexity:
`O(index)`

Where `n` is the length of the list.

<br><br>

#### InsertAfter 
```go
func (list *Sll) InsertAfter(referenceNode *Node, payload interface{}) (*Node, bool)
```

<br><br>

#### InsertBefore
```go
func (list *Sll) InsertBefore(referenceNode *Node, payload interface{}) (*Node, bool)
```

<br><br>

#### Get
```go
func (list *Sll) Get(index int) (*Node, bool)
```

<br><br>

#### IndexOf
```go
func (list *Sll) IndexOf(payload interface{}) (int, bool)
```

<br><br>

#### Print
```go
func (list *Sll) Print()
```

<br><br>

#### Length
```go
func (list *Sll) Length() int
```

<br><br>

#### ToSlice
```go
func (list *Sll) ToSlice() []interface{}
```

<br><br>

#### Clear
```go
func (list *Sll) Clear()
```