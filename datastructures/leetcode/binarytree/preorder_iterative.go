package main

import (
	"fmt"
	"github.com/vjammi/go-learn/datastructures/stack"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	p := root
	ret := []int{}
	stack := []*TreeNode{}
	for p != nil || len(stack) != 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		if len(stack) != 0 {
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			p = tmp.Right
			ret = append(ret, tmp.Val)
		}
	}
	return ret
}

func inorderIterative(root *TreeNode) []int {
	ret := make([]int, 0)
	stack := make([](*TreeNode), 0)

	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ret = append(ret, last.Val)
			root = last.Right
		}
	}
	return ret
}

//func inorderRecursive_foo(root *TreeNode) []int {
//	if root == nil {
//		return []int{}
//	}
//
//	arr := make([]int, 0)
//	recursive := inorderRecursive_foo(root.Left)
//	arr = append(arr, recursive...)
//
//	arr = append(arr, root.Val)
//	ints := inorderRecursive_foo(root.Right)
//	arr = append(arr, ints...)
//
//	return arr
//}

func preorderIterative(node *TreeNode) []int {
	//inOrderList := new(list.LinkedList)
	//inOrderList.AddNode(20)
	//inOrderList.AddNode(30)
	//inOrderList.Show()
	//
	//stack := new(stack.Stack)
	//stack.Push(1)
	//stack.Push(2)
	//stack.Push(3)
	//stack.Show()

	i := 0
	result := make([](int), 0)
	//stack := make([](*TreeNode), 0)
	stack := new(stack.Stack)
	//stack = append(stack, node)
	stack.Push(node)

	for !stack.IsEmpty() { // || poppedNode != nil
		//fmt.Print(node.Val, " > ")
		poppedNode := stack.Pop()
		//poppedNode := stack.Pop()
		fmt.Println(poppedNode)
		//result = append(result, poppedNode)
		//result[i] = poppedNode.Val
		i++
		if poppedNode != nil {
			stack.Push(poppedNode)
			//	//stack = append(stack, poppedNode.Right)
		}
		if poppedNode != nil {
			//stack = append(stack, poppedNode.Left)
			stack.Push(poppedNode)
		}
	}
	return result
}

var size = 0

func preOrderRecursive(node *TreeNode, result []int) {
	if node == nil {
		return
	}

	fmt.Print(node.Val, " > ")
	result[size] = node.Val
	size++ //TODO
	preOrderRecursive(node.Left, result)
	preOrderRecursive(node.Right, result)
}

func main() {
	root := populateTree()

	//var result [7]int
	// or
	result := make([]int, 7)
	preOrderRecursive(root, result[:])
	fmt.Println(result)

	result2 := preorderIterative(root)
	fmt.Println(result2)

	preOrderRecursive(root, result[:])
	fmt.Println(result)
	//iterativeResult := inorderIterative(root)
	//fmt.Println(iterativeResult)
	//recursiveResult := inorderRecursive_foo(root)
	//fmt.Println(recursiveResult)

}

func populateTree() *TreeNode {

	root := new(TreeNode)
	root.Val = 16

	// Populate left side of the root
	//left := new(TreeNode)
	root.Left = new(TreeNode)
	root.Left.Val = 10

	root.Left.Left = new(TreeNode)
	root.Left.Left.Val = 8

	root.Left.Right = new(TreeNode)
	root.Left.Right.Val = 12

	// Populate right side of the root
	//right := new(TreeNode)
	root.Right = new(TreeNode)
	root.Right.Val = 22

	//right := new(TreeNode)
	root.Right.Left = new(TreeNode)
	root.Right.Left.Val = 20

	root.Right.Right = new(TreeNode)
	root.Right.Right.Val = 24

	return root
}
