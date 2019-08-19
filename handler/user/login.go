package user

import (
	. "apiServerDemo/handler"
	"apiServerDemo/model"
	"apiServerDemo/pkg/auth"
	"apiServerDemo/pkg/errno"
	"apiServerDemo/pkg/token"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	// 为user结构体绑定数据
	var u model.UserModel
	if err := ctx.Bind(&u); err != nil {
		SendResponse(ctx, errno.ErrBind, nil)
		return
	}

	// 通过登陆用户名得到用户的信息
	d, err := model.GetUser(u.Username)
	if err != nil {
		SendResponse(ctx, errno.ErrUserNotFound, nil)
		return
	}

	// 比较用户密码是否正确
	if err := auth.Compare(d.Password, u.Password); err != nil {
		SendResponse(ctx, errno.ErrPasswordIncorrect, nil)
		return
	}

	//签名jwt
	t, err := token.Sign(ctx, token.Context{
		ID:       d.Id,
		Username: d.Username,
	}, "")
	if err != nil {
		SendResponse(ctx, errno.ErrToken, nil)
		return
	}
	SendResponse(ctx, nil, model.Token{Token: t})
}
