package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Server struct {
	srv *http.Server
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(addr string, handler *gin.Engine) error {
	s.srv = &http.Server{
		Addr:           addr,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		Handler:        handler,
	}
	return s.srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
