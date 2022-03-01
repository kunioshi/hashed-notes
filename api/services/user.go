package services

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kunioshi/hashed-notes/api/models"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

// Generate a Bcrypt hash from the given `string`
func Bcrypt(p string) (string, error) {
	bp := []byte(p)
	bp, err := bcrypt.GenerateFromPassword(bp, bcrypt.DefaultCost)

	return string(bp), err
}

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

func CreateUser(c *gin.Context) (*models.User, error) {
	var u CreateUserInput
	err := c.ShouldBindJSON(&u)
	if err != nil {
		return nil, err
	}

	ec, err := MD5(u.Email)
	if err != nil {
		return nil, err
	}

	hp, err := Bcrypt(u.Password)
	if err != nil {
		return nil, err
	}

	// Create and Persist User
	nu := models.NewUser(u.Username, hp, u.Email, ec)
	db := GetDB()
	r := db.Create(&u)
	if r.Error != nil {
		return nil, r.Error
	}

	return nu, nil
}

func GetUsers(c *gin.Context) ([]models.User, error) {
	var us []models.User
	r := GetDB().Find(&us)
	if r.Error != nil {
		return nil, r.Error
	}

	return us, nil
}

func GetUser(c *gin.Context) (*models.User, error) {
	id := c.Param("id")
	var u *models.User

	r := GetDB().First(&u, id)
	if r.Error != nil {
		return nil, r.Error
	}

	return u, nil
}

func UpdateUser(c *gin.Context) (*models.User, error) {
	var u *models.User

	data, _ := c.Get("password")
	fmt.Fprintf(c.Writer, "data: %v", data)
	c.ShouldBindJSON(&u)

	c.JSON(http.StatusOK, gin.H{"user": u})

	return u, nil
}
