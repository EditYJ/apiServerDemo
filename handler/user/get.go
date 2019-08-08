package user

import (
	. "apiServerDemo/handler"
	"apiServerDemo/model"
	"apiServerDemo/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

func Get(c *gin.Context) {
	log.Info("查询单个用户函数[Get()]被调用.", lager.Data{"username": c.Param("username")})
	username := c.Param("username")

	// 从数据库查询对应用户名
	user, err := model.GetUser(username)
	if err != nil{
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	SendResponse(c, nil, user)
}