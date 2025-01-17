package database

import (
	"f-commerce/model"
	"fmt"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&model.User{},
		&model.Customer{},
		&model.Admin{},
	); err != nil {
		return fmt.Errorf("failed to migrate database: %v", err)
	}

	return nil
}
