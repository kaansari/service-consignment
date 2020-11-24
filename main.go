// shippy-service-consignment/main.go
package main

import (
	"log"
	"net"

	// Import the generated protobuf code
	pb "github.com/kaansari/service-consignment/proto/consignment"
	"github.com/kaansari/service-consignment/shipping"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

func main() {

	var syncrepo shipping.ShippingRepository
	service := &shipping.ShippingService{syncrepo}

	// Set-up our gRPC server.
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// Register our service with the gRPC server, this will tie our
	// implementation into the auto-generated interface code for our
	// protobuf definition.
	pb.RegisterShippingServiceServer(s, service)

	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Println("Running on port:", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
