package util

import (
	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid" // 独特的非顺序短ID的发生器
)

// 产生不重复的ID
func GetShortId() (string, error) {
	return shortid.Generate()
}

// 得到请求附带的ID
func GetReqID(c *gin.Context) string {
	v, ok := c.Get("X-Request-Id")
	if !ok {
		return ""
	}
	if requestId, ok := v.(string); ok {
		return requestId
	}
	return ""
}
