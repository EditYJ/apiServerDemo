// 路由处理
package router

import (
	"apiServerDemo/handler/sd"
	"apiServerDemo/handler/user"
	"apiServerDemo/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// 中间件
	// Recovery返回一个中间件，该中间件从任何错误[panics]中恢复，如果有500，则写入500。
	g.Use(gin.Recovery())

	// 附加请求头
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	// 处理404
	g.NoRoute(func(ctx *gin.Context) {
		ctx.String(http.StatusNotFound, "404，这是一个不正确的API路由地址!")
	})

	// 硬件状况检查路由处理
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	// 用户相关
	u := g.Group("/v1/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.POST("", user.Create)				// 创建账号
		u.DELETE("/:id", user.Delete)	// 删除账号
		u.PUT("/:id", user.Update)		// 更新账号
		u.GET("", user.List)					// 账号列表
		u.GET("/:username", user.Get)	// 获取指定账号的详细信息
		u.POST("/login", user.Login)	// 用户登录
	}
	return g
}