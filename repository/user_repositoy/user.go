package userrepositoy

import (
	"f-commerce/model"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepo interface {
	CreateCustomer(user *model.User) error
}

type userRepo struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewUserRepo(db *gorm.DB, log *zap.Logger) UserRepo {
	return &userRepo{db, log}
}

func (c *userRepo) CreateCustomer(user *model.User) error {

	if err := c.db.Create(&user).Error; err != nil {
		c.log.Error("failed to add user to database", zap.Error(err))
		return err
	}

	return nil

}
