package api

import (
	service2 "TSM-Server/internal/service"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	service := service2.LoginService{}
	c.ShouldBind(service)
	response := service.Login()
	c.JSON(200, response)
}
