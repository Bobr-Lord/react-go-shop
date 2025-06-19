package jwt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/shop/internal/response"
	"net/http"
	"os"
)

func LoadRSAPublicKey(path string) (*rsa.PublicKey, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(data)
	if block == nil {
		return nil, response.NewCustomError("failed to decode PEM block", http.StatusInternalServerError)
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	pubKey, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		return nil, response.NewCustomError("not an RSA public key", http.StatusInternalServerError)
	}

	return pubKey, nil
}
