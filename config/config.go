// Viper 配置读取顺序：
//
//viper.Set() 所设置的值
//命令行 flag
//环境变量
//配置文件
//配置中心：etcd/consul
//默认值
package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type Config struct {
	Name string
}

func Init(cfg string) error {
	c := Config{Name: cfg}

	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}

	// 监控配置文件变化并热加载程序
	c.watchConfig()

	return nil
}

// 初始化配置
func (c *Config) initConfig() error {
	// 如果指定了配置文件则解析指定的配置文件
	if c.Name != "" {
		viper.SetConfigFile(c.Name)
	} else { // 否则解析默认的配置文件/默认配置文件目录为conf下的config文件
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	// 设置文件格式为YAML
	viper.SetConfigType("yaml")
	// 读取匹配的环境变量
	viper.AutomaticEnv()
	// 读取环境变量的前缀为[API_SERVER]
	viper.SetEnvPrefix("API_SERVER")

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

// 监控配置文件的变化并热加载程序
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("检测到配置文件变化：%s", e.Name)
	})
}
