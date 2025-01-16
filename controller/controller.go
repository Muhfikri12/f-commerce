package controller

import (
	authcontroller "f-commerce/controller/auth_controller"
	"f-commerce/database"
	"f-commerce/service"

	"go.uber.org/zap"
)

type AllController struct {
	Auth authcontroller.AuthController
}

func NewAllController(service *service.AllService, log *zap.Logger, redis *database.Cache) *AllController {
	return &AllController{
		Auth: authcontroller.NewAuthController(service, log, redis),
	}
}
