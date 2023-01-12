package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "Grpc/hellopb"
	"log"
)

func main() {
	fmt.Println("Hello client running ...")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := pb.NewHelloServiceClient(cc)
	request := &pb.HelloRequest{Name: "Saanjeev"}

	resp, err := client.Hello(context.Background(), request)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(resp)
	fmt.Printf("Receive response => [%v]", resp.Greeting)

}
