package server

import (
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(address string) error {
	s.httpServer = &http.Server{Addr: address}

	return s.httpServer.ListenAndServe()
}
