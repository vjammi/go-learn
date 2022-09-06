/*
Exercise: Loops and Functions
As a way to play with functions and loops, let's implement a square root function: given a number x, we want to find the number z for which z² is most nearly x.
Computers typically compute the square root of x using a loop. Starting with some guess z, we can adjust z based on how close z² is to x, producing a better guess:

	z -= (z*z - x) / (2*z)

Repeating this adjustment makes the guess better and better until we reach an answer that is as close to the actual square root as can be.
Implement this in the func Sqrt1 provided. A decent starting guess for z is 1, no matter what the input. To begin with, repeat the calculation 10 times and print each z along the way. See how close you get to the answer for various values of x (1, 2, 3, ...) and how quickly the guess improves.
Hint: To declare and initialize a floating point value, give it floating point syntax or use a conversion:

	z := 1.0
	z := float64(1)

Next, change the loop condition to stop once the value has stopped changing (or only changes by a very small amount). See if that's more or fewer than 10 iterations. Try other initial guesses for z, like x, or x/2. How close are your function's results to the math.Sqrt in the standard library?
(Note: If you are interested in the details of the algorithm, the z² − x above is how far away z² is from where it needs to be (x), and the division by 2z is the derivative of z², to scale how much we adjust z by how quickly z² is changing. This general approach is called Newton's method. It works well for many functions but especially well for square root.)

Solution
As a simple way to play with functions and loops, implement the square root function using Newton's method.
In this case, Newton's method is to approximate Sqrt1(x) by picking a starting point z and then repeating:

	z - (z*z - x) / (2 * z)

To begin with, just repeat that calculation 10 times and see how close you get to the answer for various values (1, 2, 3, ...).
Next, change the loop condition to stop once the value has stopped changing (or only changes by a very small delta).
See if that's more or fewer iterations. How close are you to the math.Sqrt?
Hint: to declare and initialize a floating point value, give it floating point syntax or use a conversion:

	z := float64(1)
	z := 1.0
*/
package main

import (
	"fmt"
	"math"
)

const CHANGE = 0.0000001
const INITIAL_GUESS = 1.0

func Sqrt(x float64) float64 {

	current_z := INITIAL_GUESS

	for i := 0; i < 10; i++ {
		adjusted_z := current_z - (((current_z * current_z) - x) / (2 * current_z))

		if math.Abs(adjusted_z-current_z) < CHANGE {
			break
		}
		current_z = adjusted_z
		fmt.Println("Adjusted/Improved value of Z (Guess) ", adjusted_z)
	}

	return current_z
}

func main() {
	fmt.Println(Sqrt(225))
	fmt.Println(math.Sqrt(225))
}
