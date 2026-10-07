package main

import (
	"log"

	"gin-learn/config"
	"gin-learn/seeds"
)

func main() {
	db := config.ConnectDatabase()

	if err := seeds.SeedRoles(db); err != nil {
		log.Fatal("Failed to seed roles:", err)
	}

	if err := seeds.SeedUsers(db); err != nil {
		log.Fatal("Failed to seed users:", err)
	}

	if err := seeds.SeedUserRoles(db); err != nil {
		log.Fatal("Failed to seed user roles:", err)
	}

	log.Println("Roles seeded successfully")
	log.Println("Users seeded Successfully")
	log.Println("User roles seeded successfully")
}
