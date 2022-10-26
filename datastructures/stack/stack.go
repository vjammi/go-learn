package stack

// Source: https://github.com/TheAlgorithms/Go/blob/4d9b1e48789f24959739ef03b52b26d074a98de8/structure/stack/stack_test.go
import (
	"fmt"
)

// Node structure
type TreeNode struct {
	Val  any
	Next *TreeNode
}

// Stack has just head of node and with length
type Stack struct {
	head   *TreeNode
	length int
}

// push add value to last index
func (stack *Stack) Push(n any) {
	newNode := &TreeNode{} // new node

	newNode.Val = n
	newNode.Next = stack.head

	stack.head = newNode
	stack.length++
}

// pop remove last item as first output
func (stack *Stack) Pop() any {
	//node := stack.head.Val
	node := stack.head
	if stack.head.Next == nil {
		stack.head = nil
	} else {
		//stack.head.Val, stack.head.Next = stack.head.Next.Val, stack.head.Next.Next
		stack.head, stack.head.Next = stack.head.Next, stack.head.Next.Next
	}

	stack.length--
	return node
}

// isEmpty to check our array is empty or not
func (stack *Stack) IsEmpty() bool {
	return stack.length == 0
}

// len use to return length of our stack
func (stack *Stack) Len() int {
	return stack.length
}

// peak return last input value
func (stack *Stack) Peak() any {
	node := stack.head
	return node
}

// show all value as an interface array
func (stack *Stack) Show() (in []any) {
	current := stack.head

	for current != nil {
		in = append(in, current.Val)
		fmt.Print(current.Val, " > ")
		current = current.Next
	}
	return
}

//func main() {
//	//var stack Stack
//	//   or
//	stack := new(Stack)
//
//	stack.Push(1)
//	stack.Push(2)
//	stack.Push(3)
//	stack.Push(4)
//	fmt.Println("peek ", stack.Peak())
//	stack.Push(5)
//	stack.Push(6)
//	stack.Push(7)
//	fmt.Println("peek ", stack.Peak())
//	stack.Show()
//}
