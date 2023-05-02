## Go Documentation
    https://go.dev/doc/

## A Tour of Go
    https://go.dev/tour/welcome/3

### Basics
1. Packages, variables, and functions: Declaring variables, calling functions
    - Packages
    - Imports
    - Exported Names
    - Functions
    - Multiple Results
    - Named return values
    - Variables, Variables and initializers
    - Short variable declarations
    - Basic types
    - Zero values
    - Type conversions
    - Type inference
    - Constants
    - Numeric Constants
2. Flow control statements: for, if, else, switch and defer - how to control the flow of code with conditionals, loops, switches and defers.
    - Flow control statements: for, if. else, switch and defer
    - for, for is Go's "while"
    - f, If with short statement
    - if and else
    - switch, switch evaluation order, switch with no condition
    - defer, stacking defers
3. More types: structs, slices, and maps - Learn how to define types based on existing ones: this lesson covers structs, arrays, slices, and maps.
    - Pointers
    - Structs
    - Struct Literals
    - Arrays
    - Slices, Slices are like references to arrays, slices literals, slices defaults, slice length and capacity, Nil slices
    - Creating a slice with make, Nil slices, Creating a slice with make, Slices of slices
    - Range
    - Maps, Map literals
    - Mutating Maps
    - Function values, Function closures
### Methods and interfaces
Define methods on types, declare interfaces
1. Methods and interfaces. This covers methods and interfaces, constructs that define objects and their behavior.
    - Methods, methods and functions, Pointer receivers, Pointers and functions
    - Methods and pointer indirection
    - Choosing a value or a pointer receiver
    - Interfaces, Interfaces implemented implicitly
    - Interface values, Interface values with nil underlying values
    - Nil interface values
    - Empty interface
    - type assertions
    - Type switches
    - Stringers
    - Errors
    - Readers
    - Images
### Generics
1. Type Parameters
   Using type parameters in Go functions and structs - Go supports generic programming using type parameters
2. Generic Types

### Concurrency
1. Goroutines
2. Channels
3. Buffered Channels
4. Range and CLose
5. Select
6. Default Selection
7. sync.Mutex
   Exercise: Web crawler


### GO Installation
Installation
```
wget https://storage.googleapis.com/golang/go1.8.3.linux-amd64.tar.gz
sudo tar -xzf go1.8.3.linux-amd64.tar.gz -C /usr/local 

Go environment
export GOROOT=/usr/local/go
export GOBIN=$GOROOT/bin
export PATH=$PATH:$GOBIN
export GOPATH=$HOME/gopath
```

```
export GOROOT=$HOME/go
export GOBIN=$GOROOT/bin
export PATH=$PATH:$GOROOT/bin
```

### GOPATH and Workspace
- $GOPATH
- Package directory
- Compile packages
- Install remote packages

### Go Modules
- Create a module
    - ```go mod init <module_name>```
- Check if GO111MODULE value is set to "auto" by running ```go env```
- If not set to "auto" set it using ```go env -w GO111MODULE=auto```
- We can then run ``` go mod init```  ```go mod tidy```

### Go commands
1. go build
   compile packages and dependencies
2. go clean
   remove object files and cached files
3. go fmt and gofmt
   gofmt (reformat) package sources
4. go get
   add dependencies to current module and install them
5. go install
   compile and install packages and dependencies
6. godoc
   ,,,
8. go fix
   upgrade code from an old version before go1 to a new version after go1
9. go version // get information about your version of Go
10. go env // view environment variables about Go
11. go list // list all installed packages
12. go run // compile temporary files and run the application

## References
https://go.dev/doc/
https://go.dev/tour/list
https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/01.2.html
https://stackoverflow.com/questions/66894200/error-message-go-go-mod-file-not-found-in-current-directory-or-any-parent-dire
https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/01.3.html