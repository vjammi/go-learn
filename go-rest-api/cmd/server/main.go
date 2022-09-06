package main

import (
	"fmt"
)

func Run() error {
	fmt.Println("Starting up go application")
	return nil
}

func main() {
	fmt.Println("Starting the go application")
	//handleRequests()
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
