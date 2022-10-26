package main

import "fmt"

func main() {
	// a is a slice, which means that it has a reference to the underlying data elements
	a := []int{3, 4, 5}

	// b is initialized to the same slice, which means initially it contains the same reference
	b := a

	// Using mutate on &b can change it to another slice (i.e. referencing
	// different data elements) but cannot change which data a references
	mutate(&b)
	fmt.Printf("a: %#v\n", a)
	fmt.Printf("b: %#v\n", b)

	// c is initialized to be equal to a as well
	c := a
	// But modifying c's underlying data (i.e., the data it references, rather
	// than the reference itself) changes a, because c and a refer to the same data.
	setToOnes(c)
	fmt.Printf("a: %#v\n", a)
	fmt.Printf("c: %#v\n", c)
}

func mutate(a *[]int) {
	*a = []int{1, 2, 3}
}

func setToOnes(a []int) {
	for i := range a {
		a[i] = 1
	}
}
