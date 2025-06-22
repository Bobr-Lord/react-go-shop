package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(addr string, hndl *gin.Engine) error {
	s.httpServer = &http.Server{
		Addr:           addr,
		Handler:        hndl,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
