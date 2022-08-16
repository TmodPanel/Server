package service

import (
	"TSM-Server/cmd/tmd"
	"TSM-Server/internal/serializer"
	"TSM-Server/utils"
)

type PlayerService struct {
	Nickname string `json:"nickname"`
	Ip       string `json:"ip"`
}

type Player struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
	Ip       string `json:"ip"`
	Ban      bool   `json:"ban"`
}

// GetPlayerInfoService wip
func (s *PlayerService) GetPlayerInfoService() serializer.Response {
	result := tmd.Command("playing")
	return serializer.Response{
		Msg: result,
	}
}

func (s *PlayerService) KicPlayerService() serializer.Response {
	result := tmd.Command("kick " + s.Nickname)
	return serializer.Response{
		Msg: result,
	}
}

func (s *PlayerService) BlockPlayerService() serializer.Response {
	result := tmd.Command("ban " + s.Nickname)
	return serializer.Response{
		Msg: result,
	}
}

func (s *PlayerService) DelPlayerService() serializer.Response {
	//t.Ip,t.Nickname
	//打开ban list文件并删除
	err1 := utils.RemoveFromBanList(s.Ip)
	err2 := utils.RemoveFromBanList(s.Nickname)
	return serializer.Response{
		Data:  "",
		Msg:   "已删除",
		Error: utils.ErrToString(err1) + utils.ErrToString(err2),
	}
}
