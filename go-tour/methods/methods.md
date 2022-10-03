## Methods
https://go.dev/tour/methods/1

#### Methods
Go does not have classes. However, you can define methods on types.
A method is a function with a special receiver argument.
The receiver appears in its own argument list between the func keyword and the method name.
In this example, the Abs method has a receiver of type Vertex named v.

```
import (
	"fmt"
	"math"
)

type Vertex struct { // Similar to a custom linkedList type
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
```

#### Methods are functions
Remember: a method is just a function with a receiver argument.
Here's Abs written as a regular function with no change in functionality.
```
type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
```
#### Can declare a method on non-struct types too
You can declare a method on non-struct types, too.
In this example we see a numeric type MyFloat with an Abs method.

You can only declare a method with a receiver whose type is defined in the same package as the method. 
You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).
```
type MyFloat float64 // ???

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
```
#### Pointer receivers 
You can declare methods with pointer receivers.

This means the receiver type has the literal syntax *T for some type T. (Also, T cannot itself be a pointer such as *int.)

For example, the Scale method here is defined on *Vertex.

#### pointer receivers vs value receivers
Methods with pointer receivers can modify the value to which the receiver points (as Scale does here). 
Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

Try removing the * from the declaration of the Scale function on line 16 and observe how the program's behavior changes.

With a value receiver, the Scale method operates on a copy of the original Vertex value. (This is the same behavior as
for any other function argument.) 
The Scale method must have a pointer receiver to change the Vertex value declared in the main function.

```
type Vertex struct {
	X, Y float64
}

func main() {
	v := Vertex{3, 4}	
	v.Scale(10)	
	fmt.Println(v.Abs())
}

// Methods with pointer receivers can modify the value to which the receiver points (as Scale does here)
func (v *Vertex) Scale(f float64) { 
	v.X = v.X * f
	v.Y = v.Y * f
}

// With a value receiver, the Scale method operates on a copy of the original Vertex value. (This is the same behavior as for any other function argument.) 
// The Scale method must have a pointer receiver to change the Vertex value declared in the main function.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```
#### Pointers and functions
Here we see the Abs and Scale methods rewritten as functions.
Again, try removing the * from line 16. Can you see why the behavior changes?
What else did you need to change for the example to compile?
(If you're not sure, continue to the next page.)
```
type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
```

#### Methods and pointer indirection (1)
*** Comparing the previous two programs, you might notice that functions with a pointer argument must take a pointer:

    var v Vertex
    ScaleFunc(v, 5)  // Compile error!
    ScaleFunc(&v, 5) // OK

while methods with pointer receivers take either a value or a pointer as the receiver when they are called:

    var v Vertex
    v.Scale(5)  // OK
    p := &v
    p.Scale(10) // OK

*** For the statement v.Scale(5), even though v is a value and not a pointer, 
the method with the pointer receiver is called automatically.
That is, as a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5) 
since the Scale method has a pointer receiver.

```
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//For the statement v.Scale(5), even though v is a value and not a pointer, 
//the method with the pointer receiver is called automatically.
//That is, as a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5) 
//since the Scale method has a pointer receiver.
func main() {
	v := Vertex4{3, 4}
	v.Scale4(2)    // Go interprets this as (&v).Scale4(2)
	fmt.Println(v) // {6 8}

	p := &Vertex4{3, 4}
	p.Scale4(2)    // here we are invoking the method on the address of Vertex via p
	fmt.Println(p) // &{6 8}

	ScaleFunc4(&v, 10)
	fmt.Println(v) // {60 80}

	ScaleFunc4(p, 8)
	fmt.Println(p) // 	&{48 64}

	//fmt.Println(v, p)
}
```
#### The reverse thing happens with the Methods and pointer indirection (2)
The equivalent thing happens in the reverse direction.
Functions that take a value argument must take a value of that specific type:

    var v Vertex
    fmt.Println(AbsFunc(v))  // OK
    fmt.Println(AbsFunc(&v)) // Compile error!

while methods with value receivers take either a value or a pointer as the receiver when they are called:

    var v Vertex
    fmt.Println(v.Abs()) // OK
    p := &v
    fmt.Println(p.Abs()) // OK

In this case, the method call p.Abs() is interpreted as (*p).Abs().

#### Choosing a value or pointer receiver
There are two reasons to use a pointer receiver.
The first is so that the method can modify the value that its receiver points to.
The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.
In this example, both Scale and Abs are with receiver type *Vertex, even though the Abs method needn't modify its receiver.
In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both. (We'll see why over the next few pages.)

```
type Vertex6 struct {
	X, Y float64
}

func main() {
	//v := Vertex6{3, 4}
	v := &Vertex6{3, 4}

	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs6())
	v.Scale6(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs6())
}

func (v *Vertex6) Abs6() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex6) Scale6(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
```

