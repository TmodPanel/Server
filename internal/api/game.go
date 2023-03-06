package api

import (
	"TSM-Server/internal/service"

	"github.com/gin-gonic/gin"
)

func GetGameInfo(c *gin.Context) {
	gameService := service.ServerService{}
	response := gameService.GetGameInfoService()
	c.JSON(200, response)
}

func SetTime(c *gin.Context) {
	gameService := service.ServerService{}
	c.ShouldBind(&gameService)
	response := gameService.SetTimeService()
	c.JSON(200, response)
}

func ServerAction(c *gin.Context) {
	gameService := service.ServerService{}
	c.ShouldBind(&gameService)
	response := gameService.ServerActionService()
	c.JSON(200, response)
}
