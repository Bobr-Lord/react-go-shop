package main

import (
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/shop/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
}
