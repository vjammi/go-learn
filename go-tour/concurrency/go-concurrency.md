
### Don't communicate by sharing memory; share memory by communicating
The gist is instead of protecting one fixed memory address with a lock or other concurrent primitives, you can architect the program in a way that only one stream of execution is allowed to access this memory by design.
The easy way to achieve this is to share the reference to the memory by channels. Once you send the reference over the channel you forget it. In this way only the routine that will consume that channel will have access to it.

https://go.dev/ref/mem
https://go.dev/doc/codewalk/sharemem/
https://go.dev/blog/codelab-share

### Concurrency in Go
	https://youtu.be/LvgVSSpwND8
	https://codeburst.io/why-goroutines-are-not-lightweight-threads-7c460c1f155f

### Additional References
    Concurrency is not parallelism by Rob Pike
    Analysis of Go runtime Scheduler
    Five things that make Go fast by Dave Cheney
    Discussion in golang-nuts mailing list