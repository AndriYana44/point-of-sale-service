package main

import (
	"log"
	"net/http"

	"gin-learn/config"
	"gin-learn/controllers"
	"gin-learn/repositories"
	"gin-learn/routes"
	"gin-learn/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDatabase()

	userRepository := repositories.NewUserRepository(db)
	refreshTokenRepository := repositories.NewRefreshTokenRepository(db)

	authService := services.NewAuthService(
		userRepository,
		refreshTokenRepository,
	)

	authController := controllers.NewAuthController(
		authService,
	)

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	routes.SetupRoutes(
		router,
		authController,
	)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
