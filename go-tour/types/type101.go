package main

import "fmt"

/*Define and use a type, T. */
type Foo struct {
	a, b int
}

func main() {
	var f0 *Foo = new(Foo)
	f0.a = 1
	f0.b = 2
	fmt.Println(">>", f0.a, f0.b)
	fmt.Println("> ", f0)

	// is equivalent to
	f1 := new(Foo) // t1 colon equal to new T . What colon equal means is declare and initialize.
	f1.a = 1
	f1.b = 2
	fmt.Println(">>", f1.a, f1.b)
	// type taken from expression. type of t1 is derived from the expression that’s used to declare and initialize it.
}
