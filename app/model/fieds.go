package model

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

//UserFields model
type UserFields struct {
	UserData User    `json:"user"`
	Fields   []Field `json:"fields"`
}

//Field model
type Field struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

//Save to save the feilds
func (userFields UserFields) Save(db *gorm.DB) error {
	var user User
	if gobj := db.Where("user_name", userFields.UserData.UserName).First(&user); gobj.Error != nil {
		return gobj.Error
	}
	data, err := json.Marshal(userFields.Fields)
	if err != nil {
		return err
	}
	user.Custom = string(data)
	if gobj := db.Save(user); gobj.Error != nil {
		return gobj.Error
	}
	return nil
}
