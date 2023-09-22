package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	*http.Server
}

func NewServer(port string, router http.Handler) Server {
	return Server{
		&http.Server{
			Addr:    fmt.Sprintf(":%s", port),
			Handler: router,
		},
	}
}

func (s *Server) Start() error {
	return s.ListenAndServe()
}

func (s *Server) Close() error {
	return s.Server.Close()
}
