package main

import (
	"fmt"
	"github.com/shirshendubhowmick/golang-data-structures/pkg/linkedlist/sll"
)

func main() {
	list := sll.InitList(4)

	list.Append(5)
	list.Append(6)
	nodeTobeReplace := list.Append(7)
	list.Append(8)

	list.Prepend(3)
	list.Prepend(2)
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

	list.Print()

	fmt.Println(list.ReplaceByNode(nodeTobeReplace, 1000))
	list.Print()
}
