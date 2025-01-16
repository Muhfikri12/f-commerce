package seeder

import (
	"f-commerce/model"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func User(db *gorm.DB) error {
	password, err := bcrypt.GenerateFromPassword([]byte("superadmin"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	superAdmin := model.User{
		Email:     "superadmin@example.com",
		Password:  string(password),
		Role:      "super_admin",
		Username:  "superadmin",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := db.Create(&superAdmin).Error; err != nil {
		return err
	}

	log.Println("Super Admin seeded successfully.")
	return nil
}
