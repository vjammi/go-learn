package main

import (
	"fmt"
)

type LinkedList struct {
	head   *ListNode
	length int
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list := New()
	list.populateList()

	list.show()
	list.reverseList()
	list.show()
	list.head = reverseList(list.head)
	list.show()
}

func New() *LinkedList {
	var list *LinkedList = new(LinkedList)
	return list
}

/**
 *                   head
 *                     |
 *               NIL < 1 < 2 > 3 > 4 > 5 > NIL
 *                                     P   C
 */
func (list *LinkedList) reverseList() {
	var curr *ListNode
	curr = list.head
	var prev *ListNode
	var next *ListNode

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	list.head = prev
}

func reverseList(head *ListNode) *ListNode {
	var curr *ListNode
	var prev *ListNode
	var next *ListNode
	curr = head

	for curr != nil {
		next = curr.Next

		curr.Next = prev

		prev = curr
		curr = next
	}

	head = prev
	return head
}

func (list *LinkedList) populateList() {
	// populateList
	var node0 = new(ListNode)
	var node1 = new(ListNode)
	var node2 = new(ListNode)
	var node3 = new(ListNode)
	var node4 = new(ListNode)

	node0.Val = 0
	list.head = node0
	node0.Next = node1

	node1.Val = 1
	node1.Next = node2

	node2.Val = 2
	node2.Next = node3

	node3.Val = 3
	node3.Next = node4

	node4.Val = 4
	node4.Next = nil
}

func (list *LinkedList) show() {
	var curr = list.head
	for curr != nil {
		fmt.Print(curr.Val, ", ")
		curr = curr.Next
	}
	fmt.Println()
}
