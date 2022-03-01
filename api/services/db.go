package services

import (
	"fmt"
	"log"

	"github.com/kunioshi/hashed-notes/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB = initDb()

// Initialize the database connection. It creates `*gorm.DB` and saves it in global `Db`.
func initDb() *gorm.DB {
	env := config.GetEnv()
	opt := "parseTime=true"
	srcName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?%s",
		env["DB_USER"],
		env["DB_PASS"],
		env["DB_HOST"],
		env["DB_PORT"],
		env["DB_NAME"],
		opt,
	)

	// TODO: Accept other DB drivers other than MySQL
	nd, err := gorm.Open(mysql.Open(srcName), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database. Error: %v", err)
	}

	return nd
}

// Retrieves the `database` connection
// Runs `initDb()` if needed
func GetDB() *gorm.DB {
	if Db == nil {
		Db = initDb()
	}

	return Db
}
