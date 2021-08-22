package command

import (
	"github.com/On-A-Rocket/Authorization-System/auth/config"
	iRepository "github.com/On-A-Rocket/Authorization-System/auth/domain/iReporitory"
)

type Command struct {
	Account *AccountCommandHandler
}

func NewCommand(
	repository iRepository.RepositoryInterface,
	config config.Interface) *Command {
	return &Command{
		Account: newAccountCommandHandler(repository.GetAccount()),
	}
}
