package database

import (
	"finance/database/seeder"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func SeedAll(db *gorm.DB) error {
	log := &zap.Logger{}
	err := seeder.User(db)
	if err != nil {
		log.Error("filed to seed user seeder", zap.Error(err))
		return fmt.Errorf("filed to seed user seeder")
	}

	return nil
}
