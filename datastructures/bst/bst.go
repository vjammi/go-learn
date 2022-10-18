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
		node := new(TreeNode)
		node.val = val
		return node
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
//	bst := new(BinarySearchTree)
//
//	root := bst.root
//	root = bst.Insert(root, 10)
//	root = bst.Insert(root, 5)
//	root = bst.Insert(root, 15)
//
//	bst.PreOrder(root)
//}
