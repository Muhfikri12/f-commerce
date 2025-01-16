package userservice

import (
	"f-commerce/model"
	"f-commerce/repository"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(user *model.User) error
}

type userService struct {
	Repo *repository.Repository
	Log  *zap.Logger
}

func NewUserService(Repo *repository.Repository, Log *zap.Logger) UserService {
	return &userService{Repo, Log}
}

func (us *userService) RegisterUser(user *model.User) error {

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(password)

	if err := us.Repo.User.RegisterUser(user); err != nil {
		return err
	}

	return nil
}
