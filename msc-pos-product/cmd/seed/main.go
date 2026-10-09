package main

import (
	"log"

	"github.com/AndriYana44/msc-pos-product/config"
	"github.com/AndriYana44/msc-pos-product/seeds"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found; using environment variables")
	}

	db := config.ConnectDatabase()

	if err := seeds.SeedCategories(db); err != nil {
		log.Fatal("Failed to seed categories: ", err)
	}

	log.Println("Category seeding completed successfully")
}
