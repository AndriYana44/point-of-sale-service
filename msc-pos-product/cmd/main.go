package main

import (
	"log"
	"os"

	"github.com/AndriYana44/msc-pos-product/config"
	"github.com/AndriYana44/msc-pos-product/controllers"
	"github.com/AndriYana44/msc-pos-product/repositories"
	"github.com/AndriYana44/msc-pos-product/routes"
	"github.com/AndriYana44/msc-pos-product/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	db := config.ConnectDatabase()

	productRepository := repositories.NewProductRepository(db)

	productService := services.NewProductService(
		productRepository,
	)

	productController := controllers.NewProductController(
		productService,
	)

	router := gin.Default()

	routes.SetupRoutes(
		router,
		productController,
	)

	for _, route := range router.Routes() {
		log.Printf("ROUTE: %-6s %s", route.Method, route.Path)
	}

	port := os.Getenv("APP_PORT")

	if port == "" {
		port = "8080"
	}

	log.Println("Starting product service on port:", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
