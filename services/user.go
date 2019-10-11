package services

import "github.com/jinzhu/gorm"

//User - Model
type User struct {
	gorm.Model
	Username, Password, Token, Email, Fullname string
}

//FetchUsers - function to retrieve all the users in the table.
func (db *Database) FetchUsers() {

}
