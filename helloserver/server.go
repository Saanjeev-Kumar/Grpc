package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "Grpc/hellopb"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *server) Hello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	name := request.Name
	response := &pb.HelloResponse{
		Greeting: "Hello " + name,
	}
	return response, nil
}

func main() {
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})

	s.Serve(lis)
}
