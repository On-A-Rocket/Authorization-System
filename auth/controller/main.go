package controller

import (
	"github.com/On-A-Rocket/Authorization-System/auth/application/command"
	"github.com/On-A-Rocket/Authorization-System/auth/application/query"
	"github.com/On-A-Rocket/Authorization-System/auth/config"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	query   query.Query
	command command.Command
	config  config.Interface
}

func NewController(
	query query.Query,
	command command.Command,
	config config.Interface) *Controller {
	return &Controller{
		query:   query,
		command: command,
		config:  config,
	}
}

func (ctl *Controller) Routing(router *gin.Engine) {
	ctl.accountRouting(router)
	ctl.loginRouting(router)
}
