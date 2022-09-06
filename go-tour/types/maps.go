package main

import "fmt"

type Vertexx struct {
	Lat, Long float64
}

func main() {
	basicOperations()
	mapLiterals()
}

/*
Maps
A map maps keys to values.
The zero value of a map is nil. A nil map has no keys, nor can keys be added.
The make function returns a map of the given type, initialized and ready for use.
*/
var m1 map[string]Vertexx
var m2 map[int]Vertexx
var m3 map[int]string

func basicOperations() {
	m1 = make(map[string]Vertexx)
	m1["Bell Labs"] = Vertexx{40.68433, -74.39967}
	fmt.Println(m1["Bell Labs"])

	m2 = make(map[int]Vertexx)
	m2[1000] = Vertexx{40.68433, -74.39967}
	fmt.Println(m2[1000])

	m3 = make(map[int]string)
	m3[1000] = "One Thousand"
	fmt.Println(m3[1000])
}

/*
Map literals
Map literals are like struct literals, but the keys are required.
*/
var map_literals = map[string]Vertexx{
	"Bell Labs": Vertexx{40.68433, -74.39967},
	"Google":    Vertexx{37.42202, -122.08408},
}

func mapLiterals() {
	fmt.Println(map_literals)
}
