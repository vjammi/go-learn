/*
Switch
A switch statement is a shorter way to write a sequence of if - else statements.
It runs the first case whose value is equal to the condition expression.

Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. In effect,
the break statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's
switch cases need not be constants, and the values involved need not be integers.
*/
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	writing_switch()
	switch_eval_order()
	switch_with_no_condition()
}

func writing_switch() {
	fmt.Print("Go is running on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	case "windows":
		fmt.Println("Windows.")
	default:
		// freebsd, openbsd, plan9...
		fmt.Printf("%s.\n", os)
	}
}

/*
Switch evaluation order
Switch cases evaluate cases from top to bottom, stopping when a case succeeds.

(For example,

switch i {
case 0:
case f():
}
does not call f if i==0.)
*/
func switch_eval_order() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

/*
Switch with no condition
Switch without a condition is the same as switch true.
This construct can be a clean way to write long if-then-else chains.
*/
func switch_with_no_condition() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
