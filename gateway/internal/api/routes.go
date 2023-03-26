package api

import "github.com/gin-gonic/gin"

//ApiSetupRouter ..
func ApiSetupRouter(router *gin.Engine, handle *Handle) {
	authorized := router.Group("/api")

	authorized.POST("/dish", handle.AddDish)
	authorized.GET("/dish", handle.GetAllDishes)
	authorized.POST("/order", handle.CreateOrder)
	authorized.GET("/order/:id", handle.GetOrder)
}
