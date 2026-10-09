package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"msc-pos-gateway/proxy"
	"msc-pos-gateway/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	authProxy, err := proxy.CreateReverseProxy(
		os.Getenv("AUTH_SERVICE_URL"),
	)
	if err != nil {
		log.Fatal("Failed to create auth proxy:", err)
	}

	productProxy, err := proxy.CreateReverseProxy(
		os.Getenv("PRODUCT_SERVICE_URL"),
	)
	if err != nil {
		log.Fatal("Failed to create product proxy:", err)
	}

	router := gin.Default()

	routes.SetupRoutes(
		router,
		authProxy,
		productProxy,
	)

	port := os.Getenv("GATEWAY_PORT")

	if port == "" {
		port = "8080"
	}

	log.Println("Gateway running on port", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
