package http

import (
	"github.com/Maxim2710/golang-auth-lab/internal/service"
	"github.com/Maxim2710/golang-auth-lab/internal/transport/http/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(service *service.AuthService) *gin.Engine {
	router := gin.New()

	router.SetTrustedProxies([]string{"127.0.0.1"})

	authHandler := handler.NewAuthHandler(service)

	api := router.Group("/api")
	{
		api.POST("/register", authHandler.RegisterUser)
		api.POST("/login", authHandler.Login)
	}

	return router
}
