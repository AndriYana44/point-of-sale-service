package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"msc-pos-gateway/proxy"
)

func TestProductsCollectionIsProxiedWithoutRedirect(t *testing.T) {
	gin.SetMode(gin.TestMode)

	var upstreamPath string
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upstreamPath = r.URL.Path
		w.WriteHeader(http.StatusOK)
	}))
	defer upstream.Close()

	productProxy, err := proxy.CreateReverseProxy(upstream.URL)
	if err != nil {
		t.Fatal(err)
	}

	router := gin.New()
	SetupRoutes(router, nil, productProxy)

	gateway := httptest.NewServer(router)
	defer gateway.Close()

	response, err := http.Get(gateway.URL + "/api/products")
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, response.StatusCode)
	}
	if upstreamPath != "/api/products" {
		t.Fatalf("expected upstream path /api/products, got %q", upstreamPath)
	}
}
