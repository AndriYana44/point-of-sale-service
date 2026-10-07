package routes

import (
	"gin-learn/controllers"
	"gin-learn/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	router *gin.Engine,
	authController *controllers.AuthController,
) {
	api := router.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/login", authController.Login)
		auth.POST("/refresh", authController.RefreshToken)
	}

	protected := api.Group("")
	protected.Use(middleware.JWTAuth())
	{
		protected.GET("/auth/me", authController.Me)
	}
}
