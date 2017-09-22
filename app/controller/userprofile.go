package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Shivakishore14/Own-auth/app/model"
	"github.com/gorilla/mux"
)

//UserProfile to view profile of a user
func UserProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	user := model.User{}
	user.ID = uint(id)
	user, err := user.UserData(db)
	if err != nil {
		webresponse("Error getting data", err, nil, w)
		log.Println(err)
		return
	}
	webresponse("Success", nil, user, w)
}
