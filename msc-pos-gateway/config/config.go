package config

import "os"

type Config struct {
	GatewayPort         string
	AuthServiceURL      string
	POSServiceURL       string
	InventoryServiceURL string
}

func Load() Config {
	return Config{
		GatewayPort:         os.Getenv("GATEWAY_PORT"),
		AuthServiceURL:      os.Getenv("AUTH_SERVICE_URL"),
		POSServiceURL:       os.Getenv("POS_SERVICE_URL"),
		InventoryServiceURL: os.Getenv("INVENTORY_SERVICE_URL"),
	}
}
