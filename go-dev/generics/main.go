package main

import "fmt"

func SumInts(m map[string]int64) int64 {
	var sum int64
	for key, val := range m {
		sum += val
		println("Key ", key, " Value ", val, " Sum ", sum)
	}
	return sum
}

func SumFloats(m map[string]float64) float64 {
	var sum float64
	for key, val := range m {
		sum += val
		println("Key ", key, " Value ", val, " Sum ", sum)
	}
	return sum
}

type MyKey interface {
	string | float64
}
type MyValue interface {
	int64 | float64
}

// SumNumbers sums the values of map m. It supports both integers and floats as map values.
func SumNumbers[Key MyKey, Val MyValue](m map[Key]Val) Val {
	var sum Val
	for key, val := range m {
		println("Key ", key, " Value ", val, " Sum ", sum)
		sum += val
	}
	return sum
}

func main() {
	mapOfInts := map[string]int64{
		"first":  34,
		"second": 12,
	}
	mapOfFloats := map[string]float64{
		"first":  34.19,
		"second": 12.01,
	}

	fmt.Printf("NonGeneric Sums Implementations")
	sumOfInts := SumInts(mapOfInts)
	sumOfFloats := SumFloats(mapOfFloats)
	fmt.Printf("NonGeneric Sums :%v and %v\n", sumOfInts, sumOfFloats)

	fmt.Printf("Generic Sums Implementations with Constraint")
	sumNumbersInts := SumNumbers(mapOfInts)
	sumNumbersFloats := SumNumbers(mapOfFloats)
	fmt.Printf("Generic Sums with Constraint: %v and %v\n", sumNumbersInts, sumNumbersFloats)
}
