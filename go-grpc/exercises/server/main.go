package main

import (
	"log"
	"net"

	"github.com/dindasigma/my-playground/go-grpc/exercises/server/controllers"
	pb "github.com/dindasigma/my-playground/go-grpc/exercises/server/proto/order"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	zapLogger *zap.Logger
	port      = ":50051"
	crtFile   = "../certs/server.crt"
	keyFile   = "../certs/server.key"
)

func main() {
	// seeder
	seeder()

	// create listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	cert, err := credentials.NewServerTLSFromFile(crtFile, keyFile)
	if err != nil {
		log.Fatalln(err)
	}

	zapLogger, _ = zap.NewProduction()

	// Make sure that log statements internal to gRPC library are logged using the zapLogger as well.
	grpc_zap.ReplaceGrpcLoggerV2(zapLogger)
	// Create a server, make sure we put the grpc_ctxtags context before everything else.
	s := grpc.NewServer(
		// Enable TLS for all incoming connections.
		grpc.Creds(cert),

		// add midleware
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.UnaryServerInterceptor(zapLogger),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.StreamServerInterceptor(zapLogger),
		),
	)

	// attach controllers.Server to the order server
	pb.RegisterOrderManagementServer(s, &controllers.Server{})
	log.Printf("gRPC serve on %v", port)

	// serve
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func seeder() {
	controllers.OrderMap["102"] = pb.Order{Id: "102", Items: []string{"Google Pixel 3A", "Mac Book Pro"}, Destination: "Mountain View, CA", Price: 1800.00}
	controllers.OrderMap["103"] = pb.Order{Id: "103", Items: []string{"Apple Watch S4"}, Destination: "San Jose, CA", Price: 400.00}
	controllers.OrderMap["104"] = pb.Order{Id: "104", Items: []string{"Google Home Mini", "Google Nest Hub"}, Destination: "Mountain View, CA", Price: 400.00}
	controllers.OrderMap["105"] = pb.Order{Id: "105", Items: []string{"Amazon Echo"}, Destination: "San Jose, CA", Price: 30.00}
	controllers.OrderMap["106"] = pb.Order{Id: "106", Items: []string{"Amazon Echo", "Apple iPhone XS"}, Destination: "Mountain View, CA", Price: 300.00}
}
