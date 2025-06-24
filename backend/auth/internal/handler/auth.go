package handler

import (
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/auth/internal/errors"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/auth/internal/middleware"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/auth/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/mail"
)

func (h *Handler) Register(c *gin.Context) {
	requestID, ok := c.Get(middleware.RequestIdKey)
	if !ok {
		requestID = "unknown"
	}
	loger := logrus.WithFields(logrus.Fields{
		"request_id": requestID,
	})
	loger.Info("Register")
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		loger.Errorf("Binding failed: %s", err.Error())
		errResp := errors.NewHTTPError(http.StatusBadRequest, "bad request")
		c.JSON(http.StatusBadRequest, errResp)
		return
	}
	loger.Infof("%+v", req)
	if len(req.Password) < 8 || len(req.FirstName) < 3 || len(req.LastName) < 3 || !IsValidEmail(req.Email) {
		loger.Errorf("Invalid email")
		c.JSON(400, gin.H{"error": "Invalid fields"})
		return
	}

	id, err := h.svc.Register(&req)
	if err != nil {
		loger.Errorf("Registering failed: %s", err.Error())
		c.JSON(http.StatusInternalServerError, err)
		HTTPErr := errors.ParseHTTPError(err)
		c.JSON(HTTPErr.Code, HTTPErr.Message)
		return
	}
	resp := models.RegisterResponse{ID: id}

	loger.Info("Register success")
	c.JSON(http.StatusCreated, resp)
}

func (h *Handler) Login(c *gin.Context) {
	requestID, ok := c.Get(middleware.RequestIdKey)
	if !ok {
		requestID = "unknown"
	}
	loger := logrus.WithFields(logrus.Fields{
		"request_id": requestID,
	})
	loger.Info("Login")
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		loger.Errorf("Binding failed: %s", err.Error())
		errResp := errors.NewHTTPError(http.StatusBadRequest, "bad request")
		c.JSON(http.StatusBadRequest, errResp)
		return
	}
	token, err := h.svc.Login(&req)
	if err != nil {
		loger.Errorf("Logining failed: %s", err.Error())
		HTTPErr := errors.ParseHTTPError(err)
		c.JSON(HTTPErr.Code, HTTPErr.Message)
		return
	}
	c.SetCookie(
		"access_token",
		token,
		3600,
		"/",
		"",
		false,
		true,
	)
	resp := models.LoginResponse{Token: token}
	loger.Info("Login success")
	c.JSON(http.StatusCreated, resp)
}

func (h *Handler) GetMe(c *gin.Context) {
	requestID, ok := c.Get(middleware.RequestIdKey)
	if !ok {
		requestID = "unknown"
	}
	loger := logrus.WithFields(logrus.Fields{
		"request_id": requestID,
	})
	loger.Info("Handle GetMe")
	id, ok := c.Get(middleware.IDKey)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	user, err := h.svc.GetMe(id.(string))
	if err != nil {
		loger.Errorf("Getting user failed: %s", err.Error())
		httpErr := errors.ParseHTTPError(err)
		c.JSON(httpErr.Code, httpErr.Message)
		return
	}
	loger.Infof("Handle GetMe User: %+v", user)
	c.JSON(http.StatusOK, user)
}

func (h *Handler) VerifyEmail(c *gin.Context) {
	requestID, ok := c.Get(middleware.RequestIdKey)
	if !ok {
		requestID = "unknown"
	}
	loger := logrus.WithFields(logrus.Fields{
		"request_id": requestID,
	})
	loger.Info("Handle VerifyEmail")
	token := c.Query("token")
	if token == "" {
		c.String(400, "Токен не указан")
		return
	}
	if err := h.svc.VerifyEmail(token); err != nil {
		loger.Errorf("Verifying email failed: %s", err.Error())
		httpErr := errors.ParseHTTPError(err)
		if httpErr.Code == 500 {
			c.String(500, "Ошибка сервера")

		} else if httpErr.Code == 400 {
			c.String(400, "Недействительный или просроченный токен")
		} else {
			c.String(500, "Ошибка сервера")
		}
		return
	}
	loger.Infof("Handle VerifyEmail Token: %+v", token)
	c.Data(200, "text/html; charset=utf-8", []byte(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Email подтверждён</title>
			<style>
				body { font-family: sans-serif; background: #f5f5f5; padding: 50px; text-align: center; }
				.container { background: #fff; padding: 30px; border-radius: 8px; display: inline-block; }
				.success { color: green; font-size: 20px; margin-bottom: 10px; }
			</style>
		</head>
		<body>
			<div class="container">
				<div class="success">✅ Email успешно подтверждён!</div>
				<p>Вы можете закрыть это окно и продолжить использовать приложение.</p>
			</div>
		</body>
		</html>
	`))
}

//func SendConfirmationEmail(c *gin.Context) {
//	requestID, ok := c.Get(middleware.RequestIdKey)
//	if !ok {
//		requestID = "unknown"
//	}
//	loger := logrus.WithFields(logrus.Fields{
//		"request_id": requestID,
//	})
//	loger.Info("Handle SendConfirmationEmail")
//
//}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
