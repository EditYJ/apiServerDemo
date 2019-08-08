package user

import (
	. "apiServerDemo/handler"
	"apiServerDemo/model"
	"apiServerDemo/pkg/errno"
	"apiServerDemo/util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"strconv"
)

func Update(c *gin.Context) {
	log.Info("更新用户函数[Update()]被调用.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	userId, _ := strconv.Atoi(c.Param("id"))

	// 绑定用户数据
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	// 更新记录依赖与用户id
	u.Id = uint64(userId)

	// 验证数据
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// 加密用户密码
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	// 保存更改
	if err := u.Update(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)

}
