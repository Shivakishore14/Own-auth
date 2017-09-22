package controller

import (
	"log"
	"net/http"

	"github.com/Shivakishore14/Own-auth/app/model"
)

//ListUsers controller for listing users
func ListUsers(w http.ResponseWriter, r *http.Request) {
	var userList []model.User
	if gobj := db.Find(&userList); gobj.Error != nil {
		log.Println(gobj.Error)
		webresponse("Error getting list", gobj.Error, nil, w)
		return
	}
	webresponse("Success", nil, userList, w)

}
