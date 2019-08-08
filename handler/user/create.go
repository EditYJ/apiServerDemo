package user

import (
	. "apiServerDemo/handler"
	"apiServerDemo/model"
	"apiServerDemo/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"time"
)

// 创建用户
func Create(c *gin.Context) {
	log.Info("创建用户函数[Create()]被调用.")
	var r CreateRequest

	// 检查 Content-Type 类型，将消息体作为指定的格式解析到 Go struct 变量中。
	// 此处采用的媒体类型是 JSON，所以 Bind() 是按 JSON 格式解析的。
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		//c.JSON(http.StatusOK, gin.H{"error": errno.ErrBind})
		return
	}

	u := model.UserModel{
		BaseModel: model.BaseModel{
			Id:       0,
			CreateAt: time.Now(),
			UpdateAt: time.Now(),
			DeleteAt: nil,
		},
		Username:  r.Username,
		Password:  r.Password,
	}

	// 验证数据
	if err := u.Validate();err != nil{
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// 加密用户密码
	if err:= u.Encrypt(); err!= nil{
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	// 将用户信息插入表
	if err := u.Create();err != nil{
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	// 构建返回信息，返回用户名信息
	rsp := CreateResponse{Username:r.Username}
	SendResponse(c, nil, rsp)
}
