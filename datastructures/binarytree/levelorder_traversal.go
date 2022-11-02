package main

import (
	"fmt"
	"sort"
	"strings"
)

// queueArrayToString is used to convert queue array to string
func queueArrayToString(q []*Queue) string {
	output := "["
	for i, queue := range q {
		output += queue.String() + "]"
		if i < len(q)-1 {
			output += ", "
		}
	}
	return output
}

// levelOrderTraversal is the solution function
func levelOrderTraversal(root *BinaryTreeNode) {
	// We print nil if the root is nil
	if root == nil {
		fmt.Print("nil")
	}

	fmt.Print("\n\tInitializing the queues\n")

	// Declaring an array of two queues
	queues := make([]*Queue, 0)
	queues = append(queues, &Queue{nil, nil, 0})
	queues = append(queues, &Queue{nil, nil, 0})
	fmt.Printf("\t\tqueues array: %v\n", queueArrayToString(queues))

	// Initializing the current and next queues
	currentQueue, nextQueue := queues[0], queues[1]
	fmt.Printf("\t\tCurrent queue: %v\n", currentQueue.String())
	fmt.Printf("\t\tNext queue: %v\n", nextQueue.String())

	// Enqueuing the root node into the current queue and setting
	// level to zero
	fmt.Print("\t\tEnqueuing the root node into the current queue\n")
	currentQueue.Enqueue(root)
	fmt.Printf("\t\t\tUpdated current queue: %v\n", currentQueue.String())
	levelNumber := 0
	fmt.Printf("\t\t\tLevel Number: %v\n", levelNumber)

	return

}

/*
	reverseArray() is a helper function used by main

to reverse an integer array
*/
func reverseArray(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

// main is used for driver code
func main() {
	// Creating a binary tree
	input1 := []int{100, 50, 200, 25, 75, 350}
	tree1 := new(BinaryTree)
	tree1.BinaryTreeInitWithSlice(input1)

	// Creating a right degenerate binary tree
	input2 := input1
	sort.Ints(input2)
	tree2 := new(BinaryTree)
	tree2.BinaryTreeInitWithSlice(input2)

	// Creating a left degenerate binary tree
	input2 = reverseArray(input2)
	tree3 := new(BinaryTree)
	tree3.BinaryTreeInitWithSlice(input2)

	// Creating a single node binary tree
	tree4 := new(BinaryTree)
	tree4.BinaryTreeInitWithData(100)

	roots := []*BinaryTreeNode{tree1.root, tree2.root, tree3.root, tree4.root, nil}

	for i := range roots {
		fmt.Printf("%d.\tBinary Tree:\n", i+1)
		//displayTree(roots[i])
		// Printing the in-order list using the method we just implemented
		fmt.Print("\n\tLevel order traversal:\t")
		levelOrderTraversal(roots[i])
		fmt.Printf("\n%s\n", strings.Repeat("-", 100))
	}
}
