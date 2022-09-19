package main

import "sync"

func main() {

	var wg sync.WaitGroup

	// An example of incorrect impl, where the variable i is being used in the anonymous function from teh outer scope
	// the variable i is being mutated by outer scope and by the time the function call is executed the value has already mutated by the outer for loop
	for i := 0; i < 10; i++ {
		wg.Add(i)
		go func() {
			sendRPC1(i)
			wg.Done()
		}()
	}
	wg.Wait()

	// The correct way to pass the variable i as an argument to the function
	for i := 0; i < 10; i++ {
		wg.Add(i)
		go func(x int) {
			sendRPC1(x)
			wg.Done()
		}(i)
	}
	wg.Wait()

	println("done...")
}

func sendRPC1(i int) {
	println("rpc: ", i)
}
