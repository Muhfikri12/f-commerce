package authcontroller

import (
	"f-commerce/database"
	"f-commerce/helper"
	"f-commerce/model"
	"f-commerce/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mailersend/mailersend-go"
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

	otp := helper.GenerateOTP()

	err := c.ShouldBindJSON(&login)
	if err != nil {
		ac.log.Error("error payload request: " + err.Error())
		helper.Responses(c, http.StatusBadRequest, "error payload request: "+err.Error(), nil)
		return
	}

	user, err := ac.service.Auth.Login(&login)
	if err != nil {
		ac.log.Error(err.Error())
		helper.Responses(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	toSeender := []mailersend.Recipient{
		{
			Email: user.Email,
		},
	}

	if err := helper.SendOTPEmail(toSeender, otp); err != nil {
		ac.log.Error("failed sent otp : ", zap.Error(err))
		helper.Responses(c, http.StatusBadRequest, "failed sent otp : "+err.Error(), nil)
		return
	}

	if err := ac.redis.SetRedis(user.Email, otp, 5*60); err != nil {
		ac.log.Error("failed save otp to redis : ", zap.Error(err))
		helper.Responses(c, http.StatusInternalServerError, "failed save otp to redis : "+err.Error(), nil)
		return
	}

	ac.log.Error("successfully sent otp")

	helper.Responses(c, http.StatusOK, "successfully sent otp", nil)
}
