package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-basics/api"
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

func createClient(clientNumber int, wg *sync.WaitGroup) {
	defer wg.Done()
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := api.NewSearchServiceClient(conn)
	for i := 1; i < 100; i++ {
		request := api.SearchRequest{
			MovieId:       "Avengers",
			CityId:        "Delhi",
			PageNumber:    int32(i),
			ResultPerPage: 10,
		}
		response, err := c.Search(context.Background(), &request)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Response from server: %s, client %d", response, clientNumber)
	}

}
