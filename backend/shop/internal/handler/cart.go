package handler

import (
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/shop/internal/middleware"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/shop/internal/models"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/shop/internal/response"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) AddCartItem(c *gin.Context) {
	requestID, ok := c.Get(middleware.RequestIdKey)
	if !ok {
		requestID = "unknown"
	}
	loger := logrus.WithFields(logrus.Fields{
		"request_id": requestID,
	})
	loger.Info("AddCartItem")
	var req models.AddCartItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		loger.Error("ShouldBindJSON", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, ok := c.Get(middleware.IDKey)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "id not found"})
		return
	}

	err := h.svc.AddCartItem(req.IDItem, id.(string))
	if err != nil {
		loger.Error("svc.AddCartItem", err)
		httpErr := response.ParseHttpError(err)
		c.JSON(httpErr.Code, httpErr.Message)
		return
	}
	c.Status(http.StatusCreated)
}

func (h *Handler) DeleteCartItem(c *gin.Context) {
	requestID, ok := c.Get(middleware.RequestIdKey)
	if !ok {
		requestID = "unknown"
	}
	loger := logrus.WithFields(logrus.Fields{
		"request_id": requestID,
	})
	loger.Info("DeleteCartItem")
	idItem := c.Param("id_item")
	id, ok := c.Get(middleware.IDKey)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "id not found"})
		return
	}
	err := h.svc.DeleteCartItem(idItem, id.(string))
	if err != nil {
		loger.Error("svc.DeleteCartItem", err)
		httpErr := response.ParseHttpError(err)
		c.JSON(httpErr.Code, httpErr.Message)
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) GetCartItems(c *gin.Context) {
	requestID, ok := c.Get(middleware.RequestIdKey)
	if !ok {
		requestID = "unknown"
	}
	loger := logrus.WithFields(logrus.Fields{
		"request_id": requestID,
	})
	loger.Info("GetCartItems")
	id, ok := c.Get(middleware.IDKey)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "id not found"})
		return
	}
	res, err := h.svc.GetCartItems(id.(string))
	if err != nil {
		loger.Error("svc.GetCartItems", err)
		httpErr := response.ParseHttpError(err)
		c.JSON(httpErr.Code, httpErr.Message)
		return
	}
	c.JSON(http.StatusOK, res)
}
