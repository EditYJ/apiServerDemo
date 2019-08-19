package main

import (
	"apiServerDemo/config"
	"apiServerDemo/model"
	"apiServerDemo/router"
	"apiServerDemo/router/middleware"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main() {
	pflag.Parse()

	// 初始化配置/先从命令行获取配置文件位置
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// 初始化数据库连接
	model.DB.Init()
	defer model.DB.Close()

	// 设置Gin的模式
	gin.SetMode(viper.GetString("run_mode"))

	// 创建Gin实例
	g := gin.New()

	//中间件拓展
	middlewares := []gin.HandlerFunc{
		middleware.RequestId(),
		middleware.Logging(),
	}

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
		log.Info("启动自检完成...路由部署配置成功!")
	}()

	log.Infof("开始监听网络请求，监听地址:%s", viper.GetString("addr"))
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

// 该函数用来ping服务器的http服务，确保路由是正常工作的
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count")-1; i++ {
		// 发送一个GET请求("/health")来测试服务
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		// 如果请求无错误且返回200则退出检测
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		log.Info("等待路由服务..., 1秒后重试...")
		time.Sleep(time.Second)
	}
	return errors.New("不能连接到路由服务，请检查服务器状态...")
}
