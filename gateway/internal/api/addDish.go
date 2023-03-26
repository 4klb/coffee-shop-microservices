package api

import (
	"net/http"

	"github.com/4klb/coffee-shop-microservices/gateway/internal/models"
	"github.com/4klb/coffee-shop-microservices/gateway/internal/response"
	"github.com/gin-gonic/gin"
)

//AddDish ..
func (h *Handle) AddDish(c *gin.Context) { //TODO
	var dish models.Dish

	if err := c.BindJSON(&dish); err != nil {
		response.ResponseWithError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data was added"})
}
