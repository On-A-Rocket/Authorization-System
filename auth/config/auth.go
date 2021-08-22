package config

import (
	"os"
)

type AuthInterface interface {
	AccessSecret() string
}

type Auth struct {
	accessSecret string
}

func newAuthConfig() *Auth {
	auth := &Auth{
		accessSecret: "AccessTokenSecret",
	}

	if secret := os.Getenv("ACCESS_SECRET"); secret != "" {
		auth.accessSecret = secret
	}

	return auth
}

func (auth *Auth) AccessSecret() string {
	return auth.accessSecret
}
