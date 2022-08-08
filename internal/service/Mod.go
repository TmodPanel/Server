package service

import "github.com/gin-gonic/gin"

type ModService struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	IsEnable bool   `json:"isEnable"`
}

func GetModInfo(c *gin.Context) {

}

func ModAction(c *gin.Context) {

}

func DelMod(c *gin.Context) {

}
