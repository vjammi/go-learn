package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	queue := New()
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	queue.Enqueue(40)
	queue.Enqueue(50)
	fmt.Println(queue.FrontQueue())
	fmt.Println(queue.BackQueue())
}
