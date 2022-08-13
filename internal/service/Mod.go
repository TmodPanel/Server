package service

import (
	"TSM-Server/internal/serializer"
	"TSM-Server/utils"
)

type ModService struct {
	id     string
	name   string
	enable bool
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
			Data:  nil,
			Error: err.Error(),
			Msg:   "获取mod信息失败",
		}
	}
	i := 0
	for k, v := range mods {
		t := Mod{Id: string(i), Name: k, Enable: v}
		list = append(list, t)
		i++
	}
	return serializer.Response{
		Data:  list,
		Msg:   "获取到mod信息",
		Error: err.Error(),
	}
}

func (s *ModService) ModActionService() serializer.Response {
	if s.enable {
		if err := utils.EnableMod(s.name); err != nil {
			return serializer.Response{
				Data:  nil,
				Msg:   "获取到mod信息",
				Error: err.Error(),
			}
		}
	}
	return serializer.Response{
		Data:  nil,
		Msg:   "获取到mod信息",
		Error: "",
	}
}

func (s *ModService) DelModService() serializer.Response {
	err := utils.DelMod(s.name)
	return serializer.Response{
		Data:  nil,
		Msg:   "获取到mod信息",
		Error: err.Error(),
	}
}
