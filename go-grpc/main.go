package main

import (
	"log"
	"net"

	"github.com/dindasigma/my-playground/go-grpc/orders"
	"github.com/dindasigma/my-playground/go-grpc/rest"
	"google.golang.org/grpc"
)

const (
	grpcPort = "50051"
	restPort = "8080"
)

// marry REST + gRPC.
func main() {
	grpcServer := grpc.NewServer()
	orderService := orders.UnimplementedOrderServiceServer{}
	orders.RegisterOrderServiceServer(grpcServer, &orderService)

	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {

		// Serve() is a blocking call and hence the need to put this call in
		// in a goroutine.
		grpcServer.Serve(lis)
	}()

	restServer := rest.NewRestServer(orderService, restPort)

	// Start() is also blocking but this is ok for now as we need one blocking call
	// to prevent main() from exiting abruptly. we will refactor this logic soon!
	restServer.Start()

}
