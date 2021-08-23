package command

import (
	"time"

	"github.com/On-A-Rocket/Authorization-System/auth/config"
	"github.com/On-A-Rocket/Authorization-System/auth/domain/dto"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LoginCommandHandler struct {
	config config.Interface
}

func newLoginCommandHandler(config config.Interface) *LoginCommandHandler {
	return &LoginCommandHandler{config}
}

func (handler *LoginCommandHandler) CreateToken(context *gin.Context, account AccountCommand) (dto.Token, error) {
	result := dto.Token{}
	jwtKey := handler.config.Auth().AccessSecret()

	token := &TokenDetails{}
	token.AccessUUID = uuid.NewString()
	token.AccessTokenExpiration = time.Now().Add(time.Minute * 15)
	token.RefreshUUID = uuid.NewString()
	token.RefreshTokenExpiration = time.Now().Add(time.Hour * 24 * 7)

	// accessTokenClaim := handler.accountToTokenCliams(account, token.AccessUUID, token.AccessTokenExpiration)
	accessTokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256,
		handler.accountToTokenCliams(account, token.AccessUUID, token.AccessTokenExpiration))
	accessToken, err := accessTokenClaim.SignedString([]byte(jwtKey))
	if err != nil {
		return result, err
	}
	token.AccessToken = accessToken
	result.AccessToken = accessToken

	// refreshTokenConfig := handler.accountToTokenCliams(account, token.RefreshUUID, token.RefreshTokenExpiration)
	refreshTokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256,
		handler.accountToTokenCliams(account, token.RefreshUUID, token.RefreshTokenExpiration))
	refreshToken, err := refreshTokenClaim.SignedString([]byte(jwtKey))
	if err != nil {
		return result, err
	}
	token.RefreshToken = refreshToken
	result.RefreshToken = refreshToken

	accessTokenExpiration := time.Unix(token.AccessTokenExpiration.Unix(), 0)
	refreshTokenExpiration := time.Unix(token.RefreshTokenExpiration.Unix(), 0)
	now := time.Now()

	redis := handler.config.Redis().Client()
	if accessError := redis.Set(context, token.AccessUUID, token.AccessToken, accessTokenExpiration.Sub(now)).Err(); err != nil {
		return result, accessError
	}
	if refreshError := redis.Set(context, token.RefreshUUID, token.RefreshToken, refreshTokenExpiration.Sub(now)).Err(); err != nil {
		return result, refreshError
	}

	return result, nil
}

func (handler *LoginCommandHandler) accountToTokenCliams(
	command AccountCommand, token string, expiration time.Time) TokenCliams {
	return TokenCliams{
		TokenUUID:      token,
		Name:           command.Name,
		Email:          command.Email,
		PhoneNumber:    command.PhoneNumber,
		DepartmentCode: command.DepartmentCode,
		PositionCode:   command.PositionCode,
		AuthorityCode:  command.AuthorityCode,
		FirstPaymentId: command.FirstPaymentId,
		FinalPaymentId: command.FinalPaymentId,
		WorkCode:       command.WorkCode,
		TotalAnnual:    command.TotalAnnual,
		UseAnnual:      command.UseAnnual,
		RemainAnnual:   command.RemainAnnual,
		HireDate:       command.HireDate,
		StandardClaims: jwt.StandardClaims{ExpiresAt: jwt.At(expiration)},
	}
}
