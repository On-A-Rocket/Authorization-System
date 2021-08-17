package command

import (
	"time"

	"github.com/On-A-Rocket/Authorization-System/account/domain/entity"
	iRepository "github.com/On-A-Rocket/Authorization-System/account/domain/iReporitory"
	"golang.org/x/crypto/bcrypt"
)

type AccountCommandHandler struct {
	repository iRepository.AccountInterface
}

func newAccountCommandHandler(
	repository iRepository.AccountInterface) *AccountCommandHandler {
	return &AccountCommandHandler{repository: repository}
}

func (handler *AccountCommandHandler) CreateAccountHandler(
	command CreateAccountCommand) error {
	hashedPassword, err := getHashedPassword(command.Password)
	if err != nil {
		return err
	}

	hireDate, err := time.Parse("2006-01-02 15:04:05", command.HireDate)
	if err != nil {
		return err
	}

	requestEntity := entity.Account{
		Id:          command.Id,
		Password:    hashedPassword,
		Name:        command.Name,
		Email:       command.Email,
		PhoneNumber: command.PhoneNumber,
		WorkCode:    command.WorkCode,
		HireDate:    hireDate,
	}

	transaction := handler.repository.StartTransaction()
	err = handler.repository.CreateAccount(&requestEntity)
	if err != nil {
		transaction.Rollback()
	}
	transaction.Commit()

	return err
}

func getHashedPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
