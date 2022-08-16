package api

import (
	"TSM-Server/internal/service"
	"github.com/gin-gonic/gin"
)

func GetPlayerInfo(c *gin.Context) {
	service := service.PlayerService{}
	response := service.GetPlayerInfoService()
	c.JSON(200, response)
}

func KicPlayer(c *gin.Context) {
	service := service.PlayerService{}
	c.ShouldBind(&service)
	response := service.KicPlayerService()
	c.JSON(200, response)
}

func BlockPlayer(c *gin.Context) {
	service := service.PlayerService{}
	c.ShouldBind(&service)
	response := service.BlockPlayerService()
	c.JSON(200, response)
}

func DelPlayer(c *gin.Context) {
	service := service.PlayerService{}
	c.ShouldBindJSON(&service)
	response := service.DelPlayerService()
	c.JSON(200, response)
}
