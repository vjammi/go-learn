package main

import "fmt"

// Type parameters
// Go functions can be written to work on multiple types using type parameters.
// The type parameters of a function appear between brackets, before the function's arguments.
//
// func Index[T comparable](s []T, x T) int
// This declaration means that s is a slice of any type T that fulfills the built-in constraint comparable. x is also a value of the same type.
//
//comparable is a useful constraint that makes it possible to use the == and != operators on values of the type. In this example, we use it to compare a value to all slice elements until a match is found. This Index function works for any type that supports comparison.

func main() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
