package main

import (
	"fmt"
)

type Person struct {
	Name     string
	Age      int
	personID int
}

func main() {
	basicVariableDeclaration()
	multipleVariableAssignment()
}

func basicVariableDeclaration() {
	// Go is a statically typed language. We know the type of a variable at compilation time.

	// Basic variable declaration. Declares a variable of type specified on the right.
	// The variable is initialized to the zero value of the respective type.
	var x int
	var s string
	var p Person // Assuming type Person struct {}
	person := Person{Name: "Green", Age: 25, personID: 12345}

	fmt.Println(p)
	fmt.Println(person, person.Name, person.Age)

	// Assignment of a value to a variable
	x = 3

	// Short declaration using := infers the type
	y := 4

	u := int64(100)    // declare variable of type int64 and init with 100
	var u2 int64 = 100 // declare variable of type int64 and init with 100

	fmt.Println(x, s, p, x, y, u, u2)
}

func multipleVariableAssignment() {
	xx, yy := multipleReturn()  // x = 1, y = 2
	ww, zz := multipleReturn2() // w = 3, z = 4
	fmt.Printf("xx: %d, yy: %d ww: %d, zz: %d \n", xx, yy, ww, zz)
}

func multipleReturn() (int, int) {
	return 1, 2
}

func multipleReturn2() (a int, b int) {
	a = 3
	b = 4
	return
}
