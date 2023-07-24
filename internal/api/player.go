package api

import (
	"TSM-Server/internal/server"

	"github.com/gin-gonic/gin"
)

func GetPlayerInfo(c *gin.Context) {
	player := server.Player{}
	response := player.GetPlayerInfo()
	c.JSON(200, response)
}

func KickPlayer(c *gin.Context) {
	player := server.Player{}
	c.ShouldBind(&player)
	response := player.KicPlayer()
	c.JSON(200, response)
}

func BlockPlayer(c *gin.Context) {
	player := server.Player{}
	c.ShouldBind(&player)
	response := player.BlockPlayer()
	c.JSON(200, response)
}

//func DelPlayer(c *gin.Context) {
//	player := server.Player{}
//	c.ShouldBindJSON(&player)
//	response := player.DelPlayer()
//	c.JSON(200, response)
//}
