package grpc

import (
	"context"
	"fmt"
	"log"

	wrapper "github.com/golang/protobuf/ptypes/wrappers"
)

type Server struct{}

func (s *Server) Create(ctx context.Context, orderReq *wrapper.StringValue) (*wrapper.StringValue, error) {
	message := fmt.Sprintf("Create with Request : %v", orderReq)
	log.Print(message)
	return &wrapper.StringValue{Value: message}, nil
}

func (s *Server) Retrieve(ctx context.Context, orderReq *wrapper.StringValue) (*wrapper.StringValue, error) {
	message := fmt.Sprintf("Retrieve with Request : %v", orderReq)
	log.Print(message)
	return &wrapper.StringValue{Value: message}, nil
}

func (s *Server) Update(ctx context.Context, orderReq *wrapper.StringValue) (*wrapper.StringValue, error) {
	message := fmt.Sprintf("Update with Request : %v", orderReq)
	log.Print(message)
	return &wrapper.StringValue{Value: message}, nil
}

func (s *Server) Delete(ctx context.Context, orderReq *wrapper.StringValue) (*wrapper.StringValue, error) {
	message := fmt.Sprintf("Delete with Request : %v", orderReq)
	log.Print(message)
	return &wrapper.StringValue{Value: message}, nil
}

func (s *Server) List(ctx context.Context, orderReq *wrapper.StringValue) (*wrapper.StringValue, error) {
	message := fmt.Sprintf("List with Request : %v", orderReq)
	log.Print(message)
	return &wrapper.StringValue{Value: message}, nil
}
