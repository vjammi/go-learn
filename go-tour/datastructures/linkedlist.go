package main

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

func (list *linkedList) addNode(val int) {
	//node := &listNode{val: val}
	//		or
	//node := new(listNode)
	//node.val = val
	//		which is same as
	//var node *listNode = new(listNode)
	node := new(listNode)
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

func (list *linkedList) Show() {
	currentNode := &list.head
	for currentNode != nil {
		fmt.Printf("%v -> ", list.head.val)
		list.head = list.head.next
	}
	fmt.Println()
}

// Note: When we have a mix of (list *linkedList)  &&  (list linkedList)
// Go Runtime complains Struct linkedList has methods on both value and pointer receivers. Such usage is not recommended by the Go Documentation.
func (list *linkedList) show() {
	current := list.head
	for current != nil {
		fmt.Println(current.val)
		current = current.next
	}
}

func (list *linkedList) ShowFront() (int, error) {
	if list.head == nil {
		return 0, fmt.Errorf("Cannot Find ShowFront Value in an Empty linked list")
	}
	return list.head.val, nil
}

func (list *linkedList) ShowBack() (int, error) {
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

func main() {

	list := linkedList{}
	//node1 := &listNode{val: 20}
	//node2 := &listNode{val: 30}
	//node3 := &listNode{val: 40}
	//node4 := &listNode{val: 50}
	//node5 := &listNode{val: 70}
	//
	//list.Add(node1)
	//list.Add(node2)
	//list.Add(node3)
	//list.Add(node4)
	//list.Add(node5)

	list.addNode(10)
	fmt.Println("Length = ", list.tail.val)
	list.addNode(20)
	fmt.Println("Length = ", list.tail.val)
	list.addNode(30)
	fmt.Println("Length = ", list.tail.val)
	list.addNode(40)
	fmt.Println("Length = ", list.tail.val)
	list.addNode(50)
	fmt.Println("Length = ", list.tail.val)

	//list.Show()
	list.show()
	fmt.Println("Length = ", list.Len())

	//list.Delete(40)
	//list.Reverse()
	//fmt.Println("Length = ", list.Len())
	//list.Show()
	//
	//front, _ := list.ShowFront()
	//back, _ := list.ShowBack()
	//fmt.Println("ShowFront = ", front)
	//fmt.Println("ShowBack = ", back)

}
