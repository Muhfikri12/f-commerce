package categorycontroller

import (
	"f-commerce/helper"
	"f-commerce/model"
	"f-commerce/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CategoryController interface {
	CreateCategory(c *gin.Context)
	ReadCategories(c *gin.Context)
}

type categoryController struct {
	service *service.AllService
	log     *zap.Logger
}

func NewCategoryController(service *service.AllService, log *zap.Logger) CategoryController {
	return &categoryController{service, log}
}

func (cc *categoryController) CreateCategory(c *gin.Context) {

	cat := model.Category{}

	if err := c.ShouldBindJSON(&cat); err != nil {
		cc.log.Error("Invalid payload request: " + err.Error())
		helper.Responses(c, http.StatusBadRequest, "Invalid payload request: "+err.Error(), nil)
		return
	}

	if err := cc.service.Cat.CreateCategory(&cat); err != nil {
		cc.log.Error("Failed created category: " + err.Error())
		helper.Responses(c, http.StatusInternalServerError, "Failed created category", nil)
		return
	}

	helper.Responses(c, http.StatusCreated, "Successfully Created category", nil)

}

func (cc *categoryController) ReadCategories(c *gin.Context) {

	cat, err := cc.service.Cat.ReadCategories()
	if err != nil {
		cc.log.Error("Failed Show All Categories: " + err.Error())
		helper.Responses(c, http.StatusInternalServerError, "Failed Show All Categories", nil)
		return
	}

	helper.Responses(c, http.StatusOK, "Successfully Retrieved Categories", cat)
}
