package handler

import (
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/shop/internal/middleware"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/shop/internal/models"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/shop/internal/response"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) CreateProduct(c *gin.Context) {
	requestID := c.GetHeader(middleware.RequestIdKey)
	if requestID == "" {
		requestID = "unknown"
	}
	loger := logrus.WithFields(logrus.Fields{
		"request_id": requestID,
	})
	loger.Info("Handle CreateProduct")
	var req models.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		loger.Errorf("invalid request body: %v", err)
		errResp := response.HttpError{
			Code:    http.StatusBadRequest,
			Message: "invalid request body",
		}
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	if err := h.svc.CreateProduct(&req); err != nil {
		loger.Errorf("failed to create product: %v", err)
		errResp := err
		if ok := response.IsHTTPError(err); !ok {
			errResp = &response.HttpError{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}

	HTTPResp := models.CreateProductResponse{
		Response: "success",
	}
	c.JSON(200, HTTPResp)
}
