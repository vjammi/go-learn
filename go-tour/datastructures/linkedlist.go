package datastructures

import (
	"fmt"
)

type ListNode struct {
	val  int
	next *ListNode
}

// LinkedList is a linked list
type Linkedlist struct {
	head   *ListNode
	tail   *ListNode
	length int
}

// Len Function returns Length of the LinkedList
func (list *Linkedlist) Len() int {
	return list.length
}

// Add Function inserts a new ListNode at the end of the LinkedList
func (list *Linkedlist) Add(node *ListNode) {
	if list.head == nil {
		list.head = node
		list.tail = node
		list.length++
	} else {
		list.tail.next = node
		list.tail = node
		list.length++
	}
}

func (list Linkedlist) Show() {
	for list.head != nil {
		fmt.Printf("%v -> ", list.head.val)
		list.head = list.head.next
	}
	fmt.Println()
}

func (list Linkedlist) ShowFront() (int, error) {
	if list.head == nil {
		return 0, fmt.Errorf("Cannot Find ShowFront Value in an Empty linked list")
	}
	return list.head.val, nil
}

func (list Linkedlist) ShowBack() (int, error) {
	if list.head == nil {
		return 0, fmt.Errorf("Cannot Find ShowFront Value in an Empty linked list")
	}
	return list.tail.val, nil
}

func (list *Linkedlist) Reverse() {
	curr := list.head
	list.tail = list.head
	var prev *ListNode

	for curr != nil {
		temp := curr.next
		curr.next = prev
		prev = curr
		curr = temp
	}
	list.head = prev
}

func (list *Linkedlist) Delete(key int) {
	if list.head.val == key {
		list.head = list.head.next
		list.length--
		return
	}

	var prev *ListNode = nil
	curr := list.head
	for curr != nil && curr.val != key {
		prev = curr
		curr = curr.next
	}
	if curr == nil {
		fmt.Println("Key Not found")
		return
	}
	prev.next = curr.next
	list.length--
	fmt.Println("Node Deleted")

}

func LinkedListTest3() {

	list := Linkedlist{}
	node1 := &ListNode{val: 20}
	node2 := &ListNode{val: 30}
	node3 := &ListNode{val: 40}
	node4 := &ListNode{val: 50}
	node5 := &ListNode{val: 70}

	list.Add(node1)
	list.Add(node2)
	list.Add(node3)
	list.Add(node4)
	list.Add(node5)
	fmt.Println("Length = ", list.Len())

	list.Show()
	list.Delete(40)
	list.Reverse()
	fmt.Println("Length = ", list.Len())
	list.Show()

	front, _ := list.ShowFront()
	back, _ := list.ShowBack()
	fmt.Println("ShowFront = ", front)
	fmt.Println("ShowBack = ", back)

}
