package authcontroller

import (
	"f-commerce/database"
	"f-commerce/helper"
	"f-commerce/model"
	"f-commerce/service"
	"net/http"
	"strconv"

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
		helper.Responses(c, http.StatusInternalServerError, "error payload request: "+err.Error(), nil)
		return
	}

	token, id, err := ac.service.Auth.Login(&login)
	if err != nil {
		ac.log.Error(err.Error())
		helper.Responses(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	idStr := strconv.Itoa(*id)

	if err := ac.redis.SaveToken(idStr, token); err != nil {
		ac.log.Error("failed save token to redis, ", zap.Error(err))
		helper.Responses(c, http.StatusBadRequest, "failed save token to redis, "+err.Error(), nil)
		return
	}

	ac.log.Error("login successfully")

	msg := map[string]string{
		"token": token,
	}

	helper.Responses(c, http.StatusOK, "login successfully", msg)
}
