package api

import (
	"TSM-Server/internal/server"

	"github.com/gin-gonic/gin"
)

func AddInstance(c *gin.Context) {
	instanceName := c.Query("instance_name")
	response := server.AddInstance(instanceName)
	c.JSON(200, response)
}

func GetInstanceInfo(c *gin.Context) {
	instanceName := c.Query("instance_name")
	response := server.GetInstanceInfo(instanceName)
	c.JSON(200, response)
}

func DelInstance(c *gin.Context) {
	instanceName := c.Query("instance_name")
	response := server.DelInstance(instanceName)
	c.JSON(200, response)
}

func RestartGame(c *gin.Context) {
	instanceName := c.Query("instance_name")
	response := server.RestartGame(instanceName)
	c.JSON(200, response)
}

func StartGame(c *gin.Context) {
	instanceName := c.Query("instance_name")
	response := server.StartGame(instanceName)
	c.JSON(200, response)
}

func StopGame(c *gin.Context) {
	instanceName := c.Query("instance_name")
	response := server.StopGame(instanceName)
	c.JSON(200, response)
}

func SetGameTime(c *gin.Context) {
	instanceName := c.Query("instance_name")
	gameTime := c.Query("game_time")
	response := server.SetGameTime(instanceName, gameTime)
	c.JSON(200, response)
}
