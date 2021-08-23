package query

import (
	"github.com/On-A-Rocket/Authorization-System/auth/config"
	"github.com/On-A-Rocket/Authorization-System/auth/domain/dto"
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

func (handler *LoginQueryHandler) LoginHandler(query LoginQuery) (dto.Account, error) {
	account := entity.Account{}
	dto := dto.Account{}
	condition := entity.Account{Id: query.Id}
	if err := handler.db.Where(condition).Find(&account).Error; err != nil {
		return dto, err
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(account.Password),
		[]byte(query.Password)); err != nil {
		return dto, err
	}

	return handler.entityToDTO(account), nil
}

func (handler *LoginQueryHandler) entityToDTO(entity entity.Account) dto.Account {
	return dto.Account{
		Name:           entity.Name,
		Email:          entity.Email,
		PhoneNumber:    entity.PhoneNumber,
		DepartmentCode: entity.DepartmentCode,
		PositionCode:   entity.PositionCode,
		AuthorityCode:  entity.AuthorityCode,
		FirstPaymentId: entity.FirstPaymentId,
		FinalPaymentId: entity.FinalPaymentId,
		WorkCode:       entity.WorkCode,
		TotalAnnual:    entity.TotalAnnual,
		UseAnnual:      entity.UseAnnual,
		RemainAnnual:   entity.RemainAnnual,
		HireDate:       entity.HireDate,
	}
}
