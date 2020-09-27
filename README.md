# Golang Data Structures

Packages of commonly used data structures and algorithms. Supports go module.

## Table of contents
* [Quickstart](#Quickstart)
* [Installation](#Installation)
* [Directory structure](#Directory-structure)
* [Package contents](#Package-contents)

## Quickstart

```go
package main

import (
	"fmt"
	"github.com/shirshendubhowmick/golang-data-structures/pkg/linkedlist/sll"
)

func main() {
	list := sll.InitList(4)

	list.Append(5)
	list.Append(6)
	list.Append(7)
	list.Append(8)

	list.Prepend(3)
	list.Prepend(2)
	list.Prepend(1)
	list.Prepend(0)

	list.Print()

	list.Remove(4)

	list.Print()

	fmt.Println(list.Get(4))

	fmt.Println(list.IndexOf(8))
	list.Replace(7, 99)
	list.Print()
	fmt.Println(list.ToSlice())
	fmt.Println(list.Length())
	list.Clear()

	list.Print()
}
```

## Installation
This was developved in go `v1.15` and not tested in any other lower version as of now.

Inside your main module run

```bash
go get github.com/shirshendubhowmick/golang-data-structures/pkg
```


## Directory structure

This repo uses go modules, `pkg` is the module directory and it contains all the packages.
Packages are distributed based on their data structure.
For example

`pkg/linkedlist` contains packages `sll`, `dll` etc

`pkg/tree` contains packages `btree`, `avltree` etc

Any packages inside an `internal` directory are meant for internal use, do not consume them directly.

<br>
Examples are located in `examples` directory, they are also distributed in the same way as packages.
For example

`examples/linkedList` directory contains all examples related to linked lists

## Package contents

* [**Linked lists**](#Linked-lists)
  * [Singly linked list](#Singly-linked-list)
  * [Doubly linked list](#Doubly-linked-list)
  * Circular linked list

* **Stack**
  * Array stack
  * Linked List stack

* **Queue**
  * Normal queue
  * Priority queue
  * Deque
  * Circular queue

Many more coming soon...


## Linked lists
All linked lists uses type `interface {}` for the payload, as the list allows any type of data to be stored and also internally it doesn't perform any operations on the payload.

It is recommended that consumer of linked lists define a proper type to the payload, to have compile time type safety. As it is expected that consumers are going to perform operations on the payload.

* ### Singly linked list
  Package location:
 `github.com/shirshendubhowmick/golang-data-structures/pkg/linkedlist/sll`

  [See examples](examples/linkedList/singlyLinkedList/)
	
	[Package README](pkg/linkedlist/sll)


* ### Doubly linked list
  Package location:
 `github.com/shirshendubhowmick/golang-data-structures/pkg/linkedlist/dll`

  [See examples](examples/linkedList/doublyLinkedList/)
	
	[Package README](pkg/linkedlist/sll)
  
  

#### Note: Detailed documentation WIP
