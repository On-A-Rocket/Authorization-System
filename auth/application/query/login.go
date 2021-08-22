package query

import (
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

type LoginQuery struct {
	Id       string
	Password string
}

type AuthTokenCliams struct {
	TokenUUID string `json:"tid"`
	UserID    string `json:"id"`
	jwt.StandardClaims
}

type TokenDetails struct {
	AccessToken            string
	RefreshToken           string
	AccessUUID             string
	RefreshUUID            string
	AccessTokenExpiration  time.Time
	RefreshTokenExpiration time.Time
}
