package models

import (
	"time"
)

type EmailConfirm struct {
	// Not using gorm.Model to avoid Soft Delete
	ID           uint `gorm:"primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	UserID       uint
	User         User
	Confirmation string
}

func NewEmailConfirm(u User, c string) *EmailConfirm {
	e := new(EmailConfirm)
	e.User = u
	e.Confirmation = c

	return e
}
