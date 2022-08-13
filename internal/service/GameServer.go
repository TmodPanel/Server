package service

import (
	"TSM-Server/cmd/tmd"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type GameServerService struct {
	Ip            string `json:"ip"`
	Online        string `json:"online"`
	Password      string `json:"password"`
	Port          string `json:"port"`
	World         string `json:"world"`
	Seed          string `json:"seed"`
	Time          string `json:"time"`
	ServerVersion string `json:"serverVersion"`
	ClientVersion string `json:"clientVersion"`
}

func GetServerInfo(c *gin.Context) {
	var info GameServerService
	//tmd.Command("version")
	resp, err := http.Get("https://myexternalip.com/raw")
	if err != nil {
		info.Ip = ""
	} else {
		body, _ := io.ReadAll(resp.Body)
		info.Ip = string(body)
	}

	c.JSON(200, info)
}

func SetTime(c *gin.Context) {
	t := new(struct {
		time string
	})
	c.ShouldBindJSON(t)
	//dawn noon dusk midnight
	tmd.Command(t.time)
	c.JSON(200, gin.H{"msg": "设置成功"})
}

func ServerAction(c *gin.Context) {
	t := new(struct {
		action string
	})
	c.ShouldBindJSON(t)
	//exit-nosave
	//save
	//start restart exit
	if t.action != "restart" {

	} else if t.action == "start" {

	} else {
		tmd.Command("exit")
	}

	c.JSON(200, gin.H{"msg": "设置成功"})
}
