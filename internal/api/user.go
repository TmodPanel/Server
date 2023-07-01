package api

import (
	"TSM-Server/internal/modle"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var data modle.User
	c.ShouldBindJSON(&data)
	response := data.Login()
	c.JSON(200, response)
}

func Register(c *gin.Context) {
	var data modle.User
	c.ShouldBindJSON(&data)
	response := data.Register()
	c.JSON(200, response)
}
