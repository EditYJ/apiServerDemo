package model

import (
	"fmt"
	"github.com/lexkong/log"
	"github.com/spf13/viper"

	// MySQL driver.
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Datebase struct {
	Self   *gorm.DB
	Docker *gorm.DB
}

var DB *Datebase

// 打开数据库连接
func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=%s",
		username,
		password,
		addr,
		name,
		// "Asia/Shanghai")
		"Local")
	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Errorf(err, "数据库连接失败，数据库名称为: %s", name)
	}

	// 设置数据库连接
	setupDB(db)

	return db
}

func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	// 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	//db.DB().SetMaxOpenConns(20000)
	// 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	db.DB().SetMaxOpenConns(0)
}

// 命令行界面使用
func InitSelfDB() *gorm.DB {
	return openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"))
}

func GetSelfDB() *gorm.DB {
	return InitSelfDB()
}

func InitDockerDB() *gorm.DB {
	return openDB(viper.GetString("docker_db.username"),
		viper.GetString("docker_db.password"),
		viper.GetString("docker_db.addr"),
		viper.GetString("docker_db.name"))
}

func GetDockerDB() *gorm.DB {
	return InitDockerDB()
}

func (db *Datebase) Init() {
	DB = &Datebase{
		Self:   GetSelfDB(),
		Docker: GetDockerDB(),
	}
}

func (db *Datebase) Close() {
	_ = DB.Self.Close()
	_ = DB.Docker.Close()
}