# My Golang experiments

## A Tour of Go
### Basics
1. Declaring variables, calling functions, and all the things you need to know before moving to the next lessons.
2. Packages, variables, and functions.
3. Flow control statements: for, if, else, switch and defer - Learn how to control the flow of your code with conditionals, loops, switches and defers.
4. More types: structs, slices, and maps - Learn how to define types based on existing ones: this lesson covers structs, arrays, slices, and maps.

### Methods and interfaces
Learn how to define methods on types, how to declare interfaces, and how to put everything together.
1. Methods and interfaces. This lesson covers methods and interfaces, the constructs that define objects and their behavior.

### Generics
Learn how to use type parameters in Go functions and structs.
1. Generics - Go supports generic programming using type parameters. This lesson shows some examples for employing generics in your code.

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
7. 
8. go fix // upgrade code from an old version before go1 to a new version after go1
9. go version // get information about your version of Go
10. go env // view environment variables about Go
11. go list // list all installed packages
12. go run // compile temporary files and run the application

## Building Rest API in golang
Building a REST API in golang




## References
https://go.dev/tour/list
https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/01.2.html
https://stackoverflow.com/questions/66894200/error-message-go-go-mod-file-not-found-in-current-directory-or-any-parent-dire
https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/01.3.html

https://github.com/TutorialEdge/go-rest-api-course
