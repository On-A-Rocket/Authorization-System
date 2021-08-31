package repository

import "github.com/On-A-Rocket/Authorization-System/auth/domain/entity"

func (repo *Repository) CreateAccount(account *entity.Account) error {
	return repo.db.Create(&account).Error
}
