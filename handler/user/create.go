package user

import (
	. "apiServerDemo/handler"
	"apiServerDemo/pkg/errno"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

func Create(c *gin.Context) {
	var r CreateRequest

	// 检查 Content-Type 类型，将消息体作为指定的格式解析到 Go struct 变量中。
	// 此处采用的媒体类型是 JSON，所以 Bind() 是按 JSON 格式解析的。
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		//c.JSON(http.StatusOK, gin.H{"error": errno.ErrBind})
		return
	}

	// [Param(key)]返回 URL的参数值
	// 例：
	//  router.GET("/user/:id", func(c *gin.Context) {     // a GET request to /user/john
	//     id := c.Param("id") // id == "john"
	// })
	admin2 := c.Param("username")
	log.Infof("URL username: %s", admin2)

	// [Query(key)]返回URL中的地址参数
	// 例：
	//  GET /path?id=1234&name=Manu&value=
	//  c.Query("id") == "1234"
	//  c.Query("name") == "Manu"
	//  c.Query("value") == ""
	//  c.Query("wtf") == ""
	desc := c.Query("desc")
	log.Infof("URL key param desc: %s", desc)

	// [GetHeader(key)]获取 HTTP 头
	//
	// 拓展：DefaultQuery()：类似 Query()，但是如果 key 不存在，会返回默认值。
	// 例如：
	// GET /?name=Manu&lastname=
	// c.DefaultQuery("name", "unknown") == "Manu"
	// c.DefaultQuery("id", "none") == "none"
	// c.DefaultQuery("lastname", "none") == ""
	contentType := c.GetHeader("Content-Type")
	log.Infof("Header Content-Type: %s", contentType)

	log.Debugf("username is: [%s], password is: [%s]", r.Username, r.Password)

	if r.Username == "" {
		err := errno.New(errno.ErrUserNotFound, fmt.Errorf("Username can not be NULL")).Add("用户名不能为空！")
		SendResponse(c, err, nil)
		return
	}

	if r.Password == "" {
		err := fmt.Errorf("password is empty")
		SendResponse(c, err, nil)
	}

	rsp:=CreateResponse{Username:r.Username}

	// 返回用户名信息
	SendResponse(c, nil, rsp)
}
