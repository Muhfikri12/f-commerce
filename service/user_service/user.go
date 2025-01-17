package userservice

import (
	"f-commerce/helper"
	"f-commerce/model"
	"f-commerce/repository"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(regist *model.Register) error
	UpdateCustomer(id int, cust *model.CustomerData) error
}

type userService struct {
	Repo *repository.Repository
	Log  *zap.Logger
}

func NewUserService(Repo *repository.Repository, Log *zap.Logger) UserService {
	return &userService{Repo, Log}
}

func (us *userService) RegisterUser(regist *model.Register) error {

	password, err := bcrypt.GenerateFromPassword([]byte(regist.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	regist.Password = string(password)

	user := model.CustomerData{
		User: model.User{
			Password: regist.Password,
			Email:    regist.Email,
			Role:     "customer",
			Status:   "unverified",
			Username: regist.Fullname + helper.GenerateOTP(),
		},
		Customer: model.Customer{
			Fullname: regist.Fullname,
		},
	}

	if err := us.Repo.User.RegisterUser(&user); err != nil {
		return err
	}

	return nil
}
