package main

/*
Arrays
The type [n]T is an array of n values of type T.

The expression

var a [10]int
declares a variable a as an array of ten integers.

An array's length is part of its type, so arrays cannot be resized. This seems limiting, but don't worry; Go provides a convenient way of working with arrays.
*/

import "fmt"

func main() {
	var a [2]string // array of n values of type T/string.
	a[0] = "123"
	a[1] = "456"
	fmt.Println("a[0]:", a[0], " a[1]:", a[1])

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var evens [6]int
	evens[0] = 2
	evens[1] = 4
	evens[2] = 6
	evens[3] = 8
	evens[4] = 10
	evens[5] = 12
	fmt.Println(evens)
}
