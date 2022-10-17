package main

// Source: https://github.com/TheAlgorithms/Go/blob/4d9b1e48789f24959739ef03b52b26d074a98de8/structure/stack/stack_test.go
import (
	"fmt"
)

// Node structure
type Node struct {
	Val  any
	Next *Node
}

// Stack has just head of node and with length
type Stack struct {
	head   *Node
	length int
}

// push add value to last index
func (stack *Stack) push(n any) {
	newNode := &Node{} // new node

	newNode.Val = n
	newNode.Next = stack.head

	stack.head = newNode
	stack.length++
}

// pop remove last item as first output
func (stack *Stack) pop() any {
	result := stack.head.Val
	if stack.head.Next == nil {
		stack.head = nil
	} else {
		stack.head.Val, stack.head.Next = stack.head.Next.Val, stack.head.Next.Next
	}

	stack.length--
	return result
}

// isEmpty to check our array is empty or not
func (stack *Stack) isEmpty() bool {
	return stack.length == 0
}

// len use to return length of our stack
func (stack *Stack) len() int {
	return stack.length
}

// peak return last input value
func (stack *Stack) peak() any {
	return stack.head.Val
}

// show all value as an interface array
func (stack *Stack) show() (in []any) {
	current := stack.head

	for current != nil {
		in = append(in, current.Val)
		fmt.Println(current.Val)
		current = current.Next
	}
	return
}

func main() {
	//var stack Stack
	//   or
	stack := new(Stack)

	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.push(4)
	fmt.Println("peek ", stack.peak())
	stack.push(5)
	stack.push(6)
	stack.push(7)
	fmt.Println("peek ", stack.peak())
	stack.show()
}
