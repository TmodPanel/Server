package api

import (
	"TSM-Server/internal/service"

	"github.com/gin-gonic/gin"
)

func GetPlayerInfo(c *gin.Context) {
	playerService := service.PlayerService{}
	response := playerService.GetPlayerInfoService()
	c.JSON(200, response)
}

func KicPlayer(c *gin.Context) {
	playerService := service.PlayerService{}
	c.ShouldBind(&playerService)
	response := playerService.KicPlayerService()
	c.JSON(200, response)
}

func BlockPlayer(c *gin.Context) {
	playerService := service.PlayerService{}
	c.ShouldBind(&playerService)
	response := playerService.BlockPlayerService()
	c.JSON(200, response)
}

func DelPlayer(c *gin.Context) {
	playerService := service.PlayerService{}
	c.ShouldBindJSON(&playerService)
	response := playerService.DelPlayerService()
	c.JSON(200, response)
}
