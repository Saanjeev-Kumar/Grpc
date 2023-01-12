package main

import (
	"context"
	"goapp/internal"
	"goapp/internal/mongodb"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	user_pb "goapp/proto/pb"
	"google.golang.org/grpc"
)

func main() {
	// Connect to mongodb
	mongoclient, ctx, cancel, err := mongodb.Connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer mongodb.Close(mongoclient, ctx, cancel)
	_, err = net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatal(err)
	}
	// Register grpc-server
	userServer := internal.NewServer(context.Background(), mongoclient)
	grpcServer := grpc.NewServer()
	user_pb.RegisterUserServiceServer(grpcServer, userServer)
	// Register grpc-gateway server
	gwMux := runtime.NewServeMux()
	user_pb.RegisterUserServiceHandlerServer(context.Background(), gwMux, userServer)
	log.Printf("Server listening on localhost:8082")
	log.Fatal(http.ListenAndServe("localhost:8082", gwMux))
}
