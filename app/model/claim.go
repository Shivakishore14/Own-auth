package model

import jwt "github.com/dgrijalva/jwt-go"

//UserClaim to store data in jwt
type UserClaim struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Email    string `json:"email"`
}
