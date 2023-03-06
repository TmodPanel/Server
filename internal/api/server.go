package api

import (
	"TSM-Server/internal/service"

	"github.com/gin-gonic/gin"
)

func GetServerInfo(c *gin.Context) {
	serverService := service.ServerService{}
	response := serverService.GetServerInfoService()
	c.JSON(200, response)
}
