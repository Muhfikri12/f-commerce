package authcontroller

import (
	"f-commerce/database"
	"f-commerce/helper"
	"f-commerce/model"
	"f-commerce/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mailersend/mailersend-go"
	"go.uber.org/zap"
)

type AuthController interface {
	VerificationEmail(c *gin.Context)
	Login(c *gin.Context)
	AskNewOTP(c *gin.Context)
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

	id := strconv.Itoa(user.Id)

	if err := ac.redis.SaveToken(user.Email+":"+id, token); err != nil {
		ac.log.Error("failed save token to redis : ", zap.Error(err))
		helper.Responses(c, http.StatusInternalServerError, "failed save token to redis : "+err.Error(), nil)
		return
	}

	msg := map[string]string{
		"token": token,
	}

	ac.log.Info("Login successfully")
	helper.Responses(c, http.StatusOK, "Login successfully", msg)
}

func (ac *authController) VerificationEmail(c *gin.Context) {

	verify := model.VerificationEmail{}

	if err := c.ShouldBindJSON(&verify); err != nil {
		ac.log.Error("invalid payload request: " + err.Error())
		helper.Responses(c, http.StatusBadRequest, "invalid payload request: "+err.Error(), nil)
		return
	}

	otp, err := ac.redis.Get(verify.Email)
	if err != nil {
		ac.log.Error("failed to get email from redis: " + err.Error())
		helper.Responses(c, http.StatusInternalServerError, "failed to get email from redis: "+err.Error(), nil)
		return
	}

	if otp != verify.Otp {
		ac.log.Error("otp invalid or expired")
		helper.Responses(c, http.StatusInternalServerError, "otp invalid or expired", nil)
		return
	}

	if err := ac.redis.Delete(verify.Email); err != nil {
		ac.log.Error("failed to delete otp: " + err.Error())
		helper.Responses(c, http.StatusInternalServerError, "failed to delete otp: "+err.Error(), nil)
		return
	}

	if err := ac.service.Auth.VerificationEmail(&verify); err != nil {
		ac.log.Error(err.Error())
		helper.Responses(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	ac.log.Info("verification email successfully")
	helper.Responses(c, http.StatusOK, "verification email successfully", nil)
}

func (ac *authController) AskNewOTP(c *gin.Context) {

	email := c.Query("email")

	otp := helper.GenerateOTP()

	if err := ac.service.Auth.AskNewOTP(email); err != nil {
		ac.log.Error(err.Error())
		helper.Responses(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	toSeender := []mailersend.Recipient{
		{
			Email: email,
		},
	}

	if err := helper.SendOTPEmail(toSeender, otp); err != nil {
		ac.log.Error("failed sent otp : ", zap.Error(err))
		helper.Responses(c, http.StatusBadRequest, "failed sent otp : "+err.Error(), nil)
		return
	}

	if err := ac.redis.SetRedis(email, otp, 5*60); err != nil {
		ac.log.Error("failed set otp on redis : ", zap.Error(err))
		helper.Responses(c, http.StatusBadRequest, "failed set otp on redis : "+err.Error(), nil)
	}

	ac.log.Info("Otp sent successfully")
	helper.Responses(c, http.StatusOK, "Otp sent successfully", nil)
}
