package user

import (
	"apiServerDemo/pkg/errno"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"net/http"
)

func Create(c *gin.Context) {
	var r struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var err error

	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": errno.ErrBind})
		return
	}

	log.Debugf("username is: [%s], password is: [%s]", r.Username, r.Password)

	if r.Username == "" {
		err = errno.New(errno.ErrUserNotFound, fmt.Errorf("Username can not be NULL")).Add("用户名不能为空！")
		log.Errorf(err, "发现错误==>>")
	}

	if errno.IsErrUserNotFound(err) {
		log.Debug("错误类型==>> ErrUserNotFount")
	}

	if r.Password == "" {
		err = fmt.Errorf("password is empty")
		log.Errorf(err, "发现错误==>>")
	}

	code, message := errno.DecodeErr(err)
	c.JSON(http.StatusOK, gin.H{"code": code, "message": message})
}
