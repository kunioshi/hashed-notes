package services

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/kunioshi/hashed-notes/api/models"
	"golang.org/x/crypto/bcrypt"
)

type UserInput struct {
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
	var u models.User
	err := c.ShouldBindJSON(&u)
	if err != nil {
		return nil, err
	}

	u.Password, err = Bcrypt(u.Password)
	if err != nil {
		return nil, err
	}

	db := GetDB()
	r := db.Create(&u)
	if r.Error != nil {
		return nil, r.Error
	}

	return &u, nil
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

	u := new(models.User)
	db := GetDB()
	err := db.Where("id = ?", id).First(&u).Error
	if err != nil {
		return nil, err
	}

	return u, nil
}

func UpdateUser(c *gin.Context) (*models.User, error) {
	id := c.Param("id")

	u := new(models.User)
	db := GetDB()
	err := db.Where("id = ?", id).First(&u).Error
	if err != nil {
		return nil, err
	}

	// Parse JSON body
	r, _ := c.GetRawData()
	var m map[string]interface{}
	err = json.Unmarshal(r, &m)
	if err != nil {
		return nil, err
	}

	// Remove empty fields
	for k, v := range m {
		if v == "" || v == nil {
			delete(m, k)
		}
	}

	// Hash the new password, if there is one
	p, ok := m["password"]
	if ok {
		m["password"], _ = Bcrypt(p.(string))
	}

	q := db.Model(&u).Updates(m)
	if q.Error != nil {
		return nil, q.Error
	}

	return u, nil
}

func DeleteUser(c *gin.Context) error {
	id := c.Param("id")

	u := new(models.User)
	db := GetDB()
	err := db.Where("id = ?", id).First(&u).Error
	if err != nil {
		return err
	}

	err = db.Delete(&u).Error
	if err != nil {
		return err
	}

	return nil
}
