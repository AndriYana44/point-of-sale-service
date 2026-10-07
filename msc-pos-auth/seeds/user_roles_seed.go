package seeds

import (
	"gorm.io/gorm"
)

func SeedUserRoles(db *gorm.DB) error {
	userRoles := []map[string]interface{}{
		{
			"user_id": 1,
			"role_id": 1,
		},
		{
			"user_id": 2,
			"role_id": 2,
		},
		{
			"user_id": 3,
			"role_id": 2,
		},
		{
			"user_id": 4,
			"role_id": 3,
		},
	}

	return db.Table("map_user_roles").Create(&userRoles).Error
}
