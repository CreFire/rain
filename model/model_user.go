package model

import (
	"errors"
	"github.com/CreFire/rain/dal"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// User 用户模型
type User struct {
	Id          *uint64    `xorm:"'id' bigint AUTO_INCREMENT autoincr pk " json:"id"` // 用户 ID，主键自增
	Avatar      *string    `xorm:"varchar(50) 'avatar' " json:"avatar"`               // 用户头像
	Account     *string    `xorm:"varchar(50) 'Account' " json:"Account"`             // 用户账号
	IPhone      *string    `xorm:"varchar(20)" json:"IPhone"`                         // 用户手机号码
	Name        string     `xorm:"varchar(50) 'name' " json:"name"`                   // 用户姓名
	Age         *int32     `xorm:"int(3) 'age' " json:"age"`                          // 用户年龄
	Birthday    *time.Time `xorm:"date 'birthday'" json:"birthday,omitempty"`         // 用户出生日期（可选）
	Email       *string    `xorm:"'email' varchar(255) unique_index" json:"email"`    // 用户邮箱（唯一）
	PassWord    string     `xorm:"'password' varchar(25)" json:"password,omitempty"`  // 用户密码（可选）
	Role        *int32     `xorm:"'position' int" json:"position"`                    // 用户权限
	Nickname    *string    `xorm:"'nickname' varchar(50)" json:"nickname"`            // 用户昵称
	Description *string    `xorm:"bigint" json:"description"`                         // 描述
	ExpireTime  *int64     `xorm:"bigint" json:"expireTime"`                          // 导出时间
	CreateTime  *int64     `xorm:"bigint" json:"createTime"`                          // 创建时间
	UpdateTime  *int64     `xorm:"bigint" json:"updateTime"`                          // 更新时间
}

// LoginRequest 定义了登录请求的结构
type LoginRequest struct {
	Account  string `json:"Account"`
	Password string `json:"password"`
}

func (u *User) TableName() string {
	return "user" // 用户表
}

// CheckPassword 验证提供的密码是否与用户的密码哈希值匹配
func (u *User) CheckPassword(password string) bool {
	// u.PassWord 存储的是哈希过的密码
	if u.PassWord == "" {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(u.PassWord), []byte(password))
	return err == nil
}

// CreateUser 创建用户
func CreateUser(user *User) error {
	session := dal.GetDb().NewSession()
	defer session.Close()

	// 加密密码
	if user.PassWord != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PassWord), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.PassWord = string(hashedPassword)
	}

	affected, err := session.Insert(user)
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("failed to create user")
	}
	return nil
}

// 获取
func GetUserByAccount(account string) (*User, error) {
	session := dal.GetDb().NewSession()
	defer session.Close()

	user := &User{Account: &account}
	has, err := session.Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// UpdateUser 更新用户
func UpdateUser(id uint, user *User) error {
	session := dal.GetDb().NewSession()
	defer session.Close()

	// 如果提供了新密码，则加密它
	if user.PassWord != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PassWord), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.PassWord = string(hashedPassword)
	}

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

	user := &User{Id: &id}
	has, err := session.Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// GetCount 获取用户数
func GetCount() (int64, error) {
	session := dal.GetDb().NewSession()
	defer session.Close()

	count, err := session.Count(new(User))
	if err != nil {
		return 0, err
	}
	return count, nil

}
