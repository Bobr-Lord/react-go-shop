package main

import (
	"context"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/shop/internal/config"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/shop/internal/handler"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/shop/internal/repository"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/shop/internal/server"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/shop/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		logrus.Fatal(err)
	}
	switch cfg.AppEnv {
	case "local":
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
		logrus.SetLevel(logrus.DebugLevel)
	case "prod":
		gin.SetMode(gin.ReleaseMode)
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.SetLevel(logrus.InfoLevel)
	}
	logrus.Infof("config: %+v", cfg)

	db, err := repository.NewPostgres(cfg)
	if err != nil {
		logrus.Fatal(err)
	}
	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	hndl := handler.NewHandler(svc)
	srv := server.NewServer()
	go func() {
		logrus.Info("server start")
		addr := cfg.ServerHost + ":" + cfg.ServerPort
		if err := srv.Run(addr, hndl.InitRouter()); err != nil {
			logrus.Errorf("server run error: %v", err)
		}
	}()
	wait := make(chan os.Signal, 1)
	signal.Notify(wait, syscall.SIGINT, syscall.SIGTERM)
	<-wait
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logrus.Errorf("server shutdown error: %v", err)
	}
	if err := db.Close(); err != nil {
		logrus.Error(err)
	}
	logrus.Info("Shutting down server...")
}
