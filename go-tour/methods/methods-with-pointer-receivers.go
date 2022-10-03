package main

import (
	"fmt"
	"math"
)

type Vertex6 struct {
	X, Y float64
}

func main() {
	//v := Vertex6{3, 4}
	v := &Vertex6{3, 4}

	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs6())
	v.Scale6(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs6())
}

func (v *Vertex6) Abs6() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex6) Scale6(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
