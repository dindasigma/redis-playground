package main

import (
	"context"
	"log"
	"time"

	pb "github.com/dindasigma/my-playground/go-grpc/orders"
	wrapper "github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

// create new grpc client
func newGrpcClient() pb.OrderServiceClient {
	// dial server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return pb.NewOrderServiceClient(conn)
}

func main() {
	// create new client
	client := newGrpcClient()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// get rpc list
	res, err := client.List(ctx, &wrapper.StringValue{Value: "macbook air"})
	if err != nil {
		log.Print("AddOrder error: ", err)
	}
	log.Print("AddOrder Response -> ", res)
}
