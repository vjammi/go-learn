package datastructures

import (
	"fmt"
)

type listNode struct {
	val  int
	next *listNode
}

// LinkedList is a linked list
type linkedList struct {
	head   *listNode
	tail   *listNode
	length int
}

// Len Function returns Length of the LinkedList
func (list *linkedList) Len() int {
	return list.length
}

// Add Function inserts a new listNode at the end of the LinkedList
func (list *linkedList) Add(node *listNode) {
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

func (list *linkedList) add(val int) {
	listNode := listNode{val: val}

	if list.head == nil {
		list.head = &listNode
		list.tail = &listNode
		list.length++
	} else {
		list.tail.next = &listNode
		list.tail = &listNode
		list.length++
	}

	return
}

func (list *linkedList) Show() {
	currentNode := list.head
	for currentNode != nil {
		fmt.Printf("%v -> ", list.head.val)
		list.head = list.head.next
	}
	fmt.Println()
}

func (list *linkedList) show() {
	current := list.head

	for current != nil {
		fmt.Println(current.val)
		current = current.next
	}

}

func (list linkedList) ShowFront() (int, error) {
	if list.head == nil {
		return 0, fmt.Errorf("Cannot Find ShowFront Value in an Empty linked list")
	}
	return list.head.val, nil
}

func (list linkedList) ShowBack() (int, error) {
	if list.head == nil {
		return 0, fmt.Errorf("Cannot Find ShowFront Value in an Empty linked list")
	}
	return list.tail.val, nil
}

func (list *linkedList) Reverse() {
	curr := list.head
	list.tail = list.head
	var prev *listNode

	for curr != nil {
		temp := curr.next
		curr.next = prev
		prev = curr
		curr = temp
	}
	list.head = prev
}

func (list *linkedList) Delete(key int) {
	if list.head.val == key {
		list.head = list.head.next
		list.length--
		return
	}

	var prev *listNode = nil
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

	list := linkedList{}
	node1 := &listNode{val: 20}
	node2 := &listNode{val: 30}
	node3 := &listNode{val: 40}
	node4 := &listNode{val: 50}
	node5 := &listNode{val: 70}

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
