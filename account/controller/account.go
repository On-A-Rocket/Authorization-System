package controller

import (
	"net/http"

	"github.com/On-A-Rocket/Authorization-System/account/application/command"
	"github.com/On-A-Rocket/Authorization-System/account/domain/dto"
	"github.com/gin-gonic/gin"
)

type AccountController struct {
	commandHandler command.AccountCommandHandler
}

func newAccountController(
	commandHandler command.AccountCommandHandler) *AccountController {
	return &AccountController{commandHandler: commandHandler}
}

func (ctl *AccountController) accountRouting(router *gin.Engine) {
	routerGroup := router.Group("/account")
	{
		routerGroup.POST("", ctl.CreateAccount)
	}
}

// @Summary Create Account
// @Description Create Account
// @Tags Account
// @Accept json
// @Produce json
// @Router /account [post]
// @Param Account body dto.CreateAccount true "create account"
// @Success 200 {string} create account
func (ctl *AccountController) CreateAccount(context *gin.Context) {
	dto := &dto.CreateAccount{}
	if err := context.ShouldBind(&dto); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

	command := command.CreateAccountCommand{
		Id:          dto.Id,
		Password:    dto.Password,
		Name:        dto.Name,
		Email:       dto.Email,
		PhoneNumber: dto.PhoneNumber,
		WorkCode:    dto.WorkCode,
		HireDate:    dto.HireDate,
	}

	if err := ctl.commandHandler.CreateAccountHandler(command); err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusCreated, "create account")
}
