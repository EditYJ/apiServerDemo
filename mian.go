package main

import (
	"apiServerDemo/router"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	g := gin.New()
	middlewares := []gin.HandlerFunc{}

	// Routes
	router.Load(
		// 核心
		g,
		// 中间件
		middlewares...,
	)

	// 启动另一个协程检查服务器是否启动正常
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("路由处理无响应，或者处理花的时间过长，请检查问题...", err)
		}
		log.Print("路由部署配置成功")
	}()

	log.Printf("开始监听网络请求，监听地址:%s", ":8078")
	log.Print(http.ListenAndServe(":8078", g).Error())
}

// 该函数用来ping服务器的http服务，确保路由是正常工作的
func pingServer() error {
	for i := 0; i < 2; i++ {
		// 发送一个GET请求("/health")来测试服务
		resp, err := http.Get("http://127.0.0.1:8078" + "/sd/health")
		// 如果请求无错误且返回200则退出检测
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		log.Print("等待路由服务..., 1秒后重试...")
		time.Sleep(time.Second)
	}
	return errors.New("不能连接到路由服务，请检查服务器状态...")
}
