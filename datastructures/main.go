package main

import "fmt"

type Queue

func main() {
	fmt.Println("Queue Impl in Golang")
	queue := new(Queue)
	queue.enqueue("1")
	queue.enqueue("2")
	queue.enqueue("3")
	queue.enqueue("4")
	queue.enqueue("5")
	queue.enqueue("6")
	queue.enqueue("7")
	queue.enqueue("8")
	fmt.Print(queue)
}
