package api

import (
	"TSM-Server/internal/modle"

	"github.com/gin-gonic/gin"
)

func GetPanelInfo(c *gin.Context) {
	serverService := modle.ServerService{}
	response := serverService.GetServerInfoService()
	c.JSON(200, response)
}
