package main

import (
	"fmt"
	"io"
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "../customer"
)

const (
	address = "0.0.0.0:50051"
)

// createCustomer calls the RPC method CreateCustomer of CustomerServer
func createCustomer(client pb.CustomerClient, customer *pb.CustomerRequest) {
	resp, err := client.CreateCustomer(context.Background(), customer)
	if err != nil {
		log.Fatalf("Could not create Customer: %v", err)
	}
	if resp.Success {
		log.Printf("A new Customer has been added with id: %d", resp.Id)
	}
}

// getCustomers calls the RPC method GetCustomers of CustomerServer
func getCustomers(client pb.CustomerClient, filter *pb.CustomerFilter) {
	// calling the streaming API
	stream, err := client.GetCustomers(context.Background(), filter)
	if err != nil {
		log.Fatalf("Error on get customers: %v", err)
	}
	var idsum int32
	for {
		customer, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetCustomers(_) = _, %v", client, err)
		}
		idsum += customer.Id
	}
	// fmt.Printf("idsum=%v\n", idsum)
}

func main() {
	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Creates a new CustomerClient
	client := pb.NewCustomerClient(conn)

	// Create 10 identical customers.
	for i := 0; i < 10; i++ {
		customer := &pb.CustomerRequest{
			Id:    int32(101 + i),
			Name:  "Shiju Varghese",
			Email: "shiju@xyz.com",
			Phone: "732-757-2923",
			Addresses: []*pb.CustomerRequest_Address{
				&pb.CustomerRequest_Address{
					Street:            "1 Mission Street",
					City:              "San Francisco",
					State:             "CA",
					Zip:               "94105",
					IsShippingAddress: false,
				},
				&pb.CustomerRequest_Address{
					Street:            "Greenfield",
					City:              "Kochi",
					State:             "KL",
					Zip:               "68356",
					IsShippingAddress: true,
				},
			},
		}
		createCustomer(client, customer)
	}

	const numTrials = 100
	filter := &pb.CustomerFilter{Keyword: "Shiju"}
	t1 := time.Now()
	for i := 0; i < numTrials; i++ {
		getCustomers(client, filter)
	}
	elapsed := time.Since(t1)
	fmt.Printf("trials=%d elapsed=%v per_call=%v\n", numTrials, elapsed, elapsed/time.Duration(numTrials))
}
