package services

import (
	"fmt"
	"log"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *Database

func init() {
	db, err := gorm.Open("sqlite3", "blog_api_test.db")
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})
}

func Test_GetUsers(t *testing.T) {
	fmt.Println("Testing Proper")
}
