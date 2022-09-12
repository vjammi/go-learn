package main

import (
	"fmt"
)

type listNode1 struct {
	val  int
	next *listNode
}

// LinkedList1 is a linked list
type linkedList1 struct {
	head   *listNode
	tail   *listNode
	length int
}

// Len Function returns Length of the LinkedList
func (list *linkedList1) Len() int {
	return list.length
}

// Add Function inserts a new listNode at the end of the LinkedList
func (list *linkedList1) Add(node *listNode) {
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

func (list *linkedList1) Delete(key int) {
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

func (list *linkedList1) Show() {
	currentNode := &list.head
	for currentNode != nil {
		fmt.Printf("%v -> ", list.head.val)
		list.head = list.head.next
	}
	fmt.Println()
}

func (list *linkedList1) ShowFront() (int, error) {
	if list.head == nil {
		return 0, fmt.Errorf("Cannot Find ShowFront Value in an Empty linked list")
	}
	return list.head.val, nil
}

func (list *linkedList1) Reverse() {
	list.tail = list.head

	current := list.head
	var prev *listNode
	for current != nil {
		temp := current.next
		current.next = prev
		prev = current
		current = temp
	}
	list.head = prev
}

func (list *linkedList1) ShowBack() (int, error) {
	if list.head == nil {
		return 0, fmt.Errorf("Cannot Find ShowFront Value in an Empty linked list")
	}
	return list.tail.val, nil
}

func main() {

	list := linkedList1{}
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
	fmt.Println("Length = ", list.tail.val)

	//list.Show()
	list.Show()
	fmt.Println("Length = ", list.Len())

	//list.Delete(40)

	fmt.Println("Reverse ")
	//list.Reverse()

	list.Reverse()
	//fmt.Println("Length = ", list.Len())
	list.Show()

	fmt.Println("Reverse ")
	list.Reverse()
	//fmt.Println("Length = ", list.Len())
	list.Show()

	fmt.Println("Front: ")
	list.ShowFront()

	fmt.Println("Back: ")
	list.ShowBack()

	//front, _ := list.ShowFront()
	//back, _ := list.ShowBack()
	//fmt.Println("ShowFront = ", front)
	//fmt.Println("ShowBack = ", back)

}
