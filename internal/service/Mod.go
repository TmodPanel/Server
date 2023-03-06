package service

import (
	"TSM-Server/internal/serializer"
	"TSM-Server/utils"
	"strconv"
)

type ModService struct {
	Page   int    `json:"page"`
	Name   string `json:"name"`
	Action int    `json:"action"`
}

type Mod struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Enable bool   `json:"Enable"`
}

func (s *ModService) GetModInfoService() serializer.Response {
	var list []Mod
	mods, err := utils.GetModInfo()
	if err != nil {
		return serializer.HandleErr(err, "获取mod信息失败")
	}
	id := 1
	for k, v := range mods {
		t := Mod{Id: strconv.Itoa(id), Name: k, Enable: v}
		list = append(list, t)
		id++
	}
	return serializer.Response{
		Data: list,
		Msg:  "获取到mod信息",
	}
}

func (s *ModService) ModActionService() serializer.Response {
	if s.Action == 1 {
		if err := utils.EnableMod(s.Name); err != nil {
			return serializer.HandleErr(err, "启用失败")
		}
	} else if s.Action == 2 {
		if err := utils.RemoveFromEnable(s.Name); err != nil {
			return serializer.HandleErr(err, "禁用失败")
		}
	} else if s.Action == 3 {
		if err := utils.DelMod(s.Name); err != nil {
			return serializer.HandleErr(err, "删除失败")
		}
	}
	return serializer.Response{
		Msg: "操作成功",
	}
}

func (s *ModService) DelModService() serializer.Response {
	if err := utils.DelMod(s.Name); err != nil {
		return serializer.HandleErr(err, "删除失败")
	}
	return serializer.Response{
		Msg: "删除成功",
	}
}

func (s *ModService) GetModListService() serializer.Response {
	list, err := utils.GetModList(s.Page)
	if err != nil {
		return serializer.HandleErr(err, "获取失败")
	}
	return serializer.Response{
		Data: list,
		Msg:  "获取成功",
	}
}
