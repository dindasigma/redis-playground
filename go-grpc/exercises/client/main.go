package main

import (
	"context"
	"log"
	"time"

	"github.com/dindasigma/my-playground/go-grpc/exercises/client/controllers"
	pb "github.com/dindasigma/my-playground/go-grpc/exercises/client/proto/order"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

// create new grpc client
func newGrpcClient() pb.OrderManagementClient {
	// dial server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return pb.NewOrderManagementClient(conn)
}

func main() {
	// create new client
	client := newGrpcClient()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Add Order: unary RPC
	order1 := pb.Order{
		Id:          "101",
		Items:       []string{"iPhone XS", "Mac Book Pro"},
		Destination: "San Jose, CA",
		Price:       2300.00,
	}
	res, err := controllers.AddOrder(ctx, client, order1)
	if err != nil {
		log.Print("AddOrder error: ", err)
	}
	log.Print("AddOrder Response -> ", res)

	// Get Order: unary RPC
	retrievedOrder, err := controllers.GetOrder(ctx, client, "106")
	log.Print("GetOrder Response -> : ", retrievedOrder)

	// Search Order: Server-side streaming
	controllers.SearchOrders(ctx, client, "Google")

	// Update Orders: Client-side streaming
	dataUpdate := make(map[string]pb.Order)
	dataUpdate["102"] = pb.Order{Id: "102", Items: []string{"Update Google Pixel 3A", "Google Pixel Book"}, Destination: "Mountain View, CA", Price: 1100.00}
	dataUpdate["103"] = pb.Order{Id: "103", Items: []string{"Update Apple Watch S4", "Mac Book Pro", "iPad Pro"}, Destination: "San Jose, CA", Price: 2800.00}
	dataUpdate["104"] = pb.Order{Id: "104", Items: []string{"Update Google Home Mini", "Google Nest Hub", "iPad Mini"}, Destination: "Mountain View, CA", Price: 2200.00}
	controllers.UpdateOrders(ctx, client, dataUpdate)

	// Process Orders: Bi-directional streaming
	ids := []string{"102", "103"}
	controllers.ProcessOrders(ctx, client, ids)

}
