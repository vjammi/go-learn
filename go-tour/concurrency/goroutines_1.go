package main

import (
	"fmt"
	"time"
)

func main() {
	// Synchronous function
	//count("sheep")
	//count("fish")

	// Asynchronous functions - go routine
	go count("sheep")
	go count("fish")

	//time.Sleep(time.Second * 20)
	fmt.Scanln() // blocking call - waits for user input Not very useful

}

func count(thing string) {
	for i := 0; i < 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
