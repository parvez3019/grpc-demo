package main

import (
	"google.golang.org/grpc"
	"grpc-basics/grpc/api"
	"log"
	"net"
)


func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	s := api.Handler{}

	grpcServer := grpc.NewServer()
	api.RegisterSearchServiceServer(grpcServer, &s)

	log.Println("Listening on localhost:8081")

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}