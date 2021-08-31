package query

import (
	"github.com/On-A-Rocket/Authorization-System/auth/domain/entity"
	"golang.org/x/crypto/bcrypt"
)

func (handler *Query) LoginHandler(query LoginQuery) error {
	account := entity.Account{}
	condition := entity.Account{Id: query.Id}
	if err := handler.db.Select("id, password").Where(condition).Find(&account).Error; err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(account.Password),
		[]byte(query.Password)); err != nil {
		return err
	}

	return nil
}
