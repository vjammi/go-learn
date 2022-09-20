package main

import "sync"
import "time"
import "math/rand"

/*
	// Abstract version of the conditional variable pattern - wait and broadcast

	// Do something that might affect the condition - changes the shared data
	mu.Lock()
	cond.Broadcast()
	mu.Unlock()
	-----
 	// On the other hand we are waiting for the condition to becomes true
	mu.Lock()
	// while condition is false we wait
	while conditions == false {
		cond.Wait()
	}
	// We pass the while loop when the condition is true, while we are still holding the lock
	mu.Unlock() // finally we call unlock

	Lost wakeup problem in operating systems
    Avoid funny race conditions that might happen between wait and broadcast

    // Broadcast vs Signals
	- wait
	- signal vs broadcast
		signal wakes up only 1 waiter. however broadcast wakes up everybody who is waiting
        and they will all try to grab the lock and check the condition
		and only one of them will get the lock and will proceed
		For the purpose of this class always use broadcast and not use signal.
		We think of signal for efficiency, here we are not looking for tha level of efficiency.

*/

func main() {
	rand.Seed(time.Now().UnixNano())

	count := 0
	finished := 0
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	for i := 0; i < 10; i++ {
		go func() {
			vote := requestVote4()
			mu.Lock()
			defer mu.Unlock()
			if vote {
				count++
			}
			finished++
			cond.Broadcast()
		}()
	}

	mu.Lock()
	for count < 5 && finished != 10 {
		cond.Wait()
	}
	if count >= 5 {
		println("received 5+ votes!")
	} else {
		println("lost")
	}
	mu.Unlock()
}

func requestVote4() bool {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	return rand.Int()%2 == 0
}
