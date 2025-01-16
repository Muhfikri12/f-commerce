package userservice

import (
	"f-commerce/model"
	"f-commerce/repository"

	"go.uber.org/zap"
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

	if err := us.Repo.User.CreateCustomer(user); err != nil {
		return err
	}

	return nil
}
