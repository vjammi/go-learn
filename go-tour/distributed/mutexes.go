package main

import (
	"sync"
	"time"
)

/**
By default, if we have concurrent access by different threads to some shared data
we want to makenew sure that the reads and writes of that data are atomic
the way we have blocks of code run atomically is by using locks

For example the rpc handlers within the raft structure - the updates need to be synchronized
with other concurrently happening updates.
Often the pattern would be grab a lock, defer unlock and go do some work inside.

At a high level what a lock or a mutex can do is guarantee mutual exclusion,
for a region of code that is called the critical section,
it ensures none of these critical sections execute concurrently, but serially.

*** Observation ***
The main go routine waits for a second for the thousand to finish and then prints the counter.
A different way to have written this code is to have the main go routine wait for
all the 1000 threads to be finished. This could be also done with a wait group.

We did not want to use 2 synchronization primitives together - wait group and mutexes.

At a high level you grab the lock, you mutate the shared data and you unlock.
Turns out that that is a useful starting point but not the complete story
*/

func main() {
	counter := 0
	var mut sync.Mutex

	for i := 0; i < 100; i++ {
		go func() {
			//mut.Lock()
			//counter = counter + 1
			//mut.Unlock()

			// The above can also be written as

			mut.Lock()
			defer mut.Unlock() // defer is scheduling this tyo run at teh end of the current function body
			counter = counter + 1
		}()
	}

	time.Sleep(1 * time.Second)
	mut.Lock()
	println(counter)
	mut.Unlock()
}
