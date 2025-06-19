package jwt

import (
	"crypto/rsa"
	"fmt"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/shop/internal/response"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func ValidateJWT(tokenString string, pubKey *rsa.PublicKey) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return pubKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}
	fmt.Println(token.Claims)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, response.NewCustomError("invalid claims structure", http.StatusInternalServerError)
	}

	return claims, nil
}
