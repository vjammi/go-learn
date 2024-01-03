package main

import (
	"fmt"
	"testing"
)

func TestGenerics(t *testing.T) {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts1(ints),
		SumFloats1(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats1[string, int64](ints),
		SumIntsOrFloats1[string, float64](floats))

	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
		SumIntsOrFloats1(ints),
		SumIntsOrFloats1(floats))

	sumOfInts := SumNumbers1(ints)
	sumOfFloats := SumNumbers1(floats)
	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		sumOfInts,
		sumOfFloats)

	wantSumOfInts := int64(46)
	wantSumOfFloats := float64(62.97)

	if wantSumOfInts != sumOfInts {
		t.Fatalf("wantSumOfInts = %q", wantSumOfInts)
	}

	if wantSumOfFloats != sumOfFloats {
		t.Fatalf("wantSumOfFloats = %v", wantSumOfFloats)
	}
}
