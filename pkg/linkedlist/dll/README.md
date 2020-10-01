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