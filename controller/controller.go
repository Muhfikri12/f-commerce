package controller

import (
	"f-commerce/config"
	addresscontroller "f-commerce/controller/address_controller"
	authcontroller "f-commerce/controller/auth_controller"
	categorycontroller "f-commerce/controller/category_controller"
	usercontroller "f-commerce/controller/user_controller"
	"f-commerce/database"
	"f-commerce/service"

	"go.uber.org/zap"
)

type AllController struct {
	Auth authcontroller.AuthController
	User usercontroller.UserController
	Addr addresscontroller.AddressController
	Cat  categorycontroller.CategoryController
}

func NewAllController(service *service.AllService, log *zap.Logger, redis *database.Cache, cfg *config.Config) *AllController {
	return &AllController{
		Auth: authcontroller.NewAuthController(service, log, redis),
		User: usercontroller.NewUserController(service, log, redis, cfg),
		Addr: addresscontroller.NewAddressController(service, log),
		Cat:  categorycontroller.NewCategoryController(service, log),
	}
}
