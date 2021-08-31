package query

import (
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
)

func (handler *Query) Auth(c *gin.Context) error {
	token, err := handler.veriftToken(c)
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		tokenId, ok := claims["tid"].(string)
		if !ok {
			return err
		}
		userId, ok := claims["id"].(string)
		if !ok {
			return err
		}

		id, err := handler.getTokenValue(c, tokenId)
		if err != nil {
			return err
		}
		if userId != id {
			return fmt.Errorf("unauthorized")
		}
	}

	return nil
}

func (handler *Query) veriftToken(c *gin.Context) (*jwt.Token, error) {
	tokenString := handler.extractToken(c)
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (handler *Query) extractToken(c *gin.Context) string {
	bearToken := c.Request.Header.Get("Authorization")
	bearArray := strings.Split(bearToken, " ")
	if len(bearArray) == 2 {
		return bearArray[1]
	}
	return ""
}

func (handler *Query) getTokenValue(c *gin.Context, tokenId string) (string, error) {
	redis := handler.config.Redis().Client()
	token, err := redis.Get(c, tokenId).Result()
	if err != nil {
		return "", err
	}
	return token, nil
}
