package command

import (
	iRepository "github.com/On-A-Rocket/Authorization-System/auth/domain/iReporitory"
)

type Command struct {
	Account *AccountCommandHandler
}

func NewCommand(repository iRepository.RepositoryInterface) *Command {
	return &Command{
		Account: newAccountCommandHandler(repository.GetAccount()),
	}
}