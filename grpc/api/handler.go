package api

import (
	. "context"
	"fmt"
)

type Handler struct {
}

func (h *Handler) Search(context Context, request *SearchRequest) (*SearchResponse, error) {
	fmt.Printf("Search server with request : %+v \n", request)
	return &SearchResponse{CinemaIds: []string{"pvr"}}, nil
}
