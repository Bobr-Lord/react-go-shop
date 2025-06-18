package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const RequestIdKey = "request_id"

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.New().String()
		c.Set(RequestIdKey, requestID)
		logrus.WithFields(logrus.Fields{
			RequestIdKey: requestID,
		}).Infof("Request logged in, url: %s, method: %v", c.Request.RequestURI, c.Request.Method)
		c.Next()
	}
}
