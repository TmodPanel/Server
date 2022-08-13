package service

import (
	"TSM-Server/utils"
	"github.com/gin-gonic/gin"
	"log"
)

type ModService struct {
}

type Mod struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Enable bool   `json:"Enable"`
}

func (s *ModService) GetModInfoService() []Mod {
	var list []Mod
	mods, err := utils.GetModInfo()
	if err != nil {
		log.Println(err)
	}
	i := 0
	for k, v := range mods {
		t := Mod{Id: string(i), Name: k, Enable: v}
		list = append(list, t)
		i++
	}
	return list
}

func (s *ModService) ModAction() {
	var t Mod

	if t.Enable {
		if err := utils.EnableMod(t.Name); err != nil {
			log.Println(err)
		}
	} else {
		if err := utils.RemoveFromEnable(t.Name); err != nil {
			log.Println(err)
		}
	}
	c.JSON(200, nil)
}

func (s *ModService) DelMod() {
	var t Mod
	c.ShouldBindJSON(t)
	//找到mod文件位置并且删除
	if err := utils.DelMod(t.Name); err != nil {
		log.Println(err)
	}
	c.JSON(200, gin.H{"msg": "mod已删除"})
}
