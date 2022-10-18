package main

import (
	"fmt"
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

	current := new(bst.BinarySearchTree)
	fmt.Println(current)

}

func main() {
	preorder()
}
