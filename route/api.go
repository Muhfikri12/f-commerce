package route

import (
	"f-commerce/infra"

	"github.com/gin-gonic/gin"
)

func NewRoutes(ctx *infra.IntegrationContext) *gin.Engine {

	r := gin.Default()

	r.POST("/login", ctx.Ctl.Auth.Login)
	r.POST("/register", ctx.Ctl.User.RegisterUser)

	return r

}
