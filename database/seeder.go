package database

import (
	"f-commerce/database/seeder"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func SeedAll(db *gorm.DB) error {

	log, _ := zap.NewProduction()
	defer log.Sync()

	err := seeder.User(db)
	if err != nil {
		log.Error("Failed to seed user data", zap.Error(err))
		return fmt.Errorf("failed to seed user data: %w", err)
	}

	err = seeder.SeedCategories(db)
	if err != nil {
		log.Error("Failed to seed category data", zap.Error(err))
		return fmt.Errorf("failed to seed category data: %w", err)
	}

	log.Info("Successfully seeded all data")
	return nil
}
