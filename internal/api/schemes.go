package api

import (
	"TSM-Server/internal/server"

	"github.com/gin-gonic/gin"
)

func GetSchemesInfo(c *gin.Context) {
	instanceName := c.Query("instance_name")
	response := server.GetSchemeInfo(instanceName)
	c.JSON(200, response)
}

func UpdateScheme(c *gin.Context) {
	var schemeService server.Scheme
	instanceName := c.Query("instance_name")
	c.ShouldBindJSON(&schemeService)
	response := schemeService.UpdateScheme(instanceName)
	c.JSON(200, response)
}
func DelScheme(c *gin.Context) {
	instanceName := c.Query("instance_name")
	response := server.DelScheme(instanceName)
	c.JSON(200, response)
}
