package service

import (
	"finance/helper"
	"finance/repository"
	authservice "finance/service/auth_service"

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
