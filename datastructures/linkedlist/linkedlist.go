package main

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

// list := LinkedList{}
// return list			// Will return a parameter type of LinkedList
// or
// list := newLinkedList(LinkedList)
// return list 			// Will return a parameter type of *LinkedList
func New() *LinkedList {
	//list := LinkedList{}
	// return list		// Will return a parameter type of LinkedList
	// or
	list := new(LinkedList)
	return list // Will return a parameter type of *LinkedList
}

// node := ListNode{}
// node.val = val
//
//	or
//
// node := newLinkedList(ListNode)       				// list.head = node
// node.val = val
//
//	which is similar to
//
// var node *ListNode = newLinkedList(ListNode)			// list.head = node
// node.val = val
//
//	or
//
// node := &ListNode{val: val} 							// list.head = &node
func (list *LinkedList) addNode(val int) {
	// var node *ListNode = newLinkedList(ListNode)
	// list.head = node
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

func (list *LinkedList) len() int {
	return list.length
}

// Note: When we have a mix of (list *LinkedList)  &&  (list LinkedList)
// Go Runtime complains Struct LinkedList has methods on both value and pointer receivers. Such usage is not recommended by the Go Documentation.
func (list *LinkedList) show() {
	// current := list.head		//	current = current.next
	// or
	var current = list.head //	current = current.next

	for current != nil {
		fmt.Print(current.val, ">")
		current = current.next
	}
}

func (list *LinkedList) showFront() error {
	if list.head == nil {
		return fmt.Errorf("LinkedList is empty...")
	}
	fmt.Println(list.head.val)
	return nil
}

func (list *LinkedList) showBack() (int, error) {
	if list.tail == nil {
		return -1, fmt.Errorf("LinkedList is empty...")
	}
	fmt.Println(list.tail.val)
	return list.tail.val, nil
}

func (list *LinkedList) reverse() { // list:     *main.LinkedList
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

func (list *LinkedList) delete(val int) {
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

//func (node *ListNode) next() *ListNode {
//	if p := node.next; node.next != nil && p := node.next() {
//		return p
//	}
//	return nil
//}

//// Next returns the next list element or nil.
//func (e *Element) Next() *Element {
//	if p := e.next; e.list != nil && p != &e.list.root {
//		return p
//	}
//	return nil
//}

func main() {
	list := New()

	list.addNode(10)
	list.addNode(20)
	list.addNode(30)
	list.addNode(40)
	list.addNode(50)
	fmt.Println("Tail = ", list.tail.val)

	//list.Show()
	list.show()
	fmt.Println("Length = ", list.len())

	fmt.Println("Reverse ")

	list.reverse()
	fmt.Println("Length = ", list.len())
	list.show()

	fmt.Println("Reverse ")
	list.reverse()
	list.show()

	fmt.Println("Delete ")
	list.delete(30)
	fmt.Println("Length = ", list.len())
	list.show()

	fmt.Println("Front: ")
	list.showFront()

	fmt.Println("Back: ")
	list.showBack()
}
