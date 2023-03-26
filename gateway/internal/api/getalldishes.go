package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetAllDishes ..
func (h *Handle) GetAllDishes(c *gin.Context) { //TODO
	c.JSON(http.StatusOK, gin.H{"message": ""})
}
