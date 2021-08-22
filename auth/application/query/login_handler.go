package query

import (
	"time"

	"github.com/On-A-Rocket/Authorization-System/auth/config"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LoginQueryHandler struct {
	db     *gorm.DB
	config config.Interface
}

func newLoginQueryHandler(db *gorm.DB, config config.Interface) *LoginQueryHandler {
	return &LoginQueryHandler{db, config}
}

// func (handler *LoginQueryHandler) LoginHandler(
// 	query LoginQuery) error {

// }

func (handler *LoginQueryHandler) CreateToken(userID string) (string, error) {
	at := AuthTokenCliams{
		TokenUUID: uuid.NewString(),
		UserID:    userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Minute * time.Duration(handler.config.Auth().AccessExpiration()))),
		},
	}
	atoken := jwt.NewWithClaims(jwt.SigningMethodHS256, &at)
	jwtKey := handler.config.Auth().AccessSecret()

	token, err := atoken.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return token, nil
}
