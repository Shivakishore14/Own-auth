package controller

import (
	"net/http"

	"github.com/Shivakishore14/Own-auth/app/model"
)

//UserSignUp : for creating new user
func UserSignUp(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	name := r.FormValue("name")
	phone := r.FormValue("phone")
	email := r.FormValue("email")
	user := model.User{UserName: username, Password: password, Name: name, Phone: phone, Email: email}

	msg, err := user.CreateUser(db)
	webresponse(msg, err, nil, w)
}
