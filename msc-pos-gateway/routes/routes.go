package routes

import (
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	router *gin.Engine,
	authProxy *httputil.ReverseProxy,
	posProxy *httputil.ReverseProxy,
	inventoryProxy *httputil.ReverseProxy,
) {
	router.Any(
		"/api/auth/*path",
		gin.WrapH(authProxy),
	)

	router.Any(
		"/api/pos/*path",
		gin.WrapH(posProxy),
	)

	router.Any(
		"/api/inventory/*path",
		gin.WrapH(inventoryProxy),
	)
}
