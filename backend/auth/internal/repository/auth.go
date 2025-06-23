package repository

import (
	"fmt"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/auth/internal/errors"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/auth/internal/hash"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/auth/internal/models"
	"net/http"
)

func (r *Repository) Register(req *models.RegisterRequest, token string) (string, error) {
	query := fmt.Sprintf("INSERT INTO %s (first_name, last_name, email, password, token) VALUES ($1, $2, $3, $4, $5) RETURNING id", userTableName)
	var id string
	err := r.db.QueryRow(query, req.FirstName, req.LastName, req.Email, req.Password, token).Scan(&id)
	if err != nil {
		return "", errors.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return id, nil
}

func (r *Repository) Login(req *models.LoginRequest) (string, string, error) {
	query := fmt.Sprintf("SELECT role, password, id, email, status FROM %s WHERE email = $1", userTableName)
	var user models.User
	err := r.db.QueryRow(query, req.Email).Scan(
		&user.Role,
		&user.Password,
		&user.ID,
		&user.Email,
		&user.Status,
	)
	if err != nil {
		return "", "", errors.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	if user.Status != "active" {
		return "", "", errors.NewHTTPError(http.StatusUnauthorized, "account not active")
	}
	if hash.CheckPasswordHash(req.Password, user.Password) {
		return user.ID, user.Role, nil
	}

	return "", "", errors.NewHTTPError(http.StatusUnauthorized, "wrong password")
}

func (r *Repository) GetMe(id string) (*models.GetMeResponse, error) {
	query := fmt.Sprintf("SELECT id, first_name, last_name, email, role FROM %s WHERE id = $1", userTableName)
	var user models.GetMeResponse
	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Role,
	)
	if err != nil {
		return nil, errors.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return &user, nil
}

func (r *Repository) VerifyEmail(token string) error {
	res, err := r.db.Exec(`UPDATE users SET status = 'active', token = NULL WHERE token = $1`, token)
	if err != nil {
		return errors.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token")
	}
	return nil
}
