package command

import (
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

type LoginCommand struct {
	Id string
	// Password string
}

type TokenCliams struct {
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
