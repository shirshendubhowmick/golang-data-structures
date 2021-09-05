package binarytree

func inOrderTraversal(node *Node, payloadSlice *[]interface{}) {
	if node == nil {
		return
	}
	inOrderTraversal(node.leftNode, payloadSlice)
	*payloadSlice = append(*payloadSlice, node.Payload)
	inOrderTraversal(node.rightNode, payloadSlice)
}

func preOrderTraversal(node *Node, payloadSlice *[]interface{}) {
	if node == nil {
		return
	}
	*payloadSlice = append(*payloadSlice, node.Payload)
	preOrderTraversal(node.leftNode, payloadSlice)
	preOrderTraversal(node.rightNode, payloadSlice)
}

func postOrderTraversal(node *Node, payloadSlice *[]interface{}) {
	if node == nil {
		return
	}
	postOrderTraversal(node.leftNode, payloadSlice)
	postOrderTraversal(node.rightNode, payloadSlice)
	*payloadSlice = append(*payloadSlice, node.Payload)
}
