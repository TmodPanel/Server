package api

import (
	"TSM-Server/internal/server"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var data server.User
	c.ShouldBindJSON(&data)
	response := data.Login()
	c.JSON(200, response)
}

func Register(c *gin.Context) {
	var data server.User
	c.ShouldBindJSON(&data)
	response := data.Register()
	c.JSON(200, response)
}

func Logout(c *gin.Context) {
	var data server.User
	c.ShouldBindJSON(&data)
	response := data.Logout()
	c.JSON(200, response)
}
