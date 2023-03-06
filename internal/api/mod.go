package api

import (
	"TSM-Server/internal/service"

	"github.com/gin-gonic/gin"
)

func GetModInfo(c *gin.Context) {
	modService := service.ModService{}
	response := modService.GetModInfoService()
	c.JSON(200, response)
}

func ModAction(c *gin.Context) {
	modService := service.ModService{}
	c.ShouldBind(&modService)
	response := modService.ModActionService()
	c.JSON(200, response)
}

func DelMod(c *gin.Context) {
	modService := service.ModService{}
	c.ShouldBind(&modService)
	response := modService.DelModService()
	c.JSON(200, response)
}

func GetModList(c *gin.Context) {
	modService := service.ModService{}
	response := modService.GetModListService()
	c.JSON(200, response)
}
