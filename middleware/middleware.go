package middleware

import (
	"f-commerce/helper"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Middleware struct {
	log *zap.Logger
	jwt *helper.Jwt
}

func NewMiddleware(log *zap.Logger, jwt *helper.Jwt) *Middleware {
	return &Middleware{log: log, jwt: jwt}
}

func (m *Middleware) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenHeader := ctx.GetHeader("Authorization")

		if tokenHeader == "" {
			helper.Responses(ctx, http.StatusUnauthorized, "Authorization token is required", nil)
			ctx.Abort()
			return
		}

		_, err := m.jwt.ParsingPayload(tokenHeader)
		if err != nil {
			helper.Responses(ctx, http.StatusUnauthorized, "Invalid token or token expired", nil)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
