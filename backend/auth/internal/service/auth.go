package service

import (
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/auth/internal/errors"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/auth/internal/hash"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/auth/internal/jwt"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/auth/internal/models"
	"net/http"
)

func (s *Service) Register(req *models.RegisterRequest) (string, error) {
	hashPass, err := hash.HashPassword(req.Password)
	if err != nil {
		return "", errors.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	req.Password = hashPass
	id, err := s.repo.Register(req)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *Service) Login(req *models.LoginRequest) (string, error) {
	id, role, err := s.repo.Login(req)
	if err != nil {
		return "", err
	}

	privetKey, err := jwt.LoadRSAPrivateKeyPKCS8(s.cfg.PathPrivateKey)
	if err != nil {
		return "", errors.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}
	token, err := jwt.GenerateToken(id, role, privetKey)
	if err != nil {
		return "", errors.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}
	return token, nil
}

func (s *Service) GetMe(id string) (*models.GetMeResponse, error) {
	return s.repo.GetMe(id)
}
