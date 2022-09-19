package main

import (
	"sync"
	"time"
)

var done bool
var mu sync.Mutex

func main() {
	// Often time you want code that periodically does something
	// A simple way to do that would be in a separate function in an infinite loop you do something with a sleep.
	// Here is an example of spawning a bunch of go routines/sendPRC() in parallel.
	// In lab 2 - if you were a candidate, you want to ask for votes in parallel from all the followers
	// sendRPC is a blocking operation that might take some time
	// Similarly the leader might need to send append entries to all teh followers, you wan to do in parallel not serially
	// So threads are a clean way to express this idea
	// In a for loop spawning a bunch of go routines.

	println("start...")
	go periodic1()
	println("done...")

	// Once modification of this could be
	// you might want to do something periodically until something happens
	// For example you might want to start a raft peer and periodically send heart beats, but whwn we call .kill on the raft instance
	// You might want to shut down all these go routines, you do not want these go routines running in the background

	println("start...")
	time.Sleep(1 * time.Second)
	go periodic2()

	time.Sleep(5 * time.Second)

	mu.Lock()
	done = true
	mu.Unlock()

	println("cancelled")
	time.Sleep(3 * time.Second) // observe something

	println("done...")
}

func periodic1() {
	for {
		println("tick...")
		time.Sleep(1 * time.Second)
	}
}

// you have a go routine that runs in a infinite loop, does something [tick] and we sleep for a little
// you might have a shared variable, like here we have a global variable done.
// what main does is waits for a while ans sets it to true/done
// and in this go routine which is doing some work periodically we are checking the value of done.
// If done is set to true we terminate the go routine.
// and since done is a shared variable, being read and mutated by multiple threads,
// we need to make sure we guard the use of it with a lock - mu.Lock() & mu.Unlock()
// finally when done is set and unlocked, the next lock is guaranteed to observe the done of true and will terminate the thread/return
// Now lets say, while one thread is holding thr lock, if the other thread calls the lock, that other thread will be blocked until the first one unlocks it
// ??? Only a single thread can be executing/holding the lock of the critical section
// On thing within go is if the main goroutine exits, all the other go routines are terminated
func periodic2() {
	for {
		println("tick...")
		time.Sleep(1 * time.Second)
		mu.Lock()
		if done {
			mu.Unlock()
			return
		}
		mu.Unlock()

		// or
		// we could also do something more simpler
		//for !rf.killed() {
		//	println("tick...")
		//	time.Sleep(1 * time.Second)
		//}
	}
}
