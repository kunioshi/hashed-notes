package services

import (
	"crypto/md5"
	"encoding/hex"
	"errors"

	"github.com/kunioshi/hashed-notes/api/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Generate a unique MD5 hash from the given `string`
func MD5(e string) (string, error) {
	bc, err := bcrypt.GenerateFromPassword([]byte(e), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	h := md5.New()
	h.Write(bc)

	return hex.EncodeToString(h.Sum(nil)), nil
}

func CreateEmailConfirm(u models.User) (*models.EmailConfirm, error) {
	c, err := MD5(u.Email)
	if err != nil {
		return nil, err
	}

	// TODO: Send email

	return models.NewEmailConfirm(u, c), nil
}

func ConfirmEmail(e string, c string) (bool, error) {
	db := GetDB()

	u := models.User{}
	err := db.Where("email = ?", e).First(&u).Error
	if err != nil {
		return false, errors.New("user not found")
	}

	ec, err := IsUserConfirmed(u, db)
	if err != nil {
		return false, err
	} else if ec == nil {
		return false, errors.New("user's email already confirmed")
	}

	if ec.Confirmation != c {
		return false, errors.New("confirmation code incorrect")
	}

	err = db.Delete(&ec).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func IsUserConfirmed(u models.User, db *gorm.DB) (*models.EmailConfirm, error) {
	ec := new(models.EmailConfirm)
	err := db.Preload("User").Find(&ec).Error

	return ec, err
}
