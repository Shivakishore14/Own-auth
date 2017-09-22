package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Shivakishore14/Own-auth/app/model"
)

//AddFields controller
func AddFields(w http.ResponseWriter, r *http.Request) {
	var userFields model.UserFields
	var bodyBytes []byte
	bodyBytes, _ = ioutil.ReadAll(r.Body)

	if err := json.Unmarshal(bodyBytes, &userFields); err != nil {
		log.Println(err)
		webresponse("Error decoding json", err, nil, w)
		return
	}
	if err := userFields.Save(db); err != nil {
		log.Println(err)
		webresponse("Error saving data", err, nil, w)
		return
	}
	webresponse("Success", nil, nil, w)
}
