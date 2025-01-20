package userrepositoy

import (
	"errors"
	"f-commerce/model"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepo interface {
	RegisterUser(cust *model.CustomerData) error
	GetUser(login *model.Login) (*model.User, error)
	UpdateUser(id int, user *model.User) error
	UpdateCustomer(id int, customer *model.Customer) error
	UpdateProfile(id int, image string) error
	UpdateRole(id int) error
	UpdateAdmin(id int, admin *model.Admin) error
	NonactiveAccount(id int) error
}

type userRepo struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewUserRepo(db *gorm.DB, log *zap.Logger) UserRepo {
	return &userRepo{db, log}
}

func (c *userRepo) GetUser(login *model.Login) (*model.User, error) {

	user := model.User{}
	if err := c.db.Table("users").
		Where("email = ? OR username = ?", login.Input, login.Input).
		First(&user).Error; err != nil {
		c.log.Error("Login error", zap.Error(err))
		return nil, errors.New("invalid email or username")
	}

	return &user, nil
}

func (c *userRepo) RegisterUser(cust *model.CustomerData) error {

	err := c.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&cust.User).Error; err != nil {
			c.log.Error("failed to add user to database", zap.Error(err))
			return err
		}

		if cust.User.Role != "admin" {
			cust.Customer.UserID = cust.User.Id
			if err := tx.Create(&cust.Customer).Error; err != nil {
				c.log.Error("failed to add customer to database", zap.Error(err))
				return err
			}
		} else {
			cust.Admin.UserID = cust.User.Id
			if err := tx.Create(&cust.Admin).Error; err != nil {
				c.log.Error("failed to add admin to database", zap.Error(err))
				return err
			}
		}

		return nil
	})

	return err
}

func (c *userRepo) UpdateUser(id int, user *model.User) error {

	result := c.db.Table("users").Where("id = ?", id).Updates(&user)

	if result.RowsAffected == 0 {
		return fmt.Errorf("user with id %d not found", id)
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}
