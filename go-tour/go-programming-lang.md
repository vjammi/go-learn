The Go Programming Language
https://youtu.be/rKnDgT73v8s

Go is a Concurrent, Garbage Collected, Systems language
Though designed for a systems language, it has much broader use.

Goals of the lang
Efficiencies of a statically typed compiled lang
Ease of a programming of a dynamic language
It critical that lang be type-safe and memory-safe
Good support for concurrency and communication, these 2 can help software on the network, or on multi core cpus
Efficient, latency-free garbage collection
We want it to be garbage collected. Lot of book keepoing that has in modern prog lang, specially in C/C++ has to be to do with memory managment. that can be totally automated. you can do that effeciantly and latency free. You can derive a bad address and then use it
Garbage collection has got to a point where garbage collection systems lang make a lot of sence.
High Speed compilation
Also we want the compilations to run fast.
you want carriage return times for compilations. where you want to be done.

Design Principles
https://youtu.be/rKnDgT73v8s?t=694
Avoid bookkeeping
	Go is more object-oriented lang than java or C++
	No Type Hierarchy???

Big Picture
https://youtu.be/rKnDgT73v8s?t=802
Fundamentals
Clean and concise sysmtax
lighweight type system
No implicit conversations: keep things explicit
Untyped unsized constants
Strict saperation of interface and implemetation
Runtime
Garbage collected
Goood support for stings, maps, communication channels
Concurrency
Package Model
Explicit dependencies to enable faster builds
New aproach to dependencies and dependenciy trees
If A.go depends on B.go which might depends on c.go
Inm usuall order you need to compile c, b and a
But in go, when you go to compile 1.go it does not need b to be compile and c to compiled first.
every thing that B needs from c or a needs from b is already been pulled
The compiler pulls transtitive dependency type info from the object file so now

Concurrency
Concurrent garbage collected processes - called go routines
the name has beebn changed.
Normally we have Threads, Processses, CoRoutines - Goroutines
The language takes care of managing the go routines, managing the memory they use, and the communication between them
stacks grow automatically. ytou do not have to declare how big the stack need to be
Go routines are multiplexed on teh threads
One of the need for garbage collection is that concurrent programming is very hard without garbage collection
because when you hand items between these concurrent things, between these processes, who pwns the memory, whoo's job is it to free it. there is a lot of book keeping involved if the lang does not take are of it for you. so we need garbage collection to make concurrency work well


Basics
Lexically it looks more like C
Lang requires that it be UTF-8 encoded

Declare and Initialize
/*Define and use a tppe, T. */
type T struct {a,b int}

	var t0 *T = new (T)
		is equivalent to 
	t1 := new(T) // t1 colon eqiual to new T . What colon equal means is declare and initialize. 
				 // type taken from expression. type of t1 is derived from the expression that’s used to declare and initialize it. 
			
	var f0 *Foo = new(Foo)
	f0.a = 1
	f0.b = 2
	fmt.Println(">>", f0.a, f0.b)
	fmt.Println("> ", f0)

	// is equivalent to
	f1 := new(Foo) // t1 colon equal to new T . What colon equal means is declare and initialize.
	f1.a = 1
	f1.b = 2
	fmt.Println(">>", f1.a, f1.b)
	// type taken from expression. type of t1 is derived from the expression that’s used to declare and initialize it.



// Control Structrures - No parenthesis, always braces
if len(str) > 0 {
...
}
The braces are required on the blocks and there’s no parenthesis on, for instance, the conditional expression inside an "if" statement. That's just to clean
The main difference is syntactic. The braces are required on the blocks and there’s no parenthesis on, for instance, the conditional expression inside an "if" statement. That's just to clean


Contants in go
type TZ int   // timezone

	const(
		UTC TX = 0*60*60;
		EST TZ = -5*60*60;
	)
	const (
		bit0,...
	)

	const Ln2 = 0.55972347507357105715701750170571057015470154711115798751908...
	const Log2E = 1/Ln2 // precise reciprocal


Values and types
   
    Slices
    weekend := string {"Saturday", "Sunday"}

	Maps
		We declare timezones to be a map from string to TZ (time zone)
		timezones := map[string]TZ{
			"UTC":UTC, "EST": EST, //...
		}

	func add (a,b int) int {return a+b }

	type op func (int, int) int

	type RPC struct {
		a,b int;
		op Op;
		result *int;
	}
	// declarations look backward to a C programmer.
	rpc := RPC {1,2, add, new(int)} // a and b are integers, op, and its has a result - which is a pointer to an integer	
	// we declare using a struct literal, an instance of an RPC, type name RPC open brace, list the element sof the structure, so now you got a value. You can now assign, and in this case initialize 
	// the new  value to this variable. 

Methods
let’s talk about methods, because they’re kind of different. So here’s a type called point that has X and Y. Now, those are upper case and the reason is, in Go, the complete
decision about whether an identifier is visible outside the package or not, is whether or not it starts with an uppercase Unicode letter. That’s it. So if a variable or a type or
whatever declared at the top level is uppercase, is visible to the clients of the package. If it’s lower case, it’s not. Inside the global struct, if a method or field is uppercase,	


