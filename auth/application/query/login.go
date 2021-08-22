package query

import "github.com/dgrijalva/jwt-go/v4"

type LoginQuery struct {
	Id       string
	Password string
}

type AuthTokenCliams struct {
	TokenUUID string `json:"tid"`
	UserID    string `json:"id"`
	// Role      []string `json:"role"`
	jwt.StandardClaims
}
