package controller

import (
	authcontroller "f-commerce/controller/auth_controller"
	usercontroller "f-commerce/controller/user_controller"
	"f-commerce/database"
	"f-commerce/service"

	"go.uber.org/zap"
)

type AllController struct {
	Auth authcontroller.AuthController
	User usercontroller.UserController
}

func NewAllController(service *service.AllService, log *zap.Logger, redis *database.Cache) *AllController {
	return &AllController{
		Auth: authcontroller.NewAuthController(service, log, redis),
		User: usercontroller.NewUserController(service, log, redis),
	}
}
