## Go gRPC Microservice


### Shift of http rest api http/json to grpc/protobuf microservices 
Shift of http rest api (http/json) to grpc/protobuf microservices for applications that do not need to communicate with the front end.
1. Protobuf data format over json. Far smaller message sizes.
2. Typically, maintain you connection between your client and your server. 
   InArray client requests and responses 
   Client side streaming or Server side streaming or bidirectional streaming

### generate the Go specific gRPC code using the protoc tool:
chat.proto 
```
   syntax = "proto3";
   package chat;   
   option go_package = "github.com/vjammi/go-learn/go-grpc/chat;chat";   
   message Message {
      string body = 1;
   }   
   service ChatService {
      rpc SayHello(Message) returns (Message) {}
   }
```
Option 1
```
> cd C:\ds\workspace\golang\src\github.com\vjammi\go-learn\go-grpc
> C:\ds\tools\protoc-21.5-win64\bin\protoc --go_out=plugins=grpc:chat chat.proto
```
Option 2
```
> C:\ds\workspace\golang\src
> C:\ds\tools\protoc-21.5-win64\bin\protoc --go_out=plugins=grpc:. chat.proto
```

### Install protoc-gen-go (Go protocol buffer compiler plugin) in windows

Here is the step by step directions:
1. Download protoc-win32.zip from https://developers.google.com/protocol-buffers/docs/downloads
2. Unzip and add location of the protoc.exe to your PATH environment variable
3. Run `protoc --version` from command prompt to verify
4. Verify the your GOPATH environment variable is set
5. Run `go get -u github.com/golang/protobuf/protoc-gen-go` from command prompt. This should install the binary to %GOPATH%/bin
6. Add `%GOPATH%/bin` to your PATH environment variable
7. Open a new command prompt, navigate to your .proto file, run `protoc --go_out=. *.proto`

NOTE: if you are running from a text editor or ide, you may need to reboot after modifying your environment variables

### References
https://github.com/TutorialEdge/go-grpc-services-course
https://tutorialedge.net/golang/go-grpc-beginners-tutorial/
https://github.com/TutorialEdge/go-grpc-tutorial
https://groups.google.com/g/golang-nuts/c/Qs8d56uavVs