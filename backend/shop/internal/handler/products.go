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
	requestID, ok := c.Get(middleware.RequestIdKey)
	if !ok {
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

	id, err := h.svc.CreateProduct(&req)
	if err != nil {
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
		ID: id,
	}
	loger.Infof("response %v", HTTPResp)

	c.JSON(200, HTTPResp)
}

func (h *Handler) GetAllProducts(c *gin.Context) {
	requestID, ok := c.Get(middleware.RequestIdKey)
	if !ok {
		requestID = "unknown"
	}
	loger := logrus.WithFields(logrus.Fields{
		"request_id": requestID,
	})
	loger.Info("Handle GetAllProducts")

	products, err := h.svc.GetAllProducts()
	if err != nil {
		loger.Errorf("failed to get products: %v", err)
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
	resp := models.GetAllProductsResponse{
		Products: products,
	}
	loger.Info("get products success")
	c.JSON(200, resp)
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	requestID, ok := c.Get(middleware.RequestIdKey)
	if !ok {
		requestID = "unknown"
	}
	loger := logrus.WithFields(logrus.Fields{
		"request_id": requestID,
	})
	loger.Info("Handle DeleteProduct")
	id := c.Param("id")
	if id == "" {
		errResp := response.HttpError{
			Code:    http.StatusBadRequest,
			Message: "invalid id",
		}
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	if err := h.svc.DeleteProduct(id); err != nil {
		loger.Errorf("failed to delete product: %v", err)
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
	HTTPResp := models.DeleteProductResponse{
		Response: "success",
	}
	loger.Infof("response %v", HTTPResp)
	c.JSON(200, HTTPResp)
}

func (h *Handler) GetAllProductsPrivate(c *gin.Context) {
	requestID, ok := c.Get(middleware.RequestIdKey)
	if !ok {
		requestID = "unknown"
	}
	loger := logrus.WithFields(logrus.Fields{
		"request_id": requestID,
	})
	loger.Info("Handle GetAllProductsPrivate")
	idUser, ok := c.Get(middleware.IDKey)
	if !ok {
		loger.Warn("failed to get user id")
		c.JSON(http.StatusInternalServerError, "failed to get user id")
		return
	}
	res, err := h.svc.GetAllProductsPrivate(idUser.(string))
	if err != nil {
		loger.Errorf("failed to get products: %v", err)
		httpErr := response.ParseHttpError(err)
		c.JSON(httpErr.Code, httpErr.Error())
		return
	}
	loger.Infof("response %v")
	resp := models.GetAllProductsResponse{
		Products: res,
	}
	c.JSON(200, resp)
}

func (h *Handler) DecrementProduct(c *gin.Context) {
	requestID, ok := c.Get(middleware.RequestIdKey)
	if !ok {
		requestID = "unknown"
	}
	loger := logrus.WithFields(logrus.Fields{
		"request_id": requestID,
	})
	loger.Info("Handle DecrementProduct")
	idUser, ok := c.Get(middleware.IDKey)
	if !ok {
		loger.Warn("failed to get user id")
		c.JSON(http.StatusInternalServerError, "failed to get user id")
		return
	}
	idProduct := c.Param("id")
	if idProduct == "" {
		loger.Warn("invalid id")
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}
	if err := h.svc.DecrementProduct(idUser.(string), idProduct); err != nil {
		loger.Errorf("failed to decrement product: %v", err)
		errResp := response.ParseHttpError(err)
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	c.Status(200)
}
