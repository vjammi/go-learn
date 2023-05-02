package main

import "fmt"

func main() {
	butWhyPointers()
	pointers()
}

// Reference: https://youtu.be/4B2rwYvuiBo
func butWhyPointers() {
	name := "Sam"

	update1(&name)
	fmt.Println("After update 1 ", name)

	address_of_name := &name
	fmt.Println(address_of_name)

	*address_of_name = "Sam Green"
	fmt.Println("After update 1 ", name)

	update2(address_of_name)
	fmt.Println("After update 2 ", name)

}

func update1(name *string) {
	*name = "Sam Blue"
}

func update2(name *string) {
	*name = "Sam Yellow"
}

func pointers() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
