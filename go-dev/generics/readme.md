## Generics 
```Reference: https://go.dev/doc/tutorial/generics```

### Add non-generic functions
```
In this step, you’ll add two functions that each add together the values of a map and return the total.

You’re declaring two functions instead of one because you’re working with two different types of maps: one that stores int64 values, and one that stores float64 values.
```

### Add a generic function to handle multiple types 
```
In this section, you’ll add a single generic function that can receive a map containing either integer or float values, effectively replacing the two functions you just wrote with a single function.

To support values of either type, that single function will need a way to declare what types it supports. Calling code, on the other hand, will need a way to specify whether it is calling with an integer or float map.

To support this, you’ll write a function that declares type parameters in addition to its ordinary function parameters. These type parameters make the function generic, enabling it to work with arguments of different types. You’ll call the function with type arguments and ordinary function arguments.

Each type parameter has a type constraint that acts as a kind of meta-type for the type parameter. Each type constraint specifies the permissible type arguments that calling code can use for the respective type parameter.

While a type parameter’s constraint typically represents a set of types, at compile time the type parameter stands for a single type – the type provided as a type argument by the calling code. If the type argument’s type isn’t allowed by the type parameter’s constraint, the code won’t compile.

Keep in mind that a type parameter must support all the operations the generic code is performing on it. For example, if your function’s code were to try to perform string operations (such as indexing) on a type parameter whose constraint included numeric types, the code wouldn’t compile.

In the code you’re about to write, you’ll use a constraint that allows either integer or float types.
```