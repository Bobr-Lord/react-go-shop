package handler

import (
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/auth/internal/errors"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/auth/internal/middleware"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/auth/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
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
	id, err := h.svc.Register(&req)
	if err != nil {
		loger.Errorf("Registering failed: %s", err.Error())
		c.JSON(http.StatusInternalServerError, err)
		HTTPErr := errors.ParseHTTPError(err)
		c.JSON(HTTPErr.Code, HTTPErr.Message)
		return
	}
	resp := models.RegisterResponse{ID: id}
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
