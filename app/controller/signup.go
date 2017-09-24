package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Shivakishore14/Own-auth/app/model"
)

//UserSignUp : for creating new user
func UserSignUp(w http.ResponseWriter, r *http.Request) {
	var user model.User
	var bodyBytes []byte
	bodyBytes, _ = ioutil.ReadAll(r.Body)

	if err := json.Unmarshal(bodyBytes, &user); err != nil {
		log.Println(err)
		webresponse("Error decoding json", err, nil, w)
		return
	}

	msg, err := user.CreateUser(db)
	webresponse(msg, err, nil, w)
}
