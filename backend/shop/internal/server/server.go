package server

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	srv *http.Server
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(addr string) error {
	s.srv = &http.Server{
		Addr:           addr,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
	}
	return s.srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
