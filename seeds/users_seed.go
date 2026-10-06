package seeds

import (
	"log"

	"gin-learn/helpers"

	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) error {
	hashedPassword, err := helpers.HashPassword("password123")
	if err != nil {
		log.Fatal("Failed to hash password:", err)
	}

	users := []map[string]interface{}{
		{
			"name":     "Super Administrator",
			"email":    "super.admin@example.com",
			"type":     "super_admin",
			"password": hashedPassword,
		},
		{
			"name":     "Administrator",
			"email":    "storage.admin@example.com",
			"type":     "admin_storage",
			"password": hashedPassword,
		},
		{
			"name":     "Inventory Administrator",
			"email":    "inventory.admin@example.com",
			"type":     "admin_inventory",
			"password": hashedPassword,
		},
		{
			"name":     "Cashier",
			"email":    "cashier@example.com",
			"type":     "cashier",
			"password": hashedPassword,
		},
	}

	return db.Table("mst_users").Create(&users).Error
}
