package query

import (
	"time"

	"github.com/On-A-Rocket/Authorization-System/auth/config"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
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

func (handler *LoginQueryHandler) LoginHandler(
	context *gin.Context,
	query LoginQuery) error {
	if err := handler.createToken(context, query.Id); err != nil {
		return err
	}

	return nil
}

func (handler *LoginQueryHandler) createToken(context *gin.Context, userID string) error {
	jwtKey := handler.config.Auth().AccessSecret()

	token := &TokenDetails{}
	token.AccessUUID = uuid.NewString()
	token.AccessTokenExpiration = time.Now().Add(time.Minute * 15)
	token.RefreshUUID = uuid.NewString()
	token.RefreshTokenExpiration = time.Now().Add(time.Hour * 24 * 7)

	accessTokenConfig := AuthTokenCliams{
		TokenUUID: token.AccessUUID,
		UserID:    userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(token.AccessTokenExpiration),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &accessTokenConfig)

	accessTokenValue, err := accessToken.SignedString([]byte(jwtKey))
	if err != nil {
		return err
	}
	token.AccessToken = accessTokenValue

	refreshTokenConfig := AuthTokenCliams{
		TokenUUID: token.RefreshUUID,
		UserID:    userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(token.RefreshTokenExpiration),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &refreshTokenConfig)

	refreshTokenValue, err := refreshToken.SignedString([]byte(jwtKey))
	if err != nil {
		return err
	}
	token.RefreshToken = refreshTokenValue

	accessTokenExpiration := time.Unix(token.AccessTokenExpiration.Unix(), 0)
	refreshTokenExpiration := time.Unix(token.RefreshTokenExpiration.Unix(), 0)
	now := time.Now()

	redis := handler.config.Redis().Client()
	if accessError := redis.Set(context, token.AccessUUID, token.AccessToken, accessTokenExpiration.Sub(now)).Err(); err != nil {
		return accessError
	}
	if refreshError := redis.Set(context, token.RefreshUUID, token.RefreshToken, refreshTokenExpiration.Sub(now)).Err(); err != nil {
		return refreshError
	}

	return nil
}
