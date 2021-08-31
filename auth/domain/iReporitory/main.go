package iRepository

import (
	"github.com/On-A-Rocket/Authorization-System/auth/domain/entity"
	"gorm.io/gorm"
)

type Interface interface {
	StartTransaction() *gorm.DB
	CreateAccount(*entity.Account) error
}
