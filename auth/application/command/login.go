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
	TokenUUID      string `json:"tid"`
	Name           string
	Email          string
	PhoneNumber    string
	DepartmentCode int
	PositionCode   int
	AuthorityCode  int
	FirstPaymentId string
	FinalPaymentId string
	WorkCode       int
	TotalAnnual    float32
	UseAnnual      float32
	RemainAnnual   float32
	HireDate       string
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

type AccountCommand struct {
	Name           string
	Email          string
	PhoneNumber    string
	DepartmentCode int
	PositionCode   int
	AuthorityCode  int
	FirstPaymentId string
	FinalPaymentId string
	WorkCode       int
	TotalAnnual    float32
	UseAnnual      float32
	RemainAnnual   float32
	HireDate       string
}
