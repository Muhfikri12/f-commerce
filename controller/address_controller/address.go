package addresscontroller

import (
	"f-commerce/helper"
	"f-commerce/model"
	"f-commerce/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AddressController interface {
	CreateAddress(c *gin.Context)
	FindAddressByUserID(c *gin.Context)
	UpdateAddress(c *gin.Context)
	FindAddressByID(c *gin.Context)
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
		ac.log.Error("Failed created address: " + err.Error())
		helper.Responses(c, http.StatusInternalServerError, "Failed created address", nil)
		return
	}

	helper.Responses(c, http.StatusCreated, "Successfully Created Address", nil)

}

func (ac *addressService) FindAddressByUserID(c *gin.Context) {

	token := c.GetHeader("Authorization")

	address, err := ac.service.Addr.FindAddressByUserID(token)
	if err != nil {
		ac.log.Error("Failed retrieved address: " + err.Error())
		helper.Responses(c, http.StatusInternalServerError, "Failed retrieved address", nil)
		return
	}

	helper.Responses(c, http.StatusOK, "Successfully Retrieved Address", address)
}

func (ac *addressService) UpdateAddress(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	addr := model.Address{}

	if err := c.ShouldBindJSON(&addr); err != nil {
		ac.log.Error("Invalid payload request: " + err.Error())
		helper.Responses(c, http.StatusBadRequest, "Invalid payload request: "+err.Error(), nil)
		return
	}

	if err := ac.service.Addr.UpdateAddress(id, &addr); err != nil {
		ac.log.Error("Failed updated address: " + err.Error())
		helper.Responses(c, http.StatusInternalServerError, "Failed updated address", nil)
		return
	}

	helper.Responses(c, http.StatusOK, "Successfully Updated Address", nil)
}

func (ac *addressService) FindAddressByID(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	address, err := ac.service.Addr.FindAddressByID(id)
	if err != nil {
		ac.log.Error("Failed retrieved address: " + err.Error())
		helper.Responses(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.Responses(c, http.StatusOK, "Successfully Retrieved Address", address)
}
