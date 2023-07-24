package api

import (
	"TSM-Server/internal/server"

	"github.com/gin-gonic/gin"
)

func GetPanelInfo(c *gin.Context) {
	panel := server.Panel{}
	response := panel.GetPanelInfo()
	c.JSON(200, response)
}
