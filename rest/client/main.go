package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"grpc-basics/rest/api"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	defer elapsed("page")()
	var wg sync.WaitGroup
	for i :=0 ; i < 100; i++ {
		wg.Add(1)
		go get(i, &wg)
	}
	wg.Wait()
}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}


func get(clientId int, wg *sync.WaitGroup) {
	defer wg.Done()
	var i int32
	for  i = 1; i < 500; i++ {
		makeARequest(i, clientId)
	}
}

func makeARequest(pageNumber int32, id int) {
	searchRequest := api.Request{
		MovieId:        "Avengers",
		CityId:         "Delhi",
		PageNumber:     pageNumber,
		ResultsPerPage: 10,
	}
	jsonReq, err := json.Marshal(searchRequest)
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/searchMovies", bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var searchResponse api.Response
	json.Unmarshal(bodyBytes, &searchResponse)
	fmt.Printf("API Response as struct: %+v from Client %d \n" , searchResponse, id)
}