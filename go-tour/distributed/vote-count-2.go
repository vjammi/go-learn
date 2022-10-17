package main

import "sync"
import "time"
import "math/rand"

/*

So here I declare mutex that's accessible by everybody and then in the go routines I'm launching in parallel to request votes
I'm going to and this pattern here is pretty important I'm going to first request a vote while I'm not holding the lock
and then after wear that I'm going to grab the lock and then update these shared variables
and then outside I have the same patterns as before except I makenew sure to lock and unlock between reading
these shared variables so in an infinite loop I grab the lock and check to see if the results of the election have been
determined by this point and if not I'm going to keep running in this infinite loop otherwise I'll unlock and then do
what I need to do outside of here
and so if I run this example whoops it seems to work and this is actually like a correct implementation it does the right thing

*/

func main() {
	rand.Seed(time.Now().UnixNano())

	count := 0
	finished := 0
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		go func() {
			vote := requestVote2()
			mu.Lock()
			defer mu.Unlock()
			if vote {
				count++
			}
			finished++
		}()
	}

	for {
		mu.Lock()

		if count >= 5 || finished == 10 {
			break
		}
		mu.Unlock()
	}

	if count >= 5 {
		println("received 5+ votes!")
	} else {
		println("lost")
	}
	mu.Unlock()
}

func requestVote2() bool {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	return rand.Int()%2 == 0
}
