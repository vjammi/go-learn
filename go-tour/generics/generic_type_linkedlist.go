package main

import "fmt"

// Generic types
// In addition to generic functions, Go also supports generic types.
// A type can be parameterized with a type parameter, which could be useful for implementing generic data structures.
// This example demonstrates a simple type declaration for a singly-linked list holding any type of value.
// As an exercise, add some functionality to this list implementation.

// List represents a singly-linked list that holds values of any type.

type ListNode[T any] struct {
	next *ListNode[T]
	val  T
}

type LinkedList[T any] struct {
	head *ListNode[T]
	tail *ListNode[T]
	size int
}

func (list *LinkedList[T]) add(val T) {
	node := new(ListNode[T])
	node.val = val

	if list.head == nil {
		list.head = node
		list.tail = node
		list.size++
	} else {
		list.tail.next = node
		list.tail = node
		list.size++
	}
	return
}

func (list *LinkedList[T]) len() int {
	return list.size
}

func (list *LinkedList[T]) show() {
	var current = list.head //	current = current.next
	for current != nil {
		fmt.Print(current.val, " > ")
		current = current.next
	}
}

func main() {
	listStoringStringValues()
	listStoringIntValues()
}

func listStoringStringValues() {
	list := LinkedList[string]{}

	list.add("blue")
	list.add("green")
	list.add("red")
	list.add("teal")

	fmt.Println("List size ", list.len())
	list.show()
}

func listStoringIntValues() {
	list := LinkedList[int]{}

	list.add(1)
	list.add(2)
	list.add(3)
	list.add(4)

	len := list.len()
	fmt.Println("List size ", len)
	list.show()
}
