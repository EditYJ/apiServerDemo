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
	"github.com/lexkong/log"
	"github.com/spf13/viper"
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

	// 初始化日志模块
	c.initLog()

	// 监控配置文件变化并热加载程序
	c.watchConfig()

	return nil
}

// 初始化配置读取模块
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

// 初始化日志模块
// writers：输出位置，有两个可选项 —— file 和 stdout。选择 file 会将日志记录到 logger_file 指定的日志文件中，选择 stdout 会将日志输出到标准输出，当然也可以两者同时选择
// logger_level：日志级别，DEBUG、INFO、WARN、ERROR、FATAL
// logger_file：日志文件
// log_format_text：日志的输出格式，JSON 或者 plaintext，true 会输出成非 JSON 格式，false 会输出成 JSON 格式
// rollingPolicy：rotate 依据，可选的有 daily 和 size。如果选 daily 则根据天进行转存，如果是 size 则根据大小进行转存
// log_rotate_date：rotate 转存时间，配 合rollingPolicy: daily 使用
// log_rotate_size：rotate 转存大小，配合 rollingPolicy: size 使用
// log_backup_count：当日志文件达到转存标准时，log 系统会将该日志文件进行压缩备份，这里指定了备份文件的最大个数
func (c *Config) initLog() {
	passLagerCfg := log.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("log.logger_level"),
		LoggerFile:     viper.GetString("log.logger_file"),
		LogFormatText:  viper.GetBool("log.log_format_text"),
		RollingPolicy:  viper.GetString("log.rolling_policy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}

	log.InitWithConfig(&passLagerCfg)
}

// 监控配置文件的变化并热加载程序
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("检测到配置文件变化：%s", e.Name)
	})
}
