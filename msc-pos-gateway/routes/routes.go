package routes

import (
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	router *gin.Engine,
	authProxy *httputil.ReverseProxy,
	productProxy *httputil.ReverseProxy,
) {
	router.Any(
		"/api/auth/*path",
		gin.WrapH(authProxy),
	)

	router.Any(
		"/api/products",
		gin.WrapH(productProxy),
	)

	router.Any(
		"/api/products/*path",
		gin.WrapH(productProxy),
	)
}
