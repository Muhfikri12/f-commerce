package userrepositoy

import (
	"errors"
	"f-commerce/model"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepo interface {
	RegisterUser(cust Regist) error
	GetUser(login *model.Login) (*model.User, error)
	UpdateCustomer(customer *model.Customer) error
	CreateCustomer(cust *model.Customer) error
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
		Where("email = ? OR username = ? OR id = ?", login.Input, login.Input, user.Id).
		First(&user).Error; err != nil {
		c.log.Error("Login error", zap.Error(err))
		return nil, errors.New("invalid email or username")
	}

	return &user, nil
}

type Regist struct {
	User     model.User
	Customer model.Customer
}

func (c *userRepo) RegisterUser(cust Regist) error {

	err := c.db.Transaction(func(tx *gorm.DB) error {

		if err := c.db.Create(&cust.User).Error; err != nil {
			c.log.Error("failed to add user to database", zap.Error(err))
			return err
		}

		cust.Customer.UserID = cust.User.Id

		if err := c.db.Create(&cust.Customer).Error; err != nil {
			c.log.Error("failed to add to database customer", zap.Error(err))
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil

}
