### Publishing a module [https://go.dev/doc/modules/publishing]
Steps to publish a module.
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

### Developing and publishing modules [https://go.dev/doc/modules/developing]