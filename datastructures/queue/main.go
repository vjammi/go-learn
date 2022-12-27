package main

import "fmt"

func main() {
	queue := New()
	fmt.Println("Queue Impl in Golang")
	queue.Enqueue("1")
	queue.Enqueue("2")
	queue.Enqueue("3")
	queue.Enqueue("4")
	queue.Enqueue("5")
	queue.Enqueue("6")
	queue.Enqueue("7")
	queue.Enqueue("8")
	fmt.Println(queue.FrontQueue())
	fmt.Println(queue.BackQueue())
}
