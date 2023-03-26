package api

import (
	"net/http"

	"github.com/4klb/coffee-shop-microservices/gateway/internal/models"
	"github.com/4klb/coffee-shop-microservices/gateway/internal/response"
	"github.com/gin-gonic/gin"
)

//CreateOrder ..
func (h *Handle) CreateOrder(c *gin.Context) { //TODO
	var order models.Order

	if err := c.BindJSON(&order); err != nil {
		response.ResponseWithError(http.StatusBadRequest, err, c)
		return
	}

	if err := IsValidUUID(order.DishId); err != nil {
		response.ResponseWithError(http.StatusOK, err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order was created"})
}
