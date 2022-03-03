package services

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/kunioshi/hashed-notes/api/models"
	"golang.org/x/crypto/bcrypt"
)

// Generate a Bcrypt hash from the given `string`
func Bcrypt(p string) (string, error) {
	bp := []byte(p)
	bp, err := bcrypt.GenerateFromPassword(bp, bcrypt.DefaultCost)

	return string(bp), err
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

	// Start a DB transaction to make sure all entries are persisted together
	t := db.Begin()
	r := t.Create(&u)
	if r.Error != nil {
		return nil, r.Error
	}
	ec, err := CreateEmailConfirm(u)
	if err != nil {
		t.Rollback()
		return nil, err
	}
	r = t.Create(&ec)
	if r.Error != nil {
		t.Rollback()
		return nil, r.Error
	}
	t.Commit()

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

	_, ok = m["email"]
	var ec *models.EmailConfirm
	if ok {
		ec, err = CreateEmailConfirm(*u)
		if err != nil {
			return nil, err
		}
	}

	// Start a DB transaction to make sure all entries are persisted together
	t := db.Begin()
	q := t.Model(&u).Updates(m)
	if q.Error != nil {
		t.Rollback()
		return nil, q.Error
	}
	q = t.Create(&ec)
	if q.Error != nil {
		t.Rollback()
		return nil, q.Error
	}
	t.Commit()

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
