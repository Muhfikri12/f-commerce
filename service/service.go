package service

import (
	"f-commerce/helper"
	"f-commerce/repository"
	authservice "f-commerce/service/auth_service"
	userservice "f-commerce/service/user_service"

	"go.uber.org/zap"
)

type AllService struct {
	Auth authservice.AuthService
	User userservice.UserService
}

func NewAllService(repo *repository.Repository, log *zap.Logger, jwt *helper.Jwt) *AllService {
	return &AllService{
		Auth: authservice.NewAuthService(repo, log, jwt),
		User: userservice.NewUserService(repo, log, jwt),
	}
}
