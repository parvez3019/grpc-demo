package main

import (
	"log"
	"net/http"
	"grpc-basics/rest/api"
)

func main() {
	handleRequests()
}

func handleRequests() {
	handler := api.Handler{}
	http.HandleFunc("/searchMovies", handler.GetCinemas)
	log.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
