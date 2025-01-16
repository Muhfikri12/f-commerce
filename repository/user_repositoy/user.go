package userrepositoy

import (
	"f-commerce/model"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepo interface {
	CreateCustomer(customer *model.Customer) error
}

type userRepo struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewUserRepo(db *gorm.DB, log *zap.Logger) UserRepo {
	return &userRepo{db, log}
}

func (c *userRepo) CreateCustomer(customer *model.Customer) error {
	err := c.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(&customer.User).Error; err != nil {
			c.log.Error("failed to add user to database", zap.Error(err))
			return err
		}

		customer.ID = customer.User.Id

		if err := tx.Create(&customer).Error; err != nil {
			c.log.Error("failed to add employee to database", zap.Error(err))
			return err
		}

		return nil

	})

	if err != nil {
		c.log.Error("transaction failed", zap.Error(err))
		return err
	}

	return nil

}
