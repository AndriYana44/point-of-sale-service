package seeds

import (
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) error {
	roles := []map[string]interface{}{
		{
			"name":        "super admin",
			"code":        "001",
			"description": "Super Administrator with full access to the system",
		},
		{
			"name":        "admin",
			"code":        "002",
			"description": "Administrator with limited access to the system",
		},
		{
			"name":        "cashier",
			"code":        "003",
			"description": "Cashier with access to sales and transactions",
		},
	}

	return db.Table("mst_roles").Create(&roles).Error
}
