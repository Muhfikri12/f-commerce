package usercontroller

import (
	"f-commerce/config"
	"f-commerce/database"
	"f-commerce/helper"
	"f-commerce/model"
	"f-commerce/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mailersend/mailersend-go"
	"go.uber.org/zap"
)

type UserController interface {
	RegisterUser(c *gin.Context)
	UpdateCustomer(c *gin.Context)
	UpdateRole(c *gin.Context)
	UpdateProfile(c *gin.Context)
	UpdateAdmin(c *gin.Context)
}

type userController struct {
	service *service.AllService
	log     *zap.Logger
	rdb     *database.Cache
	cfg     *config.Config
}

func NewUserController(service *service.AllService, log *zap.Logger, rdb *database.Cache, cfg *config.Config) UserController {
	return &userController{service: service, log: log, rdb: rdb, cfg: cfg}
}

func (uc *userController) RegisterUser(c *gin.Context) {
	user := model.Register{}

	otp := helper.GenerateOTP()

	if err := c.ShouldBindJSON(&user); err != nil {
		uc.log.Error("Invalid payload request: " + err.Error())
		helper.Responses(c, http.StatusBadRequest, "Invalid payload request: "+err.Error(), nil)
		return
	}

	if valid, msg := helper.ValidatePassword(user.Password); !valid {
		uc.log.Error("Validation error: " + msg.Error())
		helper.Responses(c, http.StatusBadRequest, msg.Error(), nil)
		return
	}

	if err := uc.service.User.RegisterUser(&user); err != nil {
		uc.log.Error("Error: " + err.Error())
		helper.Responses(c, http.StatusInternalServerError, "Error: "+err.Error(), nil)
		return
	}

	toSeender := []mailersend.Recipient{
		{
			Email: user.Email,
		},
	}

	if err := helper.SendOTPEmail(toSeender, otp); err != nil {
		uc.log.Error("failed sent otp : ", zap.Error(err))
		helper.Responses(c, http.StatusBadRequest, "failed sent otp : "+err.Error(), nil)
		return
	}

	if err := uc.rdb.SetRedis(user.Email, otp, 5*60); err != nil {
		uc.log.Error("failed set otp on redis : ", zap.Error(err))
		helper.Responses(c, http.StatusBadRequest, "failed set otp on redis : "+err.Error(), nil)
		return
	}

	uc.log.Info("Registration successfully")

	helper.Responses(c, http.StatusCreated, "Registration successfully, Please check email for verification otp", nil)
}
