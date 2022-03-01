package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `form:"username" json:"username" binding:"required" gorm:"unique"`
	Password     string `form:"password" json:"-" binding:"required"`
	Email        string `form:"email" json:"email" binding:"required" gorm:"unique"`
	Confirmation string `form:"confirmation" json:"-"`
}

func NewUser(u string, p string, e string, c string) *User {
	nu := new(User)
	nu.Username = u
	nu.Password = p
	nu.Email = e
	nu.Confirmation = c

	return nu
}
