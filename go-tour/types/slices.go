package main

import "fmt"

/*
Slices
	An array has a fixed size. A slice, on the other hand, is a dynamically-sized,
	flexible view into the elements of an array. In practice, slices are much more common than arrays.

	The type []T is a slice with elements of type T.
	A slice is formed by specifying two indices, a low and high bound, separated by a colon:
		a[low : high]
	This selects a half-open range which includes the first element, but excludes the last one.
	The following expression creates a slice which includes elements 1 through 3 of a:
		a[1:4]

Slices are like references to arrays
	A slice does not store any data, it just describes a section of an underlying array.
	Changing the elements of a slice modifies the corresponding elements of its underlying array.
	Other slices that share the same underlying array will see those changes.

*/

func main() {
	slices()
	slices2()
}

func slices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}

func slices2() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
