package contract

import "github.com/dgrijalva/jwt-go"

type TokenBody struct {
	UserID  uint `json:"user_id"`
	IsAdmin bool `json:"is_admin"`
}

type AccessToken struct {
	jwt.StandardClaims
	TokenBody
}
