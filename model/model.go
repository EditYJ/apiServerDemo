package model

import (
	"sync"
	"time"
)

type BaseModel struct {
	Id       uint64     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	CreateAt time.Time  `gorm:"column:createdAt" json:"-"`
	UpdateAt time.Time  `gorm:"column:updatedAt" json:"-"`
	DeleteAt *time.Time `gorm:"column:deletedAt" sql:"index" json:"-"`
}

type UserInfo struct {
	Id        uint64 `json:"id"`
	Username  string `json:"username"`
	SayHello  string `json:"sayHello"`
	Password  string `json:"password"`
	CreateAt  string `json:"createAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UserList struct {
	//互斥锁是互斥锁。互斥锁的零值是一个未锁定的互斥锁。
	//
	//互斥量在第一次使用后不能被复制。
	Lock *sync.Mutex
	IdMap map[uint64]*UserInfo
}

// 令牌
type Token struct {
	Token string `json:"token"`
}