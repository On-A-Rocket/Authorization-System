package query

import "github.com/dgrijalva/jwt-go/v4"

type TokenCliams struct {
	TokenUUID string `json:"tid"`
	UserID    string `json:"id"`
	jwt.StandardClaims
}
