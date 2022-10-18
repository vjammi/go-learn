package main

import (
	"github.com/vjammi/go-learn/datastructures/bst"
	list "github.com/vjammi/go-learn/datastructures/linkedlist"
	"github.com/vjammi/go-learn/datastructures/stack"
)

func preorder() {
	inOrderList := new(list.LinkedList)
	inOrderList.AddNode(20)
	inOrderList.AddNode(30)
	inOrderList.Show()

	stack := new(stack.Stack)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Show()

}

func main() {
	bst := new(bst.BinarySearchTree)

	root = bst.Insert(root, 10)
	root = bst.Insert(root, 5)
	root = bst.Insert(root, 15)

	bst.PreOrder(root)
	preorder()
}
