package model

import (
	"apiServerDemo/pkg/auth"
	"apiServerDemo/pkg/constvar"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type UserModel struct {
	BaseModel
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

// 实现返回表名接口[TableName]
func (u *UserModel) TableName() string {
	return "tb_users"
}

// 创建新用户
// [Create]向数据库插入新记录
func (u *UserModel) Create() error {
	return DB.Self.Create(&u).Error
}

// 删除用户
// 根据ID删除数据库中的符合条件的用户
func DeleteUser(id uint64) error {
	user := UserModel{}
	user.BaseModel.Id = id
	return DB.Self.Delete(&user).Error
}

// 更新数据
// [Save]如果提供的数据无主键则会向数据库中插入这条数据
func (u *UserModel) Update() error {
	return DB.Self.Save(u).Error
}

// 根据用户名[username]查询数据
// 返回查询到的第一条数据
func GetUser(username string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Where("username = ?", username).First(&u)
	return u, d.Error
}

// 利用用户名模糊查询相关用户
// 返回([limit]限制数量)对应的用户列表[users]，所有相关用户的数量[count]
func ListUser(username string, offset, limit int) ([]*UserModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	users := make([]*UserModel, 0)

	var count uint64

	// username like %username%
	where := fmt.Sprintf("username like '%%%s%%'", username)

	// 得到数据的数量
	if err := DB.Self.Model(&UserModel{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}
	// 得到数据列表
	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil
}

// 与纯文本密码进行比较。如果与加密的相同，则返回true(在“User”结构中)。
func (u *UserModel) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return
}

// 加密
func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

// 验证字段
func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
