package userservice

import (
	"f-commerce/helper"
	"f-commerce/model"
	"f-commerce/repository"
	"time"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(regist *model.Register) error
	UpdateCustomer(token string, cust *model.Customer) error
	UpdateProfile(token string, image string) error
	UpdateRole(token string) error
	UpdateAdmin(token string, admin *model.Admin) error
	UpdateUser(token string, user *model.User) error
}

type userService struct {
	Repo *repository.Repository
	Log  *zap.Logger
	jwt  *helper.Jwt
}

func NewUserService(Repo *repository.Repository, Log *zap.Logger, jwt *helper.Jwt) UserService {
	return &userService{Repo, Log, jwt}
}

func (us *userService) RegisterUser(regist *model.Register) error {

	password, err := bcrypt.GenerateFromPassword([]byte(regist.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	status := "active"
	if regist.Role != "admin" {
		status = "unverified"
	}

	user := model.CustomerData{
		User: model.User{
			Password: string(password),
			Email:    regist.Email,
			Role:     regist.Role,
			Status:   status,
			Username: regist.Fullname + helper.GenerateOTP(),
		},
	}

	if regist.Role != "admin" {
		user.Customer = model.Customer{
			Fullname: regist.Fullname,
		}
	} else {
		user.Admin = model.Admin{
			Fullname: regist.Fullname,
		}
	}

	if err := us.Repo.User.RegisterUser(&user); err != nil {
		return err
	}

	return nil
}

func (us *userService) UpdateUser(token string, user *model.User) error {

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(password)

	id, err := us.jwt.ParsingID(token)
	if err != nil {
		return err
	}

	user.UpdatedAt = time.Now()

	if err := us.Repo.User.UpdateUser(id, user); err != nil {
		return err
	}

	return nil
}
