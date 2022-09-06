// Example code for Chapter 2.3 from "Build Web Application with Golang"
// Purpose: Shows passing a variable by value and reference
// Reference: https://github.com/zwfang/build-web-application-with-golang/blob/master/en/code/src/apps/ch.2.3/pass_by_value_and_pointer/main.go
package main

import "fmt"

func add_by_value(a int) int {
	a = a + 1
	return a
}

func add_by_reference(a *int) int {
	*a = *a + 1
	return *a
}

func show_add_by_value() {
	x := 3
	fmt.Println("[Before] x = ", x)
	fmt.Println("add_by_value(x) =", add_by_value(x))
	fmt.Println("[After] x = ", x)
}

func show_add_by_reference() {
	x := 3
	fmt.Println("[Before] - x = ", x)
	// &x pass memory address of x
	fmt.Println("add_by_reference(&x) =", add_by_reference(&x))
	fmt.Println("[After] x = ", x)
}

func main() {
	fmt.Println()
	show_add_by_value()
	fmt.Println()
	show_add_by_reference()
}
