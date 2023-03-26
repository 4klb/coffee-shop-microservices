package api

import (
	"net/http"

	"github.com/4klb/coffee-shop-microservices/gateway/internal/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//GetOrder ..
func (h *Handle) GetOrder(c *gin.Context) { //TODO
	id := c.Param("id")

	if err := IsValidUUID(id); err != nil {
		response.ResponseWithError(http.StatusOK, err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": ""})
}

//IsValidUUID ..
func IsValidUUID(u string) error {
	_, err := uuid.Parse(u)
	if err != nil {
		return err
	}

	return nil
}
