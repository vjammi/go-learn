package main

import (
	"fmt"
)

type Stack struct {
	len  int
	head *Stack
}

// The new built-in function allocates memory. The first argument is a type,
// not a value, and the value returned is a pointer to a newly
// allocated zero value of that type.
func main() {

	// Option 1
	var stackObj1 Stack
	fmt.Println(stackObj1)

	// The new built-in function allocates memory.
	// The first argument is a type, not a value, and
	// the value returned is a pointer to a newly
	// allocated zero value of that type.

	// Option 2
	stackObj2 := new(Stack)
	fmt.Println(stackObj2)

	// Option 3
	stackObj3 := Stack{}
	fmt.Println(stackObj3)

}
