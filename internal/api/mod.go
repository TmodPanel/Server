package api

import (
	"TSM-Server/internal/service"
	"TSM-Server/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func GetModInfo(c *gin.Context) {
	service := service.ModService{}
	list := service.GetModInfoService()
	c.JSON(200, list)

}

func ModAction(c *gin.Context) {
	service := service.ModService{}
	c.ShouldBindJSON(service)
	service.ModAction()
	c.JSON(200, nil)
}

func DelMod(c *gin.Context) {
	var t ModService
	c.ShouldBindJSON(t)
	//找到mod文件位置并且删除
	if err := utils.DelMod(t.Name); err != nil {
		log.Println(err)
	}
	c.JSON(200, gin.H{"msg": "mod已删除"})
}
