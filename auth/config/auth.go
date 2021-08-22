package config

import (
	"os"
	"strconv"
)

type AuthInterface interface {
	AccessSecret() string
	AccessExpiration() int
}

type Auth struct {
	accessSecret     string
	accessExpiration string
}

func newAuthConfig() *Auth {
	auth := &Auth{
		accessSecret:     "AccessTokenSecret",
		accessExpiration: "15",
	}

	if secret := os.Getenv("ACCESS_SECRET"); secret != "" {
		auth.accessSecret = secret
	}
	if expiration := os.Getenv("ACCESS_ENPIRATION"); expiration != "" {
		auth.accessExpiration = expiration
	}

	return auth
}

func (auth *Auth) AccessSecret() string {
	return auth.accessSecret
}

func (auth *Auth) AccessExpiration() int {
	minutes, err := strconv.Atoi(auth.accessExpiration)
	if err != nil {
		panic(err)
	}
	return minutes
}
