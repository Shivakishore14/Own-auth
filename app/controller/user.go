package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/Shivakishore14/Own-auth/app/model"
	"github.com/gorilla/mux"
)

//User to manipulate profile of a user
func User(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	user := model.User{}
	user.ID = uint(id)
	if r.Method == "GET" {
		userNew, err := user.UserData(db)
		if err != nil {
			webresponse("Error getting data", err, nil, w)
			log.Println(err)
			return
		}
		webresponse("Success", nil, userNew, w)
	} else if r.Method == "DELETE" {
		if err := user.Delete(db); err != nil {
			webresponse("Error deleting user", err, nil, w)
			log.Println(err)
			return
		}
		webresponse("Success", nil, nil, w)
	} else if r.Method == "POST" {
		var bodyBytes []byte
		bodyBytes, _ = ioutil.ReadAll(r.Body)

		if err := json.Unmarshal(bodyBytes, &user); err != nil {
			log.Println(err)
			webresponse("Error decoding json", err, nil, w)
			return
		}
		if err := user.Update(db); err != nil {
			log.Println(err)
			webresponse("Error Updating user", err, nil, w)
			return
		}
		webresponse("Success", nil, nil, w)
	} else if r.Method == "PUT" {
		var bodyBytes []byte
		bodyBytes, _ = ioutil.ReadAll(r.Body)

		if err := json.Unmarshal(bodyBytes, &user); err != nil {
			log.Println(err)
			webresponse("Error decoding json", err, nil, w)
			return
		}
		if err := user.Create(db); err != nil {
			log.Println(err)
			webresponse("Error Updating user", err, nil, w)
			return
		}
		webresponse("Success", nil, nil, w)
	}
}
