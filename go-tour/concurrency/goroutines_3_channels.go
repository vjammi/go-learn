package main

import (
	"fmt"
	"time"
)

func main() {
	// Receive one message
	c := make(chan string, 50)

	go count3("sheep", c)

	msg := <-c
	fmt.Println(msg)

	// Receive all messages
	c2 := make(chan string, 50)

	go count3("sheep", c2)

	for {
		msg, open := <-c
		if !open {
			break
		}
		fmt.Println(msg)
	}
}

func count3(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing // 	fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}
