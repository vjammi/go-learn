package main

import "fmt"
import "rsc.io/quote"

func main() {
	fmt.Println("hello", quote.Go())
	fmt.Println("hello", quote.Hello())
	fmt.Println("hello", quote.Opt())
	fmt.Println("hello", quote.Glass())
}
