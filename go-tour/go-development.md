### Install Go
    https://go.dev/doc/tutorial/getting-started#install
### Download and install
     https://go.dev/doc/install

## Go modules

### Developing and publishing modules
    https://go.dev/doc/modules/developing
### Create a Go module
    https://go.dev/doc/tutorial/create-module
### Call your code from another module
    https://go.dev/doc/tutorial/call-module-code
### Get started with Go
    https://go.dev/doc/tutorial/getting-started
### Publishing a module 
    https://go.dev/doc/modules/publishing
#### Steps to publish a module.
    Open a command prompt and change to your module’s root directory in the local repository.
    Run go mod tidy, which removes any dependencies the module might have accumulated that are no longer necessary.
        $ go mod tidy
    Run go test ./... a final time to make sure everything is working.
    This runs the unit tests you’ve written to use the Go testing framework.
        $ go test ./...
        ok      example.com/mymodule       0.015s
    Tag the project with a new version number using the git tag command.
    For the version number, use a number that signals to users the nature of changes in this release. For more, see Module version numbering.
        $ git commit -m "mymodule: changes for v0.1.0"
        $ git tag v0.1.0
    Push the new tag to the origin repository.
        $ git push origin v0.1.0
    Make the module available by running the go list command to prompt Go to update its index of modules with information about the module you’re publishing.
    
    Precede the command with a statement to set the GOPROXY environment variable to a Go proxy. This will ensure that your request reaches the proxy.
        $ GOPROXY=proxy.golang.org go list -m example.com/mymodule@v0.1.0
    Developers interested in your module import a package from it and run the go get command just as they would with any other module. They can run the go get command for latest versions or they can specify a particular version, as in the following example:
        $ go get example.com/mymodule@v0.1.0

## Go Module
    github.com/vjammi/go-learn/go-tour/mr> go mod tidy
    go: finding module for package github.com/vjammi/go-learn/go-tour/mr
    go: downloading github.com/vjammi/go-learn/go-tour/mr v0.0.0-20221230025136-11275fa706d8
    go: downloading github.com/vjammi/go-learn v0.0.0-20221230025136-11275fa706d8
    go: found github.com/vjammi/go-learn/go-tour/mr in github.com/vjammi/go-learn/go-tour/mr v0.0.0-20221230025136-11275fa706d8

## Use the module in another module
    github.com\vjammi\go-learn\go-tour\main> go install  github.com/vjammi/go-learn/go-tour/mr
    github.com\vjammi\go-learn\go-tour\main> go install  github.com/vjammi/go-learn/go-tour/mr
            no required module provides package github.com/vjammi/go-learn/go-tour/mr; to add it:
            go get github.com/vjammi/go-learn/go-tour/mr
    github.com\vjammi\go-learn\go-tour\main> go get github.com/vjammi/go-learn/go-tour/mr
            go: downloading github.com/vjammi/go-learn/go-tour/mr v0.0.0-20230122043044-a2e662f89da0
            go: added github.com/vjammi/go-learn/go-tour/mr v0.0.0-20230122043044-a2e662f89da0

    ---
    github.com/vjammi/go-learn/go-tour/main> go mod init
        go: creating new go.mod: module github.com/vjammi/go-learn/go-tour/main
        go: to add module requirements and sums:
        go mod tidy
    github.com/vjammi/go-learn/go-tour/main> go mod tidy
    github.com/vjammi/go-learn/go-tour/main> go run main.go

## ??? Build and Deploy
    touch main.go
    touch main_test.go
    
    go build main.go
    go build -o main -ldflags=-X=main.version=${VERSION} main.go
    go test
    go run main.go
    Commit to git
    
    git init
    git add .
    git commit -m "initial commit"
    git remote add origin git@github.com:USERNAME/golang-pipeline
    git branch -M main
    git push -u origin main

    git tag v1.0.0
    git push --tags

### Overview of module development - Developing and publishing modules
    https://go.dev/doc/modules/developing

### Workflow for developing and publishing modules
    https://go.dev/doc/modules/developing#workflow

### High-level module development workflow including publishing - Module release and versioning workflow
    https://go.dev/doc/modules/release-workflow

### Start a module that others can use
    https://go.dev/doc/tutorial/create-module#start

### Managing dependencies
    https://go.dev/doc/modules/managing-dependencies

### Design and development
    https://go.dev/doc/modules/developing#design

### Package discovery
    https://go.dev/doc/modules/developing#discovery

### Search Packages or Symbols
    https://pkg.go.dev/