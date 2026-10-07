package main

import (
	"log"

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

	routes.SetupRoutes(
		router,
		authController,
	)

	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
