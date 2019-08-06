package handler

import (
	"apiServerDemo/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

type bodyPart struct {
	Data interface{} `json:"data"`
}

type headerPart struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Body  bodyPart   `json:"body"`
	Heads headerPart `json:"heads"`
}

// 发送消息格式封装
//
// 格式示例
//{
//	"body": {
//		"data": {}
//	},
//	"heads": {
//		"code":200,
//		"message":"success"
//	}
//}
func SendResponse(ctx *gin.Context, err error, data interface{}) {
	// 解析错误信息
	code, message := errno.DecodeErr(err)

	// 总是返回200，代表服务器返回了数据？
	ctx.JSON(http.StatusOK, Response{
		Body: bodyPart{
			Data: data,
		},
		Heads: headerPart{
			Code:    code,
			Message: message,
		},
	})
}
