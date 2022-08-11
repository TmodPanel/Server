package service

import (
	"TSM-Server/cmd/tmd"
	"TSM-Server/utils"
	"github.com/gin-gonic/gin"
	"log"
)

type PlayerService struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
	Ip       string `json:"ip"`
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
	var t PlayerService
	c.BindJSON(t)

	//t.Ip,t.Nickname
	//打开ban list文件并删除
	if err := utils.RemoveFromBanList(t.Ip); err != nil {
		log.Println(err)
	}
	if err := utils.RemoveFromBanList(t.Nickname); err != nil {
		log.Println(err)
	}

	c.JSON(200, gin.H{"msg": "success"})
}
