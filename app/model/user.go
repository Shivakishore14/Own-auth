package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

//User details
type User struct {
	gorm.Model
	UserName string `gorm:"not null;unique;index" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Name     string `gorm:"not null" json:"name"`
	Email    string `gorm:"not null;unique;index" json:"email"`
	Phone    string `json:"phone"`
	Custom   string `gorm:"type:text" json:"custom"`
}

//IsValidLogin : for checking if credentials are valid
func (user User) IsValidLogin(db *gorm.DB) (User, bool) {
	tempUser := User{}
	db.Where("user_name=?", user.UserName).First(&tempUser)

	if tempUser.Password == user.Password && user.Password != "" {
		return tempUser, true
	}
	return tempUser, false
}

//CreateUser : for creating a new user
func (user User) CreateUser(db *gorm.DB) (string, error) {
	fmt.Println(user)
	if gdb := db.Create(&user); gdb.Error != nil {
		return "Check Given details", gdb.Error
	}
	return "Created user " + user.UserName, nil

}
