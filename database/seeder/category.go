package seeder

import (
	"f-commerce/model"

	"gorm.io/gorm"
)

func SeedCategories(db *gorm.DB) error {
	categories := []model.Category{
		{Name: "Electronics"},
		{Name: "Fashion"},
		{Name: "Home Appliances"},
		{Name: "Books"},
		{Name: "Health & Beauty"},
		{Name: "Toys"},
		{Name: "Sports"},
		{Name: "Groceries"},
		{Name: "Automotive"},
		{Name: "Jewelry"},
	}

	for _, category := range categories {
		if err := db.FirstOrCreate(&category, model.Category{Name: category.Name}).Error; err != nil {
			return err
		}
	}

	return nil
}
