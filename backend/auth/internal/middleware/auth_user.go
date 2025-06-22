package middleware

import (
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/auth/internal/jwt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	IDKey   = "IDKey"
	RoleKey = "RoleKey"
)

func AuthUserMiddleware(publicKeyPath string) gin.HandlerFunc {
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
		rsa, err := jwt.LoadRSAPublicKey(publicKeyPath)
		if err != nil {
			loger.Errorf("load public key failed: %v", err)
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
			return
		}
		claims, err := jwt.ValidateJWT(tokenStr, rsa)
		if err != nil {
			loger.Errorf("validate jwt failed: %v", err)
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
			return
		}

		if claims["role"] == "" || claims["id"] == "" {
			loger.Error("unauthorized")
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
			return
		}

		c.Set(IDKey, claims["id"])
		c.Set(RoleKey, claims["role"])
		c.Next()
	}
}
