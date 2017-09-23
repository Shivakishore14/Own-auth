package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/Shivakishore14/Own-auth/app/model"
	jwt "github.com/dgrijalva/jwt-go"
)

//UserLogin : for Login functionality
func UserLogin(w http.ResponseWriter, r *http.Request) {
	// username := r.FormValue("username")
	// password := r.FormValue("password")
	// fmt.Println(username, password)
	// user := model.User{UserName: username, Password: password}
	var user model.User
	var bodyBytes []byte
	bodyBytes, _ = ioutil.ReadAll(r.Body)

	if err := json.Unmarshal(bodyBytes, &user); err != nil {
		log.Println(err)
		webresponse("Error decoding json", err, nil, w)
		return
	}
	user, check := user.IsValidLogin(db)
	//user.RegisterLogin(db, check)
	if check {
		expireToken := time.Now().Add(time.Hour * 1).Unix()
		claims := model.UserClaim{}
		claims.StandardClaims = jwt.StandardClaims{ExpiresAt: expireToken, Issuer: "localhost:9000"}
		claims.Username = user.UserName
		claims.Email = user.Email
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signedToken, _ := token.SignedString([]byte("secret"))
		data := make(map[string]interface{})
		data["profile"] = user
		data["token"] = signedToken
		webresponse("success", nil, data, w)
	} else {
		webresponse("check creds", nil, nil, w)
	}
}
