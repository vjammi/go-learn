package main

import "time"
import "math/rand"

/*
So the next synchronization primitive we're going to talk about it something called conditional variables.
We're going to do that in the context counting votes. So remember in lab 2a you have this pattern
where whenever raft pear becomes a candidate it wants to send out vote requests all of its followers
and eventually the followers come back to the candidate and say yes or no, whether or not the candidate got the vote
and one way we could write this code is have the candidate - in serial ask peer number 1,  number 2,  number 3 and so on
but that's bad right because we want the candidate ask all the peers in parallel so it can quickly win the election when possible

and then there's some other complexities there like when we ask all the peers in parallel we don't want to wait so we get a response
from all of them before making up our mind right because if a candidate gets a majority of votes like it doesn't need to wait till it hears back from everybody
else so this code is kind of complicated in some ways

and so here's a kind of stubbed out version of what that vote counting code might look like with a little bit of
infrastructure to make it actually run and so here have this main goroutine that sets count
which is like the number of yes votes I got to zero and finish to zero finished as the number of responses I've gotten
in total and the idea is I want to send out vote requests in parallel and keep track of how many yeses I've got and how
many responses I've gotten in general and then once I know whether I've won the election or whether I know that I've
lost the election then I can determine that and move on and like the real raft code you actually do whatever you need
to do don't step up to a leader or to step down to a follower after you have the result from this and so looking at
this code here I'm going to in parallel spun say I have ten peers in parallel spawn ten goroutines
here I pass in this closure here and I'm gonna do is request a vote and then if I get the vote I'm going to increment the
count by one and then I'm also going to increment this finished by one so like this is a number of yeses this is total
number of responses I've gotten and then outside here in the main go routine what I'm doing is keeping track of this
condition I'm waiting for this condition to become true that either I have enough yes votes that I've won the election or
I've heard back from enough peers and I know that I've lost and so I'm just going to in a in a loop check to see and
wait until count is greater than or equal to five or wait until finished is equal to ten and then after that's the
case I can either determine that I've lost drive one so does anybody see any problems with this code given what we
just talked about about mutexes

Yeah exactly - count and finished aren't protected by mutexes.
So one thing we certainly need to fix here is that whenever we have shared variables we need to protect access with mutexes
and so that's not too bad to fix here

contd. in vote-count-2.go
*/
func main() {
	rand.Seed(time.Now().UnixNano())

	count := 0
	finished := 0

	for i := 0; i < 10; i++ {
		go func() {
			vote := requestVote1()
			if vote {
				count++
			}
			finished++
		}()
	}

	for count < 5 && finished != 10 {
		// wait
	}
	if count >= 5 {
		println("received 5+ votes!")
	} else {
		println("lost")
	}
}

func requestVote1() bool {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	return rand.Int()%2 == 0
}
