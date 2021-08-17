package iRepository

import (
	"github.com/On-A-Rocket/Authorization-System/account/domain/entity"
	"gorm.io/gorm"
)

type AccountInterface interface {
	CreateAccount(*entity.Account) error
	StartTransaction() *gorm.DB
}
