package command

import (
	"github.com/On-A-Rocket/Authorization-System/auth/config"
	iRepository "github.com/On-A-Rocket/Authorization-System/auth/domain/iReporitory"
)

type Command struct {
	repository iRepository.Interface
	config     config.Interface
}

func NewCommand(
	repository iRepository.Interface,
	config config.Interface) *Command {
	return &Command{
		repository: repository,
		config:     config,
	}
}
