package main

import (
	"google.golang.org/grpc"
	"grpc-basics/api"
	"log"
	"net"
)


func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	s := api.Server{}

	grpcServer := grpc.NewServer()
	api.RegisterSearchServiceServer(grpcServer, &s)

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}