package middleware

import (
	"encoding/base64"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/4klb/coffee-shop-microservices/gateway/internal/config"
	"github.com/4klb/coffee-shop-microservices/gateway/internal/models"
	"github.com/4klb/coffee-shop-microservices/gateway/internal/response"
)

//AuthMiddleware ..
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)

		if auth[0] != "Basic" {
			response.ResponseMiddWithError(401, "Unauthorized", c)
			return
		}

		user := GetUserPair(auth)

		if !AuthenticateUser(user) {
			response.ResponseMiddWithError(401, "Unauthorized", c)
			return
		}

		c.Next()
	}
}

//GetUserPair ..
func GetUserPair(auth []string) models.User {
	var user models.User

	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)

	user.Username = pair[0]
	user.Password = pair[1]

	return user
}

//AuthenticateUser ..
func AuthenticateUser(user models.User) bool {
	var users gin.Accounts = make(gin.Accounts)
	var account gin.Accounts = GetAccounts(users)

	for key, value := range account {
		if key == user.Username && value == user.Password {
			return true
		}
	}

	return false
}

//GetAccounts ..
func GetAccounts(users gin.Accounts) gin.Accounts {
	var user []string

	loginAuth := config.GetConfig().MiddlewareAuth.Login
	passwordAuth := config.GetConfig().MiddlewareAuth.Password

	user = append(user, loginAuth)
	user = append(user, passwordAuth)

	users[user[0]] = user[1]

	return users
}
