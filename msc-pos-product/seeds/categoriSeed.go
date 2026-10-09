package seeds

import (
	"log"

	"github.com/AndriYana44/msc-pos-product/models"
	"gorm.io/gorm"
)

func SeedCategories(db *gorm.DB) error {
	categories := []models.Category{
		{Name: "Makanan"},
		{Name: "Minuman"},
		{Name: "Snack"},
	}

	for _, category := range categories {
		var existing models.Category

		err := db.Where("name = ?", category.Name).
			First(&existing).Error

		if err == nil {
			log.Printf(
				"Category already exists: %s (ID: %d)",
				existing.Name,
				existing.ID,
			)
			continue
		}

		if err != gorm.ErrRecordNotFound {
			return err
		}

		if err := db.Create(&category).Error; err != nil {
			return err
		}

		log.Printf(
			"Category created: %s (ID: %d)",
			category.Name,
			category.ID,
		)
	}

	return nil
}
