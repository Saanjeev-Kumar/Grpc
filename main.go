package main

import (
	"context"
	"fmt"
	"goapp/internal"
	"goapp/internal/mongodb"
	"log"
	"net"
	"net/http"

	user_pb "goapp/proto/pb"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	// Connect to mongodb
	mongoclient, ctx, cancel, err := mongodb.Connect("mongodb://localhost:27017")
	//mongodb+srv://saanjeevkumar:Sanju1102%40@cluster0.to4ajhd.mongodb.net/
	//mongodb+srv://mdm-eap-search-qa:ANurDXAvpsFEvvGE@eap-search-qa.1ls1d.mongodb.net/mls?retryWrites=true&w=majority
	//mongodb+srv://mdm-eap-search-qa:ANurDXAvpsFEvvGE@eap-zap-contacts-prd.iccne.mongodb.net/test?retryWrites=true&w=majority
	fmt.Println("main 1st connection", mongoclient)
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
	fmt.Println("MongoDb Server storage to intergrete with grpc server ", userServer)

	grpcServer := grpc.NewServer()
	fmt.Println("2nd Main grpcServer ", grpcServer)

	user_pb.RegisterUserServiceServer(grpcServer, userServer)
	//Register with storage and grpc server

	// Register grpc-gateway server
	gwMux := runtime.NewServeMux()
	user_pb.RegisterUserServiceHandlerServer(context.Background(), gwMux, userServer)
	// opts := []grpc.DialOption{grpc.WithInsecure()}
	// user_pb.RegisterUserServiceHandlerFromEndpoint(context.Background(),gwMux,fmt.Sprintf("localhost:8081"),opts)

	log.Printf("Server listening on localhost:8082")
	log.Fatal(http.ListenAndServe("localhost:8082", gwMux))
}
