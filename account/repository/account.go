package repository

import (
	"github.com/On-A-Rocket/Authorization-System/account/domain/entity"
	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func newAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db}
}

func (repo *AccountRepository) CreateAccount(account *entity.Account) error {
	return repo.db.Create(&account).Error
}

func (repo *AccountRepository) StartTransaction() *gorm.DB {
	return repo.db.Begin()
}
