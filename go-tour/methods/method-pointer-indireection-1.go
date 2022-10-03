package main

import "fmt"

type Vertex4 struct {
	X, Y float64
}

/*
Methods and pointer indirection
*** Comparing the previous two programs, you might notice that functions with a pointer argument must take a pointer:

	var v Vertex
	ScaleFunc(v, 5)  // Compile error!
	ScaleFunc(&v, 5) // OK

while methods with pointer receivers take either a value or a pointer as the receiver when they are called:

	var v Vertex
	v.Scale(5)  // OK
	p := &v
	p.Scale(10) // OK

*** For the statement v.Scale(5), even though v is a value and not a pointer,
the method with the pointer receiver is called automatically.
That is, as a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5)
since the Scale method has a pointer receiver.
*/
func main() {
	v := Vertex4{3, 4}
	v.Scale4(2)    // (&v).Scale4(2) - Go interpretation
	fmt.Println(v) // {6 8}

	p := &Vertex4{3, 4}
	p.Scale4(2)    // here we are invoking the method on the address of Vertex via p
	fmt.Println(p) // &{6 8}

	ScaleFunc4(&v, 10)
	fmt.Println(v) // {60 80}

	ScaleFunc4(p, 8)
	fmt.Println(p) // &{48 64}
}

func (v *Vertex4) Scale4(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc4(v *Vertex4, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
