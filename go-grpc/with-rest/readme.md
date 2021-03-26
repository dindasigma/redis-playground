# Go + REST + gRPC Demonstrate

## Installation
```
$ go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
$ go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc   
$ export GOPATH=$HOME/go
$ PATH=$PATH:$GOPATH/bin
```

## Compiling
Generate orders/order.pb.go and orders/order_grpc.pb.go from order.proto
```
protoc --go_out=./orders --go_opt=paths=source_relative \
--go-grpc_out=./orders --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false \
order.proto
```

- order.pb.go will contains a struct for every protobuf message type defined in order.proto.
- order_grpc.pb.go provides client/server code for interacting with the order service.

### Examples
- `/servers/grpc.go` : create gRPC server
- `/servers/rest.go` : create REST server and marrying with gRPC 
- `/grpc` : RPC function implementation

- `/client_grpc` : client gRPC example 

### Reference
[The Go Microservice Toolkit
](https://levelup.gitconnected.com/the-golang-microservice-toolkit-7521516ee4b)
