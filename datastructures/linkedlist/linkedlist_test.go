package linkedlist

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	list := New()
	list.AddNode(10)
	list.AddNode(20)
	list.AddNode(30)
	list.AddNode(40)
	list.AddNode(50)
	fmt.Println("Tail = ", list.tail.val)

	//list.Show()
	list.Show()
	fmt.Println("Length = ", list.Len())

	fmt.Println("Reverse ")

	list.Reverse()
	fmt.Println("Length = ", list.Len())
	list.Show()

	fmt.Println("Reverse ")
	list.Reverse()
	list.Show()

	fmt.Println("Delete ")
	list.Delete(30)
	fmt.Println("Length = ", list.Len())
	list.Show()

	fmt.Println("Front: ")
	list.ShowFront()

	fmt.Println("Back: ")
	list.ShowBack()
}
