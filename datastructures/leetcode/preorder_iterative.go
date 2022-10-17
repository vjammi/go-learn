package main

import (
	"fmt"
	"github.com/vjammi/go-learn/datastructures/stack"
)

func preorder() {
	stack := new(stack.Stack)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	stack.Show()

	fmt.Println(stack.Peak())
}

func main() {
	preorder()
}
