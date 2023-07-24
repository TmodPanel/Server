package api

import (
	"TSM-Server/internal/server"

	"github.com/gin-gonic/gin"
)

func GetModList(c *gin.Context) {
	instanceName := c.Query("instance_name")
	response := server.GetModInfo(instanceName)
	c.JSON(200, response)
}

func EnableMod(c *gin.Context) {
	modService := server.Mod{}
	modService.ID = c.Param("id")
	instanceName := c.Query("instance_name")
	response := modService.EnableMod(instanceName)
	c.JSON(200, response)
}

func DisableMod(c *gin.Context) {
	modService := server.Mod{}
	modService.ID = c.Param("id")
	instanceName := c.Query("instance_name")
	response := modService.DisableMod(instanceName)
	c.JSON(200, response)
}

func DelMod(c *gin.Context) {
	modService := server.Mod{}
	modService.ID = c.Param("id")
	response := modService.DelMod()
	c.JSON(200, response)
}
