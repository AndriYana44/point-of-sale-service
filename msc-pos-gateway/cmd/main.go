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

	posProxy, err := proxy.CreateReverseProxy(
		os.Getenv("POS_SERVICE_URL"),
	)
	if err != nil {
		log.Fatal("Failed to create POS proxy:", err)
	}

	inventoryProxy, err := proxy.CreateReverseProxy(
		os.Getenv("INVENTORY_SERVICE_URL"),
	)
	if err != nil {
		log.Fatal("Failed to create inventory proxy:", err)
	}

	router := gin.Default()

	routes.SetupRoutes(
		router,
		authProxy,
		posProxy,
		inventoryProxy,
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
