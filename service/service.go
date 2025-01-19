package service

import (
	"f-commerce/helper"
	"f-commerce/repository"
	addressservice "f-commerce/service/address_service"
	authservice "f-commerce/service/auth_service"
	categoryservice "f-commerce/service/category_service"
	userservice "f-commerce/service/user_service"

	"go.uber.org/zap"
)

type AllService struct {
	Auth authservice.AuthService
	User userservice.UserService
	Addr addressservice.AddressService
	Cat  categoryservice.CategoryService
}

func NewAllService(repo *repository.Repository, log *zap.Logger, jwt *helper.Jwt) *AllService {
	return &AllService{
		Auth: authservice.NewAuthService(repo, log, jwt),
		User: userservice.NewUserService(repo, log, jwt),
		Addr: addressservice.NewAddressService(repo, log, jwt),
		Cat:  categoryservice.NewCategoryService(repo, log),
	}
}
