package route

import (
	"f-commerce/infra"

	"github.com/gin-gonic/gin"
)

func NewRoutes(ctx *infra.IntegrationContext) *gin.Engine {

	r := gin.Default()

	r.POST("/login", ctx.Ctl.Auth.Login)
	r.POST("/verify-email", ctx.Ctl.Auth.VerificationEmail)
	r.POST("/register", ctx.Ctl.User.RegisterUser)
	r.GET("/new-otp", ctx.Ctl.Auth.AskNewOTP)

	user := r.Group("/users", ctx.Middleware.Middleware())
	{
		user.PUT("/", ctx.Ctl.User.UpdateUser)
		user.PUT("/admin", ctx.Ctl.User.UpdateAdmin)
		user.PUT("/profile", ctx.Ctl.User.UpdateProfile)
		user.PUT("/role", ctx.Ctl.User.UpdateRole)
		user.PUT("/user", ctx.Ctl.User.UpdateCustomer)
	}

	addr := r.Group("/address", ctx.Middleware.Middleware())
	{
		addr.POST("/", ctx.Ctl.Addr.CreateAddress)
		addr.GET("/", ctx.Ctl.Addr.FindAddressByUserID)
		addr.PUT("/:id", ctx.Ctl.Addr.UpdateAddress)
		addr.GET("/:id", ctx.Ctl.Addr.FindAddressByID)
	}

	cat := r.Group("/category", ctx.Middleware.Middleware())
	{
		cat.POST("/", ctx.Ctl.Cat.CreateCategory)
	}

	return r

}
