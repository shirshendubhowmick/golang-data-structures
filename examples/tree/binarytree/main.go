package main

import (
	"fmt"

	"github.com/shirshendubhowmick/golang-data-structures/pkg/tree/binarytree"
)

func main() {
	tree := binarytree.Init(5)

	node, _ := tree.Root.InsertRight(4)
	node.InsertLeft(8)
	node.InsertRight(9)
	node, _ = tree.Root.InsertLeft(2)
	node.InsertLeft(10)
	node.InsertRight(12)

	fmt.Println(tree.ToSliceInOrder())
	fmt.Println(tree.ToSlicePreOrder())
	fmt.Println(tree.ToSlicePostOrder())
}
