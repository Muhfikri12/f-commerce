package usercontroller

import (
	"f-commerce/helper"
	"f-commerce/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *userController) UpdateAdmin(c *gin.Context) {

	admin := model.Admin{}

	token := c.GetHeader("Authorization")

	if err := c.ShouldBindJSON(&admin); err != nil {
		uc.log.Error("Invalid payload request: " + err.Error())
		helper.Responses(c, http.StatusBadRequest, "Invalid payload request: "+err.Error(), nil)
		return
	}

	if err := uc.service.User.UpdateAdmin(token, &admin); err != nil {
		uc.log.Error("Failed to update admin: " + err.Error())
		helper.Responses(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.Responses(c, http.StatusOK, "successfully update admin", nil)

}
