package api

import (
	"TSM-Server/internal/service"

	"github.com/gin-gonic/gin"
)

func GetServerInfo(c *gin.Context) {
	service := service.ServerService{}
	response := service.GetServerInfoService()
	c.JSON(200, response)
}
