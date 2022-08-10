package service

import (
	"TSM-Server/cmd/tmd"
	"github.com/gin-gonic/gin"
)

type PlayerService struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
	Ban      bool   `json:"ban"`
}

func GetPlayerInfo(c *gin.Context) {
	var list []PlayerService
	tmd.Command("playing")
	c.JSON(200, list)
}

func KicPlayer(c *gin.Context) {
	var t PlayerService
	c.BindJSON(t)
	tmd.Command("kick " + t.Nickname)
	c.JSON(200, gin.H{"msg": t.Nickname + "has left"})
}

func BlockPlayer(c *gin.Context) {
	var t PlayerService
	c.BindJSON(t)
	tmd.Command("ban " + t.Nickname)
	c.JSON(200, gin.H{"msg": t.Nickname + "has left"})
}

func DelPlayer(c *gin.Context) {
	//打开banlist文件并删除
	c.JSON(200, gin.H{"msg": "success"})
}
