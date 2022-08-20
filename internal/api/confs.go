package api

import (
	"TSM-Server/internal/service"
	"github.com/gin-gonic/gin"
)

func GetSchemesInfo(c *gin.Context) {
	service := service.SchemeService{}
	response := service.GetSchemesInfoService()
	c.JSON(200, response)
}

func AddScheme(c *gin.Context) {
	service := service.SchemeService{}
	c.ShouldBindJSON(service)
	response := service.AddSchemeService()
	c.JSON(200, response)
}

func UpdateScheme(c *gin.Context) {
	service := service.SchemeService{}
	c.ShouldBindJSON(service)
	response := service.UpdateSchemeService()
	c.JSON(200, response)
}

func DelScheme(c *gin.Context) {
	service := service.SchemeService{}
	c.ShouldBindJSON(service)
	response := service.DelSchemeService()
	c.JSON(200, response)
}

func ResetScheme(c *gin.Context) {
	service := service.SchemeService{}
	c.ShouldBindJSON(service)
	response := service.ResetSchemeService()
	c.JSON(200, response)
}
