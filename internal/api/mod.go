package api

import (
	"TSM-Server/internal/service"
	"github.com/gin-gonic/gin"
)

func GetModInfo(c *gin.Context) {
	service := service.ModService{}
	response := service.GetModInfoService()
	c.JSON(200, response)
}

func ModAction(c *gin.Context) {
	service := service.ModService{}
	c.ShouldBindJSON(service)
	response := service.ModActionService()
	c.JSON(200, response)
}

func DelMod(c *gin.Context) {
	service := service.ModService{}
	c.ShouldBindJSON(service)
	response := service.DelModService()
	c.JSON(200, response)
}
