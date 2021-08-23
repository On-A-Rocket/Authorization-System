package query

import (
	"github.com/On-A-Rocket/Authorization-System/auth/config"
	"github.com/On-A-Rocket/Authorization-System/auth/domain/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginQueryHandler struct {
	db     *gorm.DB
	config config.Interface
}

func newLoginQueryHandler(db *gorm.DB, config config.Interface) *LoginQueryHandler {
	return &LoginQueryHandler{db, config}
}

func (handler *LoginQueryHandler) LoginHandler(query LoginQuery) error {
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
