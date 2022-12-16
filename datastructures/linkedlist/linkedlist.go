package linkedlist

import (
	"fmt"
)

type ListNode struct { // Element
	val  int
	next *ListNode
}

type LinkedList struct { // List
	head   *ListNode
	tail   *ListNode
	length int
}

func New() *LinkedList {
	//var list *LinkedList = new(LinkedList)
	//return list

	list := new(LinkedList)
	return list // Will return a parameter type of *LinkedList
}

func (list *LinkedList) AddNode(val int) {
	// node := ListNode{}
	// node := &ListNode{val: 1}
	// var node *ListNode = new(ListNode)

	node := new(ListNode)
	node.val = val

	if list.head == nil {
		list.head = node
		list.tail = node
		list.length++
	} else {
		list.tail.next = node
		list.tail = node
		list.length++
	}
	return
}

func (list *LinkedList) Len() int {
	return list.length
}

// Note: When we have a mix of (list *LinkedList)  &&  (list LinkedList)
// Go Runtime complains Struct LinkedList has methods on both value and pointer receivers. Such usage is not recommended by the Go Documentation.
func (list *LinkedList) Show() {
	// current := list.head		//	current = current.next
	// or
	var current = list.head //	current = current.next

	for current != nil {
		fmt.Print(current.val, ">")
		current = current.next
	}
}

func (list *LinkedList) ShowFront() error {
	if list.head == nil {
		return fmt.Errorf("LinkedList is empty...")
	}
	fmt.Println(list.head.val)
	return nil
}

func (list *LinkedList) ShowBack() (int, error) {
	if list.tail == nil {
		return -1, fmt.Errorf("LinkedList is empty...")
	}
	fmt.Println(list.tail.val)
	return list.tail.val, nil
}

func (list *LinkedList) Reverse() { // list:     *main.LinkedList
	var current = list.head // current:  *main.ListNode
	var previous *ListNode  // previous: nil

	var newTail = list.head
	for current != nil {
		fmt.Print(current.val, ">")
		var next = current.next

		current.next = previous

		previous = current
		current = next
	}

	list.tail = newTail
	list.head = previous
}

func (list *LinkedList) Delete(val int) {
	var current = list.head
	var previous *ListNode

	for current.next != nil {
		fmt.Print(current.val, " > ")
		var next = current.next

		if current.val == val {
			previous.next = next
			return
		}

		previous = current
		current = current.next
	}
}
