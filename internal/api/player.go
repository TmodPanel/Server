package api

import (
	"TSM-Server/internal/service"
	"github.com/gin-gonic/gin"
)

func GetPlayerInfo(c *gin.Context) {
	server := service.PlayerService{}
	response := server.GetPlayerInfoService()
	c.JSON(200, response)
}

func KicPlayer(c *gin.Context) {
	server := service.PlayerService{}
	c.ShouldBindJSON(server)
	response := server.KicPlayerService()
	c.JSON(200, response)
}

func BlockPlayer(c *gin.Context) {
	server := service.PlayerService{}
	c.ShouldBindJSON(server)
	response := server.BlockPlayerService()
	c.JSON(200, response)
}

func DelPlayer(c *gin.Context) {
	server := service.PlayerService{}
	c.ShouldBindJSON(server)
	response := server.DelPlayerService()
	c.JSON(200, response)
}
