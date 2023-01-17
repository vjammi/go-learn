package main

import (
	"fmt"
	"github.com/vjammi/go-learn/datastructures/linkedlist"
)

func main() {
	list := linkedlist.New()
	list.AddNode(1)
	list.AddNode(2)
	list.AddNode(3)
	list.Show()
	fmt.Println("")
	fmt.Println(list.Len())
}
