package main

import (
	"fmt"
)

type event struct{}
type instrument struct{}

type artist struct {
	name         string
	performances []event
	instruments  []instrument
}

func main() {

	//event1 := [...]string{"A1", "B1", "C1"}
	//event2 := [...]string{"A2", "B2", "C2"}

	// Two ways to initializing an object

	// *** Weather the variable is on the heap or on the stack is upto the compiler - escape analysis
	// https://youtu.be/3CR4UNMK_Is?t=559
	//
	// Not Accurate
	// [Here we are creating a local variable of struct and requesting its address
	// Now, we know that any local variable gets created on the stack and that variable is lost when the control goes back to the calling function
	// However in golang, when we request the address of the local variable, golang actually creates the new instance on the heap instead of the stack, which is equivalent to new
	// This sometime might be convenient for initializing the values]
	artist1 := &artist{}
	fmt.Println(artist1)

	// Instructing the go to create a new object on the heap and give a reference to the object

	// var x type
	var artist2 *artist = new(artist)
	fmt.Println(artist2)
	// the above is identical to the below
	// new type
	artist3 := new(artist)
	fmt.Println(artist3)
}
