package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `form:"username" json:"username" gorm:"unique"`
	Password string `form:"password" json:"password"`
	Email    string `form:"email" json:"email" binding:"email" gorm:"unique"`
}

func NewUser(u string, p string, e string) *User {
	nu := new(User)
	nu.Username = u
	nu.Password = p
	nu.Email = e

	return nu
}
