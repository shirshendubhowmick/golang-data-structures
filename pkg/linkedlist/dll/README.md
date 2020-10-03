# Doubly Linked List (sll)

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
func InitList(payload interface{}) *Dll
```

Initialize a doubly linked list, accepts a payload of any type for the head node, returns a pointer to the list.

### Methods

#### Append
```go
func (list *Dll) Append(payload interface{}) *Node
```
Append a paylod in the list as a node, returns a pointer to the newly created node

Time complexity:
`O(1)`

<br><br>

#### Prepend
```go
func (list *Dll) Prepend(payload interface{}) *Node
```
Prepend a payload in the list as a node (i.e. change the head node), returns a pointer to the newly created node.

Time complexity:
`O(1)`

<br><br>

#### RemoveByIndex
```go
func (list *Dll) RemoveByIndex(index int) (interface{}, bool)
```
Remove a node at a spcified index, returns the payload of the node removed and an ok value.

0 <= index <= n

Where `n` is the length of the list.

Average Time complexity:
`O(index)`

<br><br>
