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

### References
https://github.com/TutorialEdge/go-grpc-services-course
https://tutorialedge.net/golang/go-grpc-beginners-tutorial/
https://github.com/TutorialEdge/go-grpc-tutorial
