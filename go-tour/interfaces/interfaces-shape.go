package main

import (
	"fmt"
	"math"
)

// interface
type shape interface {
	area() float64
	name() string
}

// types having a common behavior
type circle struct {
	radius    float64
	shapeName string
}

type square struct {
	length    float64
	breadth   float64
	shapeName string
}

type rectangle struct {
	length    float64
	breadth   float64
	shapeName string
}

// Shapes implementing the same behavior area()
func (sq square) area() float64 {
	return sq.length * sq.breadth
}

func (rect rectangle) area() float64 {
	return rect.length * rect.breadth
}

func (crc circle) area() float64 {
	return math.Pi * crc.radius * crc.radius
}

// implementing the behavior name()
func (sq square) name() string {
	return sq.shapeName
}

func (rect rectangle) name() string {
	return rect.shapeName
}

func (crc circle) name() string {
	return crc.shapeName
}

// putting it all together
func main() {
	sq := square{5, 5, "square"}
	rect := rectangle{5, 10, "rectangle"}
	circ := circle{5, "circle"}

	shapes := []shape{sq, rect, circ}
	for _, shape := range shapes {
		fmt.Println("Area of shape [", shape.name(), "] is ", shape.area())
	}
}
