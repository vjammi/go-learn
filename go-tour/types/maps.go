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

We declare timezones to be a map from string to TZ (time zone)
	timezones := map[string]TZ{
		"UTC":UTC, "EST": EST, //...
	}
*/

var map1 map[string]Vertexx
var map2 map[int]Vertexx
var m3 map[int]string

func basicOperations() {
	map1 = make(map[string]Vertexx)
	map1["Bell Labs"] = Vertexx{40.68433, -74.39967}
	fmt.Println(map1["Bell Labs"])

	map2 = make(map[int]Vertexx)
	map2[1000] = Vertexx{40.68433, -74.39967}
	fmt.Println(map2[1000])

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
