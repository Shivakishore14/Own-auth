package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Shivakishore14/Own-auth/app/model"
	//mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var database = "own_auth"
var user = "test"
var password = "test"

var db *gorm.DB

func init() {
	var err error

	if db, err = gorm.Open("mysql", user+":"+password+"@/"+database+"?charset=utf8&parseTime=True&loc=Local"); err != nil {
		log.Fatal("Error Connecting to database")
	}

	if db.HasTable(&model.User{}) == false {
		log.Print("Created new User table")
		db.CreateTable(&model.User{})
	}
}

func webresponse(msg string, err error, data interface{}, w http.ResponseWriter) (string, error) {
	obj := model.WebResponse{}
	obj.Message = msg
	obj.Error = err
	obj.Data = data

	var resErr error
	var resTxt string

	if jsonData, e := json.Marshal(obj); e != nil {
		resErr = e
	} else {
		resTxt = string(jsonData)
	}
	fmt.Fprint(w, resTxt)
	return resTxt, resErr
}
