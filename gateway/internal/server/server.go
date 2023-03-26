package server

import (
	"github.com/4klb/coffee-shop-microservices/gateway/internal/api"
	"github.com/4klb/coffee-shop-microservices/gateway/internal/config"
	mw "github.com/4klb/coffee-shop-microservices/gateway/internal/middleware"
	"github.com/gin-gonic/gin"
)

//Run ..
func Run() {
	router := gin.Default()

	handle := api.GetHandle()

	router.Use(mw.AuthMiddleware())

	api.ApiSetupRouter(router, handle)

	router.Run(config.GetConfig().Api.Port)
}
