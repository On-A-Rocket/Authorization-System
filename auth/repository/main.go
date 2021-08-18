package repository

import (
	iRepository "github.com/On-A-Rocket/Authorization-System/auth/domain/iReporitory"
	"gorm.io/gorm"
)

type Repository struct {
	Account AccountRepository
}

func NewRepository(db *gorm.DB) *iRepository.Repository {
	return &iRepository.Repository{
		Account: newAccountRepository(db),
	}
}
