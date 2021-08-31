package repository

import (
	iRepository "github.com/On-A-Rocket/Authorization-System/auth/domain/iReporitory"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) iRepository.Interface {
	return &Repository{db}
}

func (repo *Repository) StartTransaction() *gorm.DB {
	return repo.db.Begin()
}
