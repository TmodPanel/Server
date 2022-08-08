package service

import "github.com/gin-gonic/gin"

type PlayerService struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
	Ban      bool   `json:"ban"`
}

func GetPlayerInfo(c *gin.Context) {

}

func KicPlayer(c *gin.Context) {

}

func BlockPlayer(c *gin.Context) {

}

func DelPlayer(c *gin.Context) {

}
