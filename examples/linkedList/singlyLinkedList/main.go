package main

import (
	"fmt"

	"github.com/shirshendubhowmick/golang-data-structures/pkg/linkedlist/sll"
)

func main() {
	list := sll.InitList(4)

	list.Append(5)
	nodeRef0 := list.Append(6)
	list.Append(7)
	list.Append(8)

	list.Prepend(3)
	nodeRef1 := list.Prepend(2)
	list.Prepend(1)
	list.Prepend(0)

	list.Print()

	list.RemoveByIndex(4)

	list.Print()

	fmt.Println(list.Get(4))

	fmt.Println(list.IndexOf(8))

	list.ReplaceByIndex(7, 99)
	list.Print()

	fmt.Println(list.ToSlice())
	fmt.Println(list.Length())

	fmt.Println(list.InsertAt(8, 200))
	list.Print()

	_, nodeRef2, _ := list.ReplaceByNode(nodeRef1, 999)
	list.Print()

	list.RemoveByNode(nodeRef2)
	list.Print()

	list.InsertAfter(nodeRef0, 500)
	list.Print()

	list.InsertBefore(nodeRef0, 600)
	list.Print()

}
