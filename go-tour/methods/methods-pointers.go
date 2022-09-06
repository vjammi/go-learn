package main

import (
	"fmt"
	"math"
)

type Vertex2 struct {
	X, Y float64
}

func main() {
	// Can declare methods with pointer receivers
	v1 := Vertex2{3, 4}
	v1.Scale(10)
	abs := v1.Abs()
	fmt.Println(abs)

	// see the Abs and Scale methods rewritten as functions.
	v2 := Vertex2{3, 4}
	Scale2(&v2, 10)
	fmt.Println(Abs2(v2))
}

/*
Pointer receivers
You can declare methods with pointer receivers.
This means the receiver type has the literal syntax *T for some type T. (Also, T cannot itself be a pointer such as *int.)
For example, the Scale method here is defined on *Vertex.
Methods with pointer receivers can modify the value to which the receiver points (as Scale does here). Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
Try removing the * from the declaration of the Scale function on line 16 and observe how the program's behavior changes.
With a value receiver, the Scale method operates on a copy of the original Vertex value. (This is the same behavior as for any other function argument.) The Scale method must have a pointer receiver to change the Vertex value declared in the main function.
*/
func (v Vertex2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex2) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

/*
Pointers and functions
Here we see the Abs and Scale methods rewritten as functions.
Again, try removing the * from line 16. Can you see why the behavior changes? What else did you need to change for the example to compile?
(If you're not sure, continue to the next page.)
*/
func Abs2(v Vertex2) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale2(v *Vertex2, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

/*
	Methods and pointer indirection
	Comparing the previous two programs, you might notice that functions with a pointer argument must take a pointer:

	var v Vertex
	ScaleFunc(v, 5)  // Compile error!
	ScaleFunc(&v, 5) // OK
	while methods with pointer receivers take either a value or a pointer as the receiver when they are called:

	var v Vertex
	v.Scale(5)  // OK
	p := &v
	p.Scale(10) // OK
	For the statement v.Scale(5), even though v is a value and not a pointer, the method with the pointer receiver is called automatically. That is, as a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver.
*/
