package api

import (
	"TSM-Server/internal/service"

	"github.com/gin-gonic/gin"
)

func GetGameInfo(c *gin.Context) {
	service := service.ServerService{}
	response := service.GetGameInfoService()
	c.JSON(200, response)
}

func SetTime(c *gin.Context) {
	service := service.ServerService{}
	c.ShouldBind(&service)
	response := service.SetTimeService()
	c.JSON(200, response)
}

func ServerAction(c *gin.Context) {
	service := service.ServerService{}
	c.ShouldBind(&service)
	response := service.ServerActionService()
	c.JSON(200, response)
}
