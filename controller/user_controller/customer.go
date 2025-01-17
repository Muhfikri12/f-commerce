package usercontroller

import (
	"f-commerce/helper"
	"f-commerce/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userController) UpdateCustomer(c *gin.Context) {

	cust := model.CustomerData{}

	token := c.GetHeader("Authorization")
	fmt.Println(token)

	jwt := helper.NewJwt(uc.cfg, uc.log)

	idStr, err := jwt.ParsingPayload(token)
	if err != nil {
		uc.log.Error("Failed parsing payload JWT: " + err.Error())
		helper.Responses(c, http.StatusBadRequest, "Failed parsing payload JWT: "+err.Error(), nil)
		return
	}

	id, _ := strconv.Atoi(idStr.(string))

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

	if err := uc.service.User.UpdateCustomer(id, &cust); err != nil {
		uc.log.Error("Failed to update customer: "+err.Error(), zap.Int("id", id))
		helper.Responses(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

}
