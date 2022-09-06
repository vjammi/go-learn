package main

import (
	"fmt"
	"math"
)

// If and else
// Variables declared inside an if short statement are also available inside any of the else blocks.
// (Both calls to pow return their results before the call to fmt.Println in main begins.)
func power(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("v >=lim ", x, n)
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(power(3, 2, 10), power(3, 3, 20))
}
