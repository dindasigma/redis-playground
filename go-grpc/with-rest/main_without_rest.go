package main

import (
	"log"
	"net"

	"github.com/dindasigma/my-playground/go-grpc/orders"
	"google.golang.org/grpc"
)

func mainWithoutRest() {
	grpcServer := grpc.NewServer()
	orderService := orders.UnimplementedOrderServiceServer{}
	orders.RegisterOrderServiceServer(grpcServer, &orderService)

	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start gRPC server: %v", err)
	}
}
