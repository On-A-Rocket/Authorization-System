package controller

import (
	"net/http"

	"github.com/On-A-Rocket/Authorization-System/auth/application/query"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	queryHandler query.LoginQueryHandler
}

func newLoginController(
	queryHandler query.LoginQueryHandler) *LoginController {
	return &LoginController{queryHandler: queryHandler}
}

func (ctl *LoginController) loginRouting(router *gin.Engine) {
	router.GET("/login", ctl.Login)
}

// @Summary Login
// @Description Login
// @Tags Login
// @Produce json
// @Router /login [get]
// @Param id query string false "Id"
// @Param password query string false "Password"
// @Success 200 {string} login
func (ctl *LoginController) Login(context *gin.Context) {
	id := context.Query("id")

	query := query.LoginQuery{
		Id:       id,
		Password: "1",
	}

	if err := ctl.queryHandler.LoginHandler(context, query); err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}
	context.JSON(http.StatusOK, "login")
}
