package model

import (
	"time"
)

type User struct {
	Id       uint       `xorm:"'id' AUTO_INCREMENT" json:"id"`
	Name     string     `xorm:"size:50" json:"name"`                        // Name of the user
	Age      int        `xorm:"size:3" json:"age"`                          // Age of the user
	Birthday *time.Time `json:"birthday,omitempty"`                         // Date of birth of the user (optional)
	Email    string     `xorm:"type:varchar(50);unique_index" json:"email"` // Email address of the user (unique)
	PassWord string     `xorm:"type:varchar(25)" json:"password"`           // Password for the user's account
}

func TableName() string {
	return "user"
}
