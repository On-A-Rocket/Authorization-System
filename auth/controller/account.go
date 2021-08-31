package controller

import (
	"net/http"

	"github.com/On-A-Rocket/Authorization-System/auth/application/command"
	"github.com/On-A-Rocket/Authorization-System/auth/domain/dto"
	"github.com/gin-gonic/gin"
)

func (ctl *Controller) accountRouting(router *gin.Engine) {
	routerGroup := router.Group("/account")
	{
		routerGroup.POST("", ctl.CreateAccount)
	}
}

// @Summary Create Account
// @Description 회원가입
// @Tags Account
// @Accept json
// @Produce json
// @Router /account [post]
// @Param Account body dto.CreateAccount true "create account"
// @Success 200 {string} create account
func (ctl *Controller) CreateAccount(context *gin.Context) {
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

	if err := ctl.command.CreateAccountHandler(command); err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusCreated, "create account")
}
