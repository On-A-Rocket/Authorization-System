package controller

import (
	"github.com/On-A-Rocket/Authorization-System/auth/application/command"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Account *AccountController
}

func NewController(command command.Command) *Controller {
	return &Controller{
		Account: newAccountController(*command.Account),
	}
}

func (ctl *Controller) Routing(router *gin.Engine) {
	ctl.Account.accountRouting(router)
}
