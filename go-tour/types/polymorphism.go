package main

import "fmt"

type Vehicle struct {
	wheels int
}
type Car struct {
	wheels int
}
type MotorCycle struct{}

// https://youtu.be/1ZjvhGfpwJ8?t=654
func (v *Vehicle) ComputeSpeed() int {
	return v.wheels * 5
}

type Speeder interface {
	GetSpeed() int
}

func (c *Car) GetSpeed() int {
	return c.wheels * 3
}

func (m *MotorCycle) GetSpeed() int {
	//return m.Base.ComputeSpeed()
	return 0
}

func printSpeed(s Speeder) {
	fmt.Println(s.GetSpeed())
}

func main() {
	car := new(Car)
	motorCycle := new(MotorCycle)

	fmt.Println(car.GetSpeed())
	fmt.Println(motorCycle.GetSpeed())

	printSpeed(car)
	printSpeed(motorCycle)
}
