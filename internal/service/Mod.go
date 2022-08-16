package service

import (
	"TSM-Server/internal/serializer"
	"TSM-Server/utils"
	"strconv"
)

type ModService struct {
	Name   string `json:"name"`
	Enable bool   `json:"enable"`
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
		return serializer.Response{
			Error: utils.ErrToString(err),
			Msg:   "获取mod信息失败",
		}
	}
	i := 1
	for k, v := range mods {
		t := Mod{Id: strconv.Itoa(i), Name: k, Enable: v}
		list = append(list, t)
		i++
	}
	return serializer.Response{
		Data: list,
		Msg:  "获取到mod信息",
	}
}

func (s *ModService) ModActionService() serializer.Response {
	if s.Enable {
		err := utils.EnableMod(s.Name)
		return serializer.Response{
			Msg:   "已启用" + s.Name,
			Error: utils.ErrToString(err),
		}
	}
	err := utils.RemoveFromEnable(s.Name)
	return serializer.Response{
		Msg:   "已禁用" + s.Name,
		Error: utils.ErrToString(err),
	}
}

func (s *ModService) DelModService() serializer.Response {
	err := utils.DelMod(s.Name)
	return serializer.Response{
		Msg:   "已删除" + s.Name,
		Error: utils.ErrToString(err),
	}
}
