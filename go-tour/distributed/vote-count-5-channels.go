package main

import "time"
import "math/rand"

/*
The final topic

We're going to cover in terms of go concurrency primitives is channels so two high level channels are like a
queue like synchronization primitive but they don't behave quite like cues in the
intuitive sense like I think some people think of channels is like there's this
data structure we can sticks that stick things in and eventually someone will pull those things out but in fact channels have no queuing capacity they
have no internal storage basically channels are synchronous if you have to
go routines that are going to send and receive on a channel if someone tries to send on the channel while nobody's receiving that thread will block until
somebody's ready to receive and at that point synchronously it will exchange that data over to the receiver and the
same is true the other direction if someone tries to receive from a channel while nobody's sending that receive will block until there's another goroutine
that's about to send on the channel and that send will happen synchronously so here's a little demo program that
demonstrates this here I have a I declare channel and then I spawn a go routine that waits for a second and then
sent and then receives from a channel and then in my main girl routine I keep
track of the time then I send on the channel so I just put some dummy data into the channel and then I'm going to print out how long the send took
and if you think of channels as cues with internal storage capacity you might
think of this thing as completing very fast but that's not how channels work this send is going to block until this
receive happens and this one happened till this one second is the elapsed and so from here to here we're actually blocked in the main go routine for one whole second

alright So don't think of channels as queues think of them as this synchronous like the synchronous communication mechanism

Another example that'll make this really obvious is here we have a goroutine that creates a channel then sends on the
channel and tries receiving from it doesn't anybody know what'll happen when I try running this I think the file name might give it away
	func main(){
		c:= make(chan bool)
		c <- true // send will block until there is someone to receive
		  <- c    // but there is no one to receive
	}
Yeah exactly the send is going to block till somebody's ready to receive but there is no receiver and go actually
detects this condition if all your threads are sleeping it to text this is a deadlock condition and it'll actually crash but you can have more subtle bugs

where if you have some other thread like off doing something if I spawn this go routine that you know for loop does nothing
	func main(){
		go func(){
		  for{}
		} ()
		c := make(chan bool)
		c <- true
		  <- c
	}
and I try running this program again now go's deadlock detector won't notice that all threads are not doing any useful work
like there's one thread running it's just this is never receiving and we can tell by looking at
this program that it'll never terminate but here it just looks like it hangs so if you're not careful with channels you
can get these subtle bugs where you have double X as a result yeah yeah exactly
there's no data nobody's sending on this channel so this is gonna block here it's never gonna get to this line
yeah so channels as you pointed out can't really be used just within a single goroutine it doesn't really make
sense because in order to send or in order to receive there has to be another go routine doing the opposite action at
the same time so if there isn't you're just gonna block forever and then that chant but thread will no longer do any
useful work yeah
send waits for receives and receives wait for sends and it happens synchronously once there's both the sender and receiver present what I
talked about so far is unbuffered channels I was going to avoid talking about buffered channels because there are very few problems that they're
actually useful for solving so buffered channels can take in a capacity and then you can think of it as
	func main(){
		go func(){
		  for{}
		}()
		c := make(chan bool, 1)
		c <- true
		  <- c
	}
so here's a buffered channel with a capacity of one this program does terminate because buffered channels are like they have some internal storage
space and until that space fills up sends are non blocking because they can just put that data in the internal
storage space but once the channel does fill up then it does behave like a nun buffer channel in the sense that further
sends will block until there's a receive to make space in the channel

but I think at a high level we should avoid buffered channels because they basically don't solve any problems and another path and
other things should be thinking about is whenever you to make up arbitrary numbers like this one here to make your code work you're probably doing something wrong yeah
so I think this is a question about terminology like what exactly does deadlock mean into this count as a deadlock like yes this counts as a
deadlock like no useful progress will be made here. like this these threads are just stuck forever


## So what our channels useful for
I think channels are useful for a small set of things
like for example I think for producer consumer queues sort of situations like here I have a program
```
	func main(){
			c := make(chan int)

			for i := 0; i< 4; i++ {
				doWork(c)
			}

			for {
				v := <-c
				println(v)
			}
	}

	func doWork(c chan int){
		for{
			time.Sleep()
			c <- rand.Int()
		}
	}
```
that makes a channel and this spawns a bunch of goroutines that are going to be doing some work, like say they're computing
some result and producing some data and I have a bunch of these go routines running in parallel and I want to
collect all that data as it comes in and do something with it so this do work thing just like waits for a bit and produces a random number
and in the main goroutine I'm going to continuously receive on this channel and print it out like this is a great use of channels

## Another good use of channels is to achieve something similar to what wait groups do
So rather than use a wait group suppose I want to spawn a bunch of threads and wait till they're all done doing something
``` wait.go

	func main(){
			done := make(chan bool)
			for i := 0; i < 5; i++ {
				go func(x int) {
					sendRPC(x)
					done <-true
				}(i)
			}
			for i := 0; i < 5; i++ {
				 <-done
			}
	}
	func sendRPC(c chan int){
		println(i)
	}
```
one way to do that is to create a channel and then I spawn a bunch of threads and know how many threads I've spawned
so five goroutines created here they're going to do something and then send on this channel when they're done and then
in the main go routine I can just receive from that channel the same number of times and this has the same effect as a
wait group. question ...
Ah so the question here is could you use a buffered channel with a capacity of five because you're waiting for five receives
I think in this particular case yes that would have the equivalent effect but I think there's not really a reason to do
that and I think at a high level in your code you should avoid buffer channels and also maybe even channels unless you
think very hard about what you're doing yeah

## Wait Group
So a wait group is a yet another synchronization primitive provided by go in the sync package
and it kind of does what its name advertises like it lets you wait for a certain number of threads to be done
``` loop.go / waitgroup.go
	func main() {
		var wg sync.WaitGroup
		for i := 0; i < 5; i++ {
			wg.Add(i)
			go func(x int) {
				sendRPC(x)
				wg.Done()
			}(i)
		}
		wg.Wait()
		println("done...")
	}

	func sendRPC(i int) {
		println(i)
	}
```
the way it works is,  you call wait group dot add and that basically increments some internal counter
and then when you call wait group dot wait it waits till done has been called as many times as add was called
so this code is basically the same as the code I just showed you that was using a channel
except this is using wait group they have the exact same effect you can use either one

## <question> So the question here is about race conditions
I think like what happens if this add doesn't happen fast enough before this wait happens or something like that
Well so here notice that the pattern here is we call wait group data outside of this go routine and it's called before spawning this go routine
so this happens first this happens next and so we'll never have the situation we're done happens after this add happens for this particular routine how's this
implemented by the compiler and I will not talk about that now but talk to me after class or in office hours but I
think for the purposes class like you need to know the API for these things not the implementation all right and so
I think that's basically all I have on go concurrency primitives

## Good Advice - So what are channels good for?
So one final thought is on channels like channels are good for a specific set of things like I just showed you
The producer consumer queue or like implementing something like wait groups but I think when you try to do fancier things
with them like if you want to say like kick another go routine that may or may not be waiting for you to be like woken up
that's a kind of tricky thing to do with channels there's also a bunch of other ways to shoot yourself in the foot
with them I'm going to avoid showing you examples of bad code with channels just because it's not useful to see
but I personally avoid using channels for the most part and just use shared memory and mutexes and condition variables
and set and I personally find those much easier to reason about so feel free to use channels for when they make sense
but if anything looks especially awkward to do with channels like just use mutexes and condition variables and they're probably
a better tool

## So the question is what is the difference between this producer-consumer pattern and a thread-safe FIFO
I think they're kind of equivalent like you could do this with the thread-safe FIFO and it like that is basically what a
like buffered channel is roughly

## Buffered Channel
if you're in queueing things in and queueing things,
like if you want this line to finish and have this thread go do something else while that data sits there in a queue
rather than this go routine waiting to send it then a buffered channel might make sense
but I think at least in the labs you will not have a pattern like that
*/

func main() {
	rand.Seed(time.Now().UnixNano())

	count := 0
	finished := 0
	ch := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			ch <- requestVote5()
		}()
	}

	for count < 5 && finished < 10 {
		v := <-ch
		if v {
			count += 1
		}
		finished += 1
	}

	if count >= 5 {
		println("received 5+ votes!")

	} else {
		println("lost")
	}
}

func requestVote5() bool {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	return rand.Int()%2 == 0
}
