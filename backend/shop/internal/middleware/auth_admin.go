package middleware

import (
	myJWT "github.com/Bobr-Lord/react-go-shop/tree/main/backend/shop/internal/jwt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	RoleKey = "role"
	IDKey   = "id"
)

func AuthAdminMiddleware(publicKeyPath string) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID, ok := c.Get(RequestIdKey)
		if !ok {
			requestID = "unknown"
		}
		loger := logrus.WithFields(logrus.Fields{
			"request_id": requestID,
		})
		loger.Info("auth middleware")
		tokenStr, err := c.Cookie("access_token")
		if err != nil {
			loger.Errorf("get cookie failed: %v", err)
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
			return
		}
		if tokenStr == "" {
			loger.Error("get cookie failed")
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
			return
		}
		rsa, err := myJWT.LoadRSAPublicKey(publicKeyPath)
		if err != nil {
			loger.Errorf("load public key failed: %v", err)
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
			return
		}
		claims, err := myJWT.ValidateJWT(tokenStr, rsa)
		if err != nil {
			loger.Errorf("validate jwt failed: %v", err)
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
			return
		}

		if claims["role"] != "admin" || claims["id"] == "" {
			loger.Error("unauthorized")
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
			return
		}

		c.Set(RoleKey, claims["id"])
		c.Set(RoleKey, claims["role"])
		c.Next()
	}
}
