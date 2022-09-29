package main

import (
	"fmt"
)

type listNode struct {
	val  int
	next *listNode
}

type linkedList struct {
	head   *listNode
	tail   *listNode
	length int
}

func newLinkedListNode() linkedList {
	list := linkedList{}
	return list
}

func (list *linkedList) addNode(val int) {
	// node := &listNode{val: val} 				// list.head = &node
	//	 or
	// node := new(listNode)       				// list.head = node
	// node.val = val
	//   is same as
	// var node *listNode = new(listNode)		// list.head = node

	node := new(listNode) // var node *listNode = new(listNode)  // list.head = node
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

func (list *linkedList) len() int {
	return list.length
}

// Note: When we have a mix of (list *linkedList)  &&  (list linkedList)
// Go Runtime complains Struct linkedList has methods on both value and pointer receivers. Such usage is not recommended by the Go Documentation.
func (list *linkedList) show() {
	// current := list.head		//	current = current.next
	// or
	var current = list.head //	current = current.next

	for current != nil {
		fmt.Print(current.val, ">")
		current = current.next
	}
}

func (list *linkedList) showFront() error {
	if list.head == nil {
		return fmt.Errorf("linkedList is empty...")
	}
	fmt.Println(list.head.val)
	return nil
}

func (list *linkedList) showBack() (int, error) {
	if list.tail == nil {
		return -1, fmt.Errorf("linkedList is empty...")
	}
	fmt.Println(list.tail.val)
	return list.tail.val, nil
}

func (list *linkedList) reverse() { // list:     *main.linkedList
	var current = list.head // current:  *main.listNode
	var previous *listNode  // previous: nil

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

func (list *linkedList) delete(val int) {
	var current = list.head
	var previous *listNode

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

func main() {
	list := newLinkedListNode()

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
