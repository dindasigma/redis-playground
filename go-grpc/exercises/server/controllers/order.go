package controllers

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"

	pb "github.com/dindasigma/my-playground/go-grpc/exercises/server/proto/order"
	"github.com/golang/protobuf/ptypes/wrappers"
	wrapper "github.com/golang/protobuf/ptypes/wrappers"
)

const (
	orderBatchSize = 3
)

var OrderMap = make(map[string]pb.Order)

type Server struct{}

// Unary RPC Example
func (s *Server) Create(ctx context.Context, orderReq *pb.Order) (*wrapper.StringValue, error) {
	log.Printf("Order Added. ID : %v", orderReq.Id)
	OrderMap[orderReq.Id] = *orderReq
	return &wrapper.StringValue{Value: "Order Added: " + orderReq.Id}, nil
}

// Unary RPC Example
func (s *Server) Retrieve(ctx context.Context, orderId *wrapper.StringValue) (*pb.Order, error) {
	log.Printf("Order Given. ID : %v", orderId.Value)
	ord, exists := OrderMap[orderId.Value]
	if exists {
		return &ord, nil
	}

	return nil, errors.New("Order does not exist.")
}

// Server-side Streaming RPC Example
func (s *Server) Search(query *wrappers.StringValue, stream pb.OrderManagement_SearchServer) error {
	for key, order := range OrderMap {
		log.Print(key, order)
		for _, itemStr := range order.Items {
			log.Print(itemStr)
			// find mathing orders
			if strings.Contains(itemStr, query.Value) {
				// Send the matching orders in a stream
				err := stream.Send(&order)
				if err != nil {
					return fmt.Errorf("error sending message to stream : %v", err)
				}
				log.Print("Matching Order Found : " + key)
				break
			}
		}
	}
	return nil
}

// Client-side Streaming RPC Example
func (s *Server) Update(stream pb.OrderManagement_UpdateServer) error {

	ordersStr := "Updated Order IDs : "
	for {
		// read message from the client stream
		order, err := stream.Recv()
		if err == io.EOF {
			// Finished reading the order stream when found io.EOF
			return stream.SendAndClose(&wrapper.StringValue{Value: "Orders processed " + ordersStr})
		}

		if err != nil {
			return err
		}
		// Update order
		OrderMap[order.Id] = *order

		log.Printf("Order ID : %s - %s", order.Id, "Updated")
		ordersStr += order.Id + ", "
	}
}

// Bi-directional Streaming RPC Example
func (s *Server) Process(stream pb.OrderManagement_ProcessServer) error {
	batchMarker := 1
	var combinedShipmentMap = make(map[string]pb.CombinedShipment)
	for {
		// read order ids from the incomming stream
		orderId, err := stream.Recv()
		log.Printf("Reading Proc order : %s", orderId)

		// keep reading until the end of the stream is found
		if err == io.EOF {
			// Client has sent all the messages
			// Send remaining shipments
			log.Printf("EOF : %s", orderId)
			for _, shipment := range combinedShipmentMap {
				if err := stream.Send(&shipment); err != nil {
					return err
				}
			}
			return nil
		}
		if err != nil {
			log.Println(err)
			return err
		}

		// logic to organized orders into shipments based on the destinatino
		destination := OrderMap[orderId.GetValue()].Destination
		shipment, found := combinedShipmentMap[destination]

		if found {
			ord := OrderMap[orderId.GetValue()]
			shipment.OrdersList = append(shipment.OrdersList, &ord)
			combinedShipmentMap[destination] = shipment
		} else {
			comShip := pb.CombinedShipment{Id: "cmb - " + (OrderMap[orderId.GetValue()].Destination), Status: "Processed!"}
			ord := OrderMap[orderId.GetValue()]
			comShip.OrdersList = append(shipment.OrdersList, &ord)
			combinedShipmentMap[destination] = comShip
			log.Print(len(comShip.OrdersList), comShip.GetId())
		}

		// orders are process in batch. When the batch size is met, all the created combined shipment are streamed to the client.
		if batchMarker == orderBatchSize {
			// stream combined orders to client in batches
			for _, comb := range combinedShipmentMap {
				log.Printf("Shipping : %v -> %v", comb.Id, len(comb.OrdersList))

				// send combined shipment to the client
				if err := stream.Send(&comb); err != nil {
					return err
				}
			}
			batchMarker = 0
			combinedShipmentMap = make(map[string]pb.CombinedShipment)
		} else {
			batchMarker++
		}
	}
}
