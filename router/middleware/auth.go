package middleware

import (
	"apiServerDemo/handler"
	"apiServerDemo/pkg/errno"
	"apiServerDemo/pkg/token"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 解析jwt
		if _, err := token.ParseRequest(ctx); err != nil {
			handler.SendResponse(ctx, errno.ErrTokenInvalid, nil)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
