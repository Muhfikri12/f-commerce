package addresscontroller

import (
	"f-commerce/helper"
	"f-commerce/model"
	"f-commerce/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AddressController interface {
	CreateAddress(c *gin.Context)
}

type addressService struct {
	service *service.AllService
	log     *zap.Logger
}

func NewAddressController(service *service.AllService, log *zap.Logger) AddressController {
	return &addressService{service, log}
}

func (ac *addressService) CreateAddress(c *gin.Context) {

	addr := model.Address{}
	token := c.GetHeader("Authorization")

	if err := c.ShouldBindJSON(&addr); err != nil {
		ac.log.Error("Invalid payload request: " + err.Error())
		helper.Responses(c, http.StatusBadRequest, "Invalid payload request: "+err.Error(), nil)
		return
	}

	if err := ac.service.Addr.CreateAddress(token, &addr); err != nil {
		ac.log.Error("Failed to create address: " + err.Error())
		helper.Responses(c, http.StatusInternalServerError, "Failed to create address", nil)
		return
	}

	helper.Responses(c, http.StatusCreated, "Successfully Created Address", nil)

}
