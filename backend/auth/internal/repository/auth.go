package repository

import (
	"fmt"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/auth/internal/errors"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/auth/internal/hash"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/auth/internal/models"
	"net/http"
)

func (r *Repository) Register(req *models.RegisterRequest) (string, error) {
	query := fmt.Sprintf("INSERT INTO %s (first_name, last_name, email, role, password) VALUES ($1, $2, $3, $4, $5) RETURNING id", userTableName)
	var id string
	err := r.db.QueryRow(query, req.FirstName, req.LastName, req.Email, "user", req.Password).Scan(&id)
	if err != nil {
		return "", errors.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return id, nil
}

func (r *Repository) Login(req *models.LoginRequest) (string, string, error) {
	query := fmt.Sprintf("SELECT role, password, id, email FROM %s WHERE email = $1", userTableName)
	var user models.User
	err := r.db.QueryRow(query, req.Email).Scan(
		&user.Role,
		&user.Password,
		&user.ID,
		&user.Email,
	)
	if err != nil {
		return "", "", errors.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	if hash.CheckPasswordHash(req.Password, user.Password) {
		return user.ID, user.Role, nil
	}

	return "", "", errors.NewHTTPError(http.StatusUnauthorized, "wrong password")
}
