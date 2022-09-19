package main

import (
	"fmt"
	//"golang.org/x/text/date"
	"time"
)

/*
So far count has been outputting to he terminal,
but what if we have to communicate back to the go routine
well for this we can accept a channel as an argument



*/

func main() {
	ch := make(chan int)
	arr := [6]int{2, 3, 5, 7, 11, 13}

	fmt.Println("*****3.1*****")
	go count31(arr, ch)
	sum := <-ch
	fmt.Println("Final Sum: ", sum)

	fmt.Println("*****3.2*****")
	go count32(arr, ch)
	for {
		sum, open := <-ch
		if !open {
			fmt.Println("breaking ")
			break
		}
		fmt.Println("Intermediate Sum: ", sum)
	}
	// or
	fmt.Println("*****3.3*****")
	go count32(arr, ch)
	for sum := range ch {
		fmt.Println("Intermediate Sum: ", sum)
	}

}

func count31(arr [6]int, c chan int) {
	var sum int
	for i := 0; i <= len(arr)-1; i++ {
		sum = sum + arr[i]
		time.Sleep(time.Millisecond * 500)
		fmt.Println("Adding element ", arr[i], " sum: ", sum)
	}
	c <- sum
	close(c)
}

func count32(arr [6]int, c chan int) {
	var sum int
	for i := 0; i <= len(arr)-1; i++ {
		sum = sum + arr[i]
		time.Sleep(time.Millisecond * 500)

		fmt.Println("Adding element ", arr[i], " sum: ", sum)
		c <- sum
	}
	close(c)
}
