package controller

import (
	"net/http"

	"github.com/On-A-Rocket/Authorization-System/auth/application/query"
	"github.com/On-A-Rocket/Authorization-System/auth/domain/dto"
	"github.com/gin-gonic/gin"
)

func (ctl *Controller) loginRouting(router *gin.Engine) {
	router.POST("/login", ctl.Login)
}

// @Summary Login
// @Description Login
// @Tags Login
// @Accept json
// @Produce json
// @Router /login [post]
// @Param login body dto.Login true "login"
// @Success 200 {object} dto.Token
func (ctl *Controller) Login(context *gin.Context) {
	dto := &dto.Login{}
	if err := context.ShouldBind(&dto); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

	query := query.LoginQuery{
		Id:       dto.Id,
		Password: dto.Password,
	}
	if err := ctl.query.LoginHandler(query); err != nil {
		context.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	token, err := ctl.command.CreateToken(context, dto.Id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, token)
}
