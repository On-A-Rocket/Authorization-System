package command

import (
	"time"

	"github.com/On-A-Rocket/Authorization-System/auth/domain/dto"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (handler *Command) CreateToken(context *gin.Context, id string) (dto.Token, error) {
	result := dto.Token{}
	jwtKey := handler.config.Auth().AccessSecret()

	token := &TokenDetails{}
	token.AccessUUID = uuid.NewString()
	token.AccessTokenExpiration = time.Now().Add(time.Minute * 15)
	token.RefreshUUID = uuid.NewString()
	token.RefreshTokenExpiration = time.Now().Add(time.Hour * 24 * 7)

	accessTokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256,
		handler.accountToTokenCliams(id, token.AccessUUID, token.AccessTokenExpiration))
	accessToken, err := accessTokenClaim.SignedString([]byte(jwtKey))
	if err != nil {
		return result, err
	}
	result.AccessToken = accessToken

	refreshTokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256,
		handler.accountToTokenCliams(id, token.RefreshUUID, token.RefreshTokenExpiration))
	refreshToken, err := refreshTokenClaim.SignedString([]byte(jwtKey))
	if err != nil {
		return result, err
	}
	result.RefreshToken = refreshToken

	accessTokenExpiration := time.Unix(token.AccessTokenExpiration.Unix(), 0)
	refreshTokenExpiration := time.Unix(token.RefreshTokenExpiration.Unix(), 0)
	now := time.Now()

	redis := handler.config.Redis().Client()
	if accessError := redis.Set(context, token.AccessUUID, id, accessTokenExpiration.Sub(now)).Err(); err != nil {
		return result, accessError
	}
	if refreshError := redis.Set(context, token.RefreshUUID, id, refreshTokenExpiration.Sub(now)).Err(); err != nil {
		return result, refreshError
	}

	return result, nil
}

func (handler *Command) accountToTokenCliams(
	id string, token string, expiration time.Time) TokenCliams {
	return TokenCliams{
		TokenUUID:      token,
		UserID:         id,
		StandardClaims: jwt.StandardClaims{ExpiresAt: jwt.At(expiration)},
	}
}
