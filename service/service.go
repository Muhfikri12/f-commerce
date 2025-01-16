package service

import (
	"f-commerce/helper"
	"f-commerce/repository"
	authservice "f-commerce/service/auth_service"

	"go.uber.org/zap"
)

type AllService struct {
	Auth authservice.AuthService
}

func NewAllService(repo *repository.Repository, log *zap.Logger, jwt *helper.Jwt) *AllService {
	return &AllService{
		Auth: authservice.NewAuthService(repo, log, jwt),
	}
}
