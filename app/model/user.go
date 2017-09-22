package model

import (
	"encoding/json"
	"log"

	"github.com/jinzhu/gorm"
)

//User details
type User struct {
	gorm.Model
	UserName string  `gorm:"not null;unique;index" json:"username"`
	Password string  `gorm:"not null" json:"password"`
	Name     string  `gorm:"not null" json:"name"`
	Email    string  `gorm:"not null;unique;index" json:"email"`
	Phone    string  `json:"phone"`
	Custom   string  `gorm:"type:text" json:"-"`
	Fields   []Field `gorm:"-" json:"fields"`
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
	if gobj := db.Create(&user); gobj.Error != nil {
		return "Check Given details", gobj.Error
	}
	return "Created user " + user.UserName, nil

}

//UserData : to return data about user
func (user User) UserData(db *gorm.DB) (User, error) {
	if gobj := db.Where("id=?", user.ID).First(&user); gobj.Error != nil {
		log.Println(gobj.Error)
		return user, gobj.Error
	}
	fields := make([]Field, 0, 100)

	if err := json.Unmarshal([]byte(user.Custom), &fields); err != nil {
		log.Print(err)
	}
	user.Fields = fields
	return user, nil
}

//Delete user
func (user User) Delete(db *gorm.DB) error {
	if gobj := db.Where("id=?", user.ID).First(&user); gobj.Error != nil {
		return gobj.Error
	}
	if gobj := db.Delete(&user); gobj.Error != nil {
		return gobj.Error
	}
	return nil
}

//Update user
func (user User) Update(db *gorm.DB) error {
	if gobj := db.Where("id=?", user.ID).First(&user); gobj.Error != nil {
		return gobj.Error
	}
	if gobj := db.Save(&user); gobj.Error != nil {
		return gobj.Error
	}
	return nil
}

//Create user
func (user User) Create(db *gorm.DB) error {
	if gobj := db.Create(&user); gobj.Error != nil {
		return gobj.Error
	}
	return nil
}
