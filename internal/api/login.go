package api

import (
	"TSM-Server/internal/service"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	loginService := service.LoginService{}
	c.ShouldBind(loginService)
	response := loginService.Login()
	c.JSON(200, response)
}
