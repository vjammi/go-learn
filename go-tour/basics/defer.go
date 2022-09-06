package main

import "fmt"

func main() {
	defer_statement()
	stacking_deferred_fnc_calls()
}

/*
Defer
A defer statement defers the execution of a function until the surrounding function returns.
The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
*/
func defer_statement() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

/*
Stacking defers
Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
To learn more about defer statements read this blog post.
*/
func stacking_deferred_fnc_calls() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
