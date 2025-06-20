package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"log"
)

const RequestIdKey = "request_id"

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Request.URL.Path, c.Request.Method)
		if c.Request.Method == "OPTIONS" {
			log.Println("OPTIONS allowed by AuthUserMiddleware")
			c.AbortWithStatus(204)
			return
		}

		requestID := uuid.New().String()
		c.Set(RequestIdKey, requestID)
		logrus.WithFields(logrus.Fields{
			RequestIdKey: requestID,
		}).Infof("Request logged in, url: %s, method: %v", c.Request.RequestURI, c.Request.Method)
		c.Next()
	}
}
