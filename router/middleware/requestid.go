package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	)

func RequestId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 检查请求头的"X-Request-Id"，如果存在则直接使用
		requesId := ctx.Request.Header.Get("X-Request-Id")

		// 通过UUID4创建请求ID
		if requesId == ""{
			u4 := uuid.NewV4()
			requesId = u4.String()
		}

		// 暴露它 以便在应用程序中使用
		ctx.Set("X-Request-Id", requesId)

		// 设置"X-Request-Id"头
		ctx.Writer.Header().Set("X-Request-Id", requesId)
		ctx.Next()
	}
}