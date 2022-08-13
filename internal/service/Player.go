package service

import (
	"TSM-Server/cmd/tmd"
	"TSM-Server/internal/serializer"
	"TSM-Server/utils"
)

type PlayerService struct {
	nickname string
	ip       string
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
		Data:  "",
		Msg:   result,
		Error: "",
	}
}

func (s *PlayerService) KicPlayerService() serializer.Response {
	result := tmd.Command("kick " + s.nickname)
	return serializer.Response{
		Data:  "",
		Msg:   result,
		Error: "",
	}
}

func (s *PlayerService) BlockPlayerService() serializer.Response {
	result := tmd.Command("ban " + s.nickname)
	return serializer.Response{
		Data:  "",
		Msg:   result,
		Error: "",
	}
}

func (s *PlayerService) DelPlayerService() serializer.Response {
	//t.Ip,t.Nickname
	//打开ban list文件并删除
	err1 := utils.RemoveFromBanList(s.ip)
	err2 := utils.RemoveFromBanList(s.nickname)
	return serializer.Response{
		Data:  "",
		Msg:   "已删除",
		Error: err1.Error() + err2.Error(),
	}
}
