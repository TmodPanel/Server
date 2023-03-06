package api

import (
	"TSM-Server/internal/service"

	"github.com/gin-gonic/gin"
)

func GetSchemesInfo(c *gin.Context) {
	schemeService := service.SchemeService{}
	response := schemeService.GetSchemesInfoService()
	c.JSON(200, response)
}

func AddScheme(c *gin.Context) {
	schemeService := service.SchemeService{}
	c.ShouldBindJSON(schemeService)
	response := schemeService.AddSchemeService()
	c.JSON(200, response)
}

func UpdateScheme(c *gin.Context) {
	schemeService := service.SchemeService{}
	c.ShouldBindJSON(schemeService)
	response := schemeService.UpdateSchemeService()
	c.JSON(200, response)
}

func DelScheme(c *gin.Context) {
	schemeService := service.SchemeService{}
	c.ShouldBindJSON(schemeService)
	response := schemeService.DelSchemeService()
	c.JSON(200, response)
}

func ResetScheme(c *gin.Context) {
	schemeService := service.SchemeService{}
	c.ShouldBindJSON(schemeService)
	response := schemeService.ResetSchemeService()
	c.JSON(200, response)
}
