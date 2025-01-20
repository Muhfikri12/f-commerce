package usercontroller

import (
	"f-commerce/helper"
	"f-commerce/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *userController) UpdateCustomer(c *gin.Context) {

	cust := model.Customer{}

	token := c.GetHeader("Authorization")

	if err := c.ShouldBindJSON(&cust); err != nil {
		uc.log.Error("Invalid payload request: " + err.Error())
		helper.Responses(c, http.StatusBadRequest, "Invalid payload request: "+err.Error(), nil)
		return
	}

	if err := uc.service.User.UpdateCustomer(token, &cust); err != nil {
		uc.log.Error("Failed to update customer: " + err.Error())
		helper.Responses(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.Responses(c, http.StatusOK, "successfully update customer", nil)

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
		uc.log.Error("Failed to update profile: " + err.Error())
		helper.Responses(c, http.StatusInternalServerError, "Failed to update profile: "+err.Error(), nil)
		return
	}

	helper.Responses(c, http.StatusOK, "successfully update profile", nil)

}

func (uc *userController) UpdateRole(c *gin.Context) {

	token := c.GetHeader("Authorization")

	err := uc.service.User.UpdateRole(token)
	if err != nil {
		uc.log.Error("Failed to update role: " + err.Error())
		helper.Responses(c, http.StatusInternalServerError, "Failed to update role: "+err.Error(), nil)
		return
	}

	helper.Responses(c, http.StatusOK, "successfully update role", nil)

}

func (uc *userController) NonactiveAccount(c *gin.Context) {

	token := c.GetHeader("Authorization")

	err := uc.service.User.NonactiveAccount(token)
	if err != nil {
		uc.log.Error("Failed to nonactive role: " + err.Error())
		helper.Responses(c, http.StatusInternalServerError, "Failed to nonactive role: "+err.Error(), nil)
		return
	}

	helper.Responses(c, http.StatusOK, "successfully nonactive role", nil)

}
