package authcontroller

import (
	"f-commerce/database"
	"f-commerce/helper"
	"f-commerce/model"
	"f-commerce/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthController interface {
	Login(c *gin.Context)
}
type authController struct {
	service *service.AllService
	log     *zap.Logger
	redis   *database.Cache
}

func NewAuthController(service *service.AllService, log *zap.Logger, redis *database.Cache) AuthController {
	return &authController{service, log, redis}
}

func (ac *authController) Login(c *gin.Context) {

	login := model.Login{}

	err := c.ShouldBindJSON(&login)
	if err != nil {
		ac.log.Error("error payload request: " + err.Error())
		helper.Responses(c, http.StatusBadRequest, "error payload request: "+err.Error(), nil)
		return
	}

	user, token, err := ac.service.Auth.Login(&login)
	if err != nil {
		ac.log.Error(err.Error())
		helper.Responses(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := ac.redis.SaveToken(user.Email, token); err != nil {
		ac.log.Error("failed save token to redis : ", zap.Error(err))
		helper.Responses(c, http.StatusInternalServerError, "failed save token to redis : "+err.Error(), nil)
		return
	}

	msg := map[string]string{
		"token": token,
	}

	ac.log.Info("successfully sent Token")
	helper.Responses(c, http.StatusOK, "successfully sent otp", msg)
}
