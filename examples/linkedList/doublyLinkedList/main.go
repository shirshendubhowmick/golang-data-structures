package main

import (
	"fmt"
	"github.com/shirshendubhowmick/golang-data-structures/pkg/linkedlist/dll"
)

func main() {
	list := dll.InitList(4)

	list.Append(5)
	list.Append(6)
	list.Append(7)
	list.Append(8)

	list.Prepend(3)
	list.Prepend(2)
	list.Prepend(1)
	list.Prepend(0)

	list.Print(false)
	list.Print(true)

	list.Remove(4)

	list.Print(false)

	fmt.Println(list.Get(4))

	fmt.Println(list.IndexOf(8))
	list.Replace(7, 99)
	list.Print(false)
	fmt.Println(list.ToSlice(false))
	fmt.Println(list.ToSlice(true))
	fmt.Println(list.Length())
	list.Clear()

	list.Print(true)
}
