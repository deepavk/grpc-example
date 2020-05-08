1. Define proto file
2.  To compile the proto file and generate code install protoc-gen
    go get github.com/golang/protobuf/protoc-gen-go
    export PATH="$PATH:$(go env GOPATH)/bin" 
3. mkdir chat 
   run: protoc --go_out=plugins=grpc:chat chat.proto

4. Start grpc server 
    ~~~
    In main.go -> 
        start grpc server
        add listeners for chat code 
    ~~~
     
5. Add chat.go in package chat to implement the ChatServer
 interface methods, register chat methods with grpc server 
 
6. Add client code to connect to grpc server.
 Refer the chat.pb.go for available methods and structs

https://tutorialedge.net/golang/go-grpc-beginners-tutorial/