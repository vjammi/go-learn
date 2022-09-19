package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup // we create a WaitGroup, its has a counter
	wg.Add(1)             // we can then increment by 1 to say that we have 1 go routine to wait for

	go func() {
		count2("sheep", &wg)
		wg.Done() // We can then decrement the counter after the go routine finishes
	}()

	wg.Wait() // This will wait until count is 0. If any goroutine has not finished it wil wait
}

// Notice that we could also pass a pointer to the Waitgroup as an argument to the count function,
// but its really not the responsibility of this count function to decrement the counter,
// but the responsibility of the enclosing calling function,
// but instead we are going to wrap the call to this count function in an anonymous function
// wg.Done() - Since we have created the anonymous function inline we have access to the wg variable from the enclosing function scape
// wg.Wait() // This will wait until count is 0. If any goroutine has not finished it wil wait
func count2(thing string, wg *sync.WaitGroup) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
	// wg.Done() // We can then decrement the counter, when the go routine finishes
}
