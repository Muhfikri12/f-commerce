package usercontroller

import (
	"f-commerce/helper"
	"f-commerce/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *userController) UpdateCustomer(c *gin.Context) {

	cust := model.CustomerData{}

	token := c.GetHeader("Authorization")

	if err := c.ShouldBindJSON(&cust); err != nil {
		uc.log.Error("Invalid payload request: " + err.Error())
		helper.Responses(c, http.StatusBadRequest, "Invalid payload request: "+err.Error(), nil)
		return
	}

	if valid, msg := helper.ValidatePassword(cust.User.Password); !valid {
		uc.log.Error("Validation error: " + msg.Error())
		helper.Responses(c, http.StatusBadRequest, msg.Error(), nil)
		return
	}

	if err := uc.service.User.UpdateCustomer(token, &cust); err != nil {
		uc.log.Error("Failed to update customer: " + err.Error())
		helper.Responses(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

}

func (uc *userController) UpdateProfile(c *gin.Context) {

	token := c.GetHeader("Authorization")

	filePath, err := helper.UploadImage(c, "image")
	if err != nil {
		uc.log.Error("Failed to upload image: " + err.Error())
		helper.Responses(c, http.StatusBadRequest, "Failed to upload image: "+err.Error(), nil)
		return
	}

	err = uc.service.User.UpdateProfile(token, filePath)
	if err != nil {
		uc.log.Error("Failed parsing payload JWT: " + err.Error())
		helper.Responses(c, http.StatusBadRequest, "Failed parsing payload JWT: "+err.Error(), nil)
		return
	}

}
