package user

import (
	. "apiServerDemo/handler"
	"apiServerDemo/model"
	"apiServerDemo/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"strconv"
)

func Delete(c *gin.Context)  {
	log.Info("删除用户函数[Delete()]被调用.", lager.Data{"id": c.Param("id")})
	userId,_ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userId)); err!=nil{
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}