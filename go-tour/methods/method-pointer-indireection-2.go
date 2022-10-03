package main

import (
	"fmt"
	"math"
)

/*
Methods and pointer indirection (2)
The equivalent thing happens in the reverse direction.

Functions that take a value argument must take a value of that specific type:

	var v Vertex
	fmt.Println(AbsFunc(v))  // OK
	fmt.Println(AbsFunc(&v)) // Compile error!

while methods with value receivers take either a value or a pointer as the receiver when they are called:

	var v Vertex
	fmt.Println(v.Abs()) // OK
	p := &v
	fmt.Println(p.Abs()) // OK

In this case, the method call p.Abs() is interpreted as (*p).Abs().
*/

type Vertex5 struct {
	X, Y float64
}

func main() {

	v := Vertex5{3, 4}
	fmt.Println(v.Abs()) // (*v).Abs()
	//fmt.Println(AbsFunc(&v)) // Compile error!  Cannot use &v (value of type *Vertex5) as type Vertex5 in argument to AbsFunc
	fmt.Println(AbsFunc(v))

	p := &Vertex5{3, 4}
	fmt.Println(p.Abs())     //(*p).Abs()
	fmt.Println(AbsFunc(*p)) // need to send the pointer
}

func (v Vertex5) Abs() float64 {
	sqrt := math.Sqrt(v.X*v.X + v.Y*v.Y)
	print(sqrt)
	return sqrt
}

func AbsFunc(v Vertex5) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
