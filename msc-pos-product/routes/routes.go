package routes

import (
	"github.com/AndriYana44/msc-pos-product/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	router *gin.Engine,
	productController *controllers.ProductController,
) {
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "ok",
		})
	})

	router.POST("/api/products", productController.Create)
	router.GET("/api/products", productController.FindAll)
	router.GET("/api/products/:id", productController.FindByID)
	router.PUT("/api/products/:id", productController.Update)
	router.DELETE("/api/products/:id", productController.Delete)
}
