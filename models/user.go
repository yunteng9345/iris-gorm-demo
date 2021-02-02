package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	Name     string
	Email    string
	Mobile   string
	QQ       string
	Gender   int
	Age      int
	Remark   string
	Token    string `gorm:"-"`
	Session  string `gorm:"-"`
}
