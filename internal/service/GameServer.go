package service

import "github.com/gin-gonic/gin"

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

}

func SetTime(c *gin.Context) {

}

func ServerAction(c *gin.Context) {

}
