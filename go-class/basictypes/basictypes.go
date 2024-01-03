package main

import (
	"fmt"
)

func main() {
	a := 2
	b := 2.01
	fmt.Printf("a: %-4v %v\n", a, a) // fmt.Printf("a: %-4d %[1]T\n", a)
	fmt.Printf("b: %-4v %v\n", b, b) // fmt.Printf("b: %-4.2f %[1]T\n", b)
	var size float32 = 1.9
	y := int(size)  // truncated to 1
	z := float32(y) // still 1.0 from 1
	print(y, "  ", z)
}
