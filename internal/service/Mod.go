package service

import "github.com/gin-gonic/gin"

type ModService struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	IsEnable bool   `json:"isEnable"`
}

func GetModInfo(c *gin.Context) {
	var list []ModService
	//查看enable.json文件信息
	c.JSON(200, list)

}

func ModAction(c *gin.Context) {
	t := new(struct {
		action string
	})
	c.ShouldBindJSON(t)

	//启用

	//禁用
	c.JSON(200, nil)
}

func DelMod(c *gin.Context) {

	//找到mod文件位置并且删除

	c.JSON(200, gin.H{"msg": "mod已删除"})
}
