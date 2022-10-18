package bst

import (
	"fmt"
)

type TreeNode struct {
	val         int
	left, right *TreeNode
}

type BinarySearchTree struct {
	root *TreeNode
}

func New() *BinarySearchTree {
	bst := new(BinarySearchTree)
	return bst
}

func (bst *BinarySearchTree) PreOrder(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.val, " > ")
	bst.PreOrder(node.left)
	bst.PreOrder(node.right)
}

func (bst *BinarySearchTree) Insert(node *TreeNode, val int) *TreeNode {
	if node == nil {
		newNode := &TreeNode{val, nil, nil}
		return newNode
	}

	if val < node.val {
		node.left = bst.Insert(node.left, val)
	} else if val > node.val {
		node.right = bst.Insert(node.right, val)
	} else {
		node.val = val
		return node
	}
	return node
}

//func main() {
//	//bst := new(BinarySearchTree)
//	// or
//	bst := New()
//
//	root := bst.root
//	root = bst.Insert(root, 20)
//	root = bst.Insert(root, 15)
//	root = bst.Insert(root, 25)
//	root = bst.Insert(root, 10)
//	root = bst.Insert(root, 30)
//	bst.PreOrder(root)
//}
