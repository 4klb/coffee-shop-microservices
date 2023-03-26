package response

import (
	"github.com/gin-gonic/gin"
)

//ResponseMiddWithError ..
func ResponseMiddWithError(code int, message string, c *gin.Context) {
	c.JSON(code, gin.H{"error": message})
	c.Abort()
}

//ResponseWithError ..
func ResponseWithError(code int, statusText error, c *gin.Context) {
	c.JSON(code, gin.H{"statusText": statusText.Error()})
	c.Abort()
}
