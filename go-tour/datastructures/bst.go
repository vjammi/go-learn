package main

import (
	"fmt"
)

type treeNode struct {
	val         int
	left, right *treeNode
}

type binarySearchTree struct {
	root *treeNode
}

func (bst *binarySearchTree) preOrder(node *treeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.val, " > ")
	bst.preOrder(node.left)
	bst.preOrder(node.right)
}

func (bst *binarySearchTree) insert(node *treeNode, val int) *treeNode {
	if node == nil {
		node := new(treeNode)
		node.val = val
		return node
	}

	if val < node.val {
		node.left = bst.insert(node.left, val)
	} else if val > node.val {
		node.right = bst.insert(node.right, val)
	} else {
		node.val = val
		return node
	}
	return node
}

func main() {
	bst := new(binarySearchTree)

	root := bst.root
	root = bst.insert(root, 10)
	root = bst.insert(root, 5)
	root = bst.insert(root, 15)

	bst.preOrder(root)
}
