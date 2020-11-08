package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"grpc-basics/rest/api"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	get()
}

func get() {
	fmt.Println("Get")
	searchRequest := api.Request{
		MovieId:        "",
		CityId:         "",
		PageNumber:     0,
		ResultsPerPage: 0,
	}
	jsonReq, err := json.Marshal(searchRequest)
	req, err := http.NewRequest(http.MethodGet, "localhost:8080/searchMovies", bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	// Convert response body to Todo struct
	var searchResponse api.Response
	json.Unmarshal(bodyBytes, &searchResponse)
	fmt.Printf("API Response as struct:\n%+v\n", searchResponse)
}