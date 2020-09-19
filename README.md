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
	Sll "github.com/shirshendubhowmick/golang-data-structures/src/linkedlist/sll"
)

func main() {
	list := Sll.InitList(1)

	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)
	list.Prepend(6)
	list.Prepend(7)

	list.Print()

	list.Remove(0)
	list.Remove(4)
	list.Remove(5)

	list.Replace(0, 100)
	list.Clear()
  
	list.Append(2)
	list.Append(5)
	list.Prepend(6)
	list.Prepend(7)

	list.Print()
}
```

## Installation
This was developved in go `v1.15` and not tested in any other lower version as of now.

Inside your main module run

```bash
go get github.com/shirshendubhowmick/golang-data-structures/src
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

`examples/linkedlist` directory contains all examples related to linked lists

## Package contents

* Linked List
  * Singly linked list
  * Doubly linked list
  * Circular linked list

* Stack
  * Array stack
  * Linked List stack

* Queue
  * Normal queue
  * Priority queue
  * Deque
  * Circular queue

Many more coming soon...



#### Note: Detailed documentation WIP
