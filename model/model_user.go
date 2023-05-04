package model

import (
	"errors"
	"github.com/CreFire/rain/dal"
	"time"
)

// User 用户模型
type User struct {
	Id       uint64     `xorm:"'id' AUTO_INCREMENT pk" json:"id"`                 // 用户 ID，主键自增
	Name     string     `xorm:"'name' varchar(50)" json:"name"`                   // 用户姓名
	Age      int32      `xorm:"'age' int(3)" json:"age"`                          // 用户年龄
	Birthday *time.Time `xorm:"'birthday' date" json:"birthday,omitempty"`        // 用户出生日期（可选）
	Email    string     `xorm:"'email' varchar(255) unique_index" json:"email"`   // 用户邮箱（唯一）
	PassWord string     `xorm:"'password' varchar(25)" json:"password,omitempty"` // 用户密码（可选）
	Role     int32      `xorm:"'position' int" json:"position"`                   // 用户职位
	Nickname string     `xorm:"'nickname' varchar(50)" json:"nickname"`           // 用户昵称
	iPhone   string     `xorm:"varchar(20)"`                                      // 用户手机号码
}

func (u *User) TableName() string {
	return "user" // 用户表
}

// CreateUser 创建用户
func CreateUser(user *User) error {
	session := dal.GetDb().NewSession()
	defer session.Close()

	affected, err := session.Insert(user)
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("failed to create user")
	}
	return nil
}

// UpdateUser 更新用户
func UpdateUser(id uint, user *User) error {
	session := dal.GetDb().NewSession()
	defer session.Close()

	affected, err := session.ID(id).Update(user)
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("failed to update user")
	}
	return nil
}

// DeleteUser 删除用户
func DeleteUser(id uint) error {
	session := dal.GetDb().NewSession()
	defer session.Close()

	affected, err := session.ID(id).Delete(&User{})
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("failed to delete user")
	}
	return nil
}

// GetUser 获取用户
func GetUser(id uint64) (*User, error) {
	session := dal.GetDb().NewSession()
	defer session.Close()

	user := &User{Id: id}
	has, err := session.Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("user not found")
	}
	return user, nil
}
