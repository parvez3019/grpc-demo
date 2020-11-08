package api

import (
	. "context"
	"fmt"
)

type Server struct {
}

func (s *Server) Search(context Context, request *SearchRequest) (*SearchResponse, error) {
	fmt.Printf("Search server with request : %+v \n", request)
	return &SearchResponse{CinemaIds: []string{"pvr"}}, nil
}
