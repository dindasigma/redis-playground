package controllers

import (
	"context"
	"io"
	"log"

	pb "github.com/dindasigma/my-playground/go-grpc/exercises/client/proto/order"
	wrapper "github.com/golang/protobuf/ptypes/wrappers"
)

// Add Order: Unary RPC scenario
func AddOrder(ctx context.Context, client pb.OrderManagementClient, order pb.Order) (string, error) {

	res, err := client.Create(ctx, &order)
	if err != nil {
		return "", err
	}
	return res.Value, nil
}

// Get Order: Unary RPC scenario
func GetOrder(ctx context.Context, client pb.OrderManagementClient, id string) (*pb.Order, error) {
	retrievedOrder, err := client.Retrieve(ctx, &wrapper.StringValue{Value: id})
	if err != nil {
		return retrievedOrder, err
	}

	return retrievedOrder, nil
}

// Search Order: Server streaming scenario
func SearchOrders(ctx context.Context, client pb.OrderManagementClient, query string) ([]*pb.Order, error) {
	// create searchStream which has a Recv method
	result := []*pb.Order{}
	searchStream, _ := client.Search(ctx, &wrapper.StringValue{Value: query})

	for {
		// calling order responses one by one with the client stream's Recv()
		searchOrder, err := searchStream.Recv()
		// when the end of the stream is found return io.EOF
		if err == io.EOF {
			log.Print("EOF")
			break
		}

		if err == nil {
			result = append(result, searchOrder)
			// receiving the response
		}
	}
	return result, nil
}

// Update Orders: Client streaming scenario
func UpdateOrders(ctx context.Context, client pb.OrderManagementClient, data map[string]pb.Order) (*wrapper.StringValue, error) {
	updateStream, err := client.Update(ctx)

	if err != nil {
		log.Fatalf("%v.UpdateOrders(_) = _, %v", client, err)
	}

	// sending orders update via client stream
	for _, order := range data {
		// Updating order
		if err := updateStream.Send(&order); err != nil {
			log.Fatalf("%v.Send(%v) = %v", updateStream, order, err)
			return &wrapper.StringValue{}, err
		}
	}

	// closing the stream
	updateRes, err := updateStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", updateStream, err, nil)
		return &wrapper.StringValue{}, err
	}

	// receiving the response
	return updateRes, nil
}

// Process Order: Bidirectional streaming scenario
func ProcessOrders(ctx context.Context, client pb.OrderManagementClient, ids []string) {

	streamProcOrder, err := client.Process(ctx)
	if err != nil {
		log.Fatalf("%v.ProcessOrders(_) = _, %v", client, err)
	}

	for _, id := range ids {
		// send order
		if err := streamProcOrder.Send(&wrapper.StringValue{Value: id}); err != nil {
			log.Fatalf("%v.Send(%v) = %v", client, id, err)
		}
	}

	// create channel
	channel := make(chan struct{})
	// read the message from server in parallel
	go asncClientBidirectionalRPC(streamProcOrder, channel)
	// just mimic a delay
	//time.Sleep(time.Millisecond * 1000)
	// canceling the rpc
	// close in the end of the client stream
	if err := streamProcOrder.CloseSend(); err != nil {
		log.Fatal(err)
	}
	<-channel
}

func asncClientBidirectionalRPC(streamProcOrder pb.OrderManagement_ProcessClient, c chan struct{}) {
	for {
		// read message
		combinedShipment, errProcOrder := streamProcOrder.Recv()
		// end of stream
		if errProcOrder == io.EOF {
			break
		}
		log.Printf("Combined shipment : ", combinedShipment.OrdersList)
	}
	<-c
}
