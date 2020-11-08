package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Handler struct {
}

func (h *Handler) GetCinemas(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var request Request
	json.Unmarshal(reqBody, &request)
	fmt.Printf("Request : %+v \n", request)

	fmt.Println("Endpoint Hit: /GetCinemas")
	_ = json.NewEncoder(w).Encode(Response{[]string{"PVR", "INOX"}})
}

type Request struct {
	MovieId        string `json:movieId`
	CityId         string `json:cityId`
	PageNumber     int32  `json:pageNumber`
	ResultsPerPage int32  `json:resultsPerPage`
}

type Response struct {
	CinemaIds []string `json:"cinemaIds"`
}
