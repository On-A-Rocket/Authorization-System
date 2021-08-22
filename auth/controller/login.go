package controller

import (
	"net/http"

	"github.com/On-A-Rocket/Authorization-System/auth/domain/dto"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
}

func newLoginController() *LoginController {
	return &LoginController{}
}

// func (ctl *LoginController) loginRouting(router *gin.Engine) {
// 	router.GET("/login")
// }

func (ctl *LoginController) Login(context *gin.Context) {
	dto := &dto.Login{}
	if err := context.ShouldBind(&dto); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

}
