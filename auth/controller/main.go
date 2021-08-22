package controller

import (
	"github.com/On-A-Rocket/Authorization-System/auth/application/command"
	"github.com/On-A-Rocket/Authorization-System/auth/application/query"
	"github.com/On-A-Rocket/Authorization-System/auth/config"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Account *AccountController
	Login   *LoginController
}

func NewController(
	query query.Query,
	command command.Command,
	config config.Interface) *Controller {
	return &Controller{
		Account: newAccountController(*command.Account),
		Login:   newLoginController(*query.Login),
	}
}

func (ctl *Controller) Routing(router *gin.Engine) {
	ctl.Account.accountRouting(router)
	ctl.Login.loginRouting(router)
}
