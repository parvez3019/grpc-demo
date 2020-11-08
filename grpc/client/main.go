package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	api2 "grpc-basics/grpc/api"
	"log"
	"sync"
	"time"
)

func main() {
	defer elapsed("page")()
	var wg sync.WaitGroup
	for i :=0 ; i < 100; i++ {
		wg.Add(1)
		go createClient(i, &wg)
	}
	wg.Wait()

}
func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func createClient(clientId int, wg *sync.WaitGroup) {
	defer wg.Done()
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := api2.NewSearchServiceClient(conn)
	var i int32
	for  i = 1; i < 500; i++ {
		request := api2.SearchRequest{
			MovieId:       "Avengers",
			CityId:        "Delhi",
			PageNumber:    i,
			ResultPerPage: 10,
		}
		response, err := c.Search(context.Background(), &request)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Response from server: %s for client %d", response, clientId)
	}

}
