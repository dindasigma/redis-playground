package servers

import (
	"net"

	"github.com/dindasigma/my-playground/go-grpc/orders"
	"google.golang.org/grpc"
)

// GrpcServer implements a gRPC Server for the Order service
type GrpcServer struct {
	server   *grpc.Server
	errCh    chan error
	listener net.Listener
}

// NewGrpcServer is a convenience func to create a GrpcServer
func NewGrpcServer(service orders.OrderServiceServer, port string) (GrpcServer, error) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return GrpcServer{}, err
	}
	server := grpc.NewServer()
	orders.RegisterOrderServiceServer(server, service)

	return GrpcServer{
		server:   server,
		listener: lis,
		errCh:    make(chan error),
	}, nil
}

// Start starts the server in the background, pushing any error to the error channel
func (g GrpcServer) Start() {
	go func() {
		g.errCh <- g.server.Serve(g.listener)
	}()
}

// Stop stops the gRPC server
func (g GrpcServer) Stop() {
	g.server.GracefulStop()
}

// Error returns the server's error channel
func (g GrpcServer) Error() chan error {
	return g.errCh
}
