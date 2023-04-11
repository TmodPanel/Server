package service

import (
	"TSM-Server/cmd/process"
	"TSM-Server/internal/serializer"
	"TSM-Server/utils"
)

type PlayerService struct {
	Ip       string `json:"ip"`
	Id       int    `json:"id"`
	Nickname string `json:"nickname"`
}

type Player struct {
	Id       int    `json:"id"`
	Ip       string `json:"ip"`
	Ban      bool   `json:"ban"`
	Nickname string `json:"nickname"`
	JoinTime string `json:"join_time"`
	LeftTime string `json:"left_time"`
}

// GetPlayerInfoService wip
func (s *PlayerService) GetPlayerInfoService() serializer.Response {
	proc, err := process.Pool.GetWorker(s.Id)
	if err != nil {
		serializer.HandleErr(err, "获取进程失败")
	}

	list := proc.Ps
	result := make([]Player, 0)
	for k, v := range list {
		result = append(result, Player{
			Id:       k,
			Nickname: v.Name,
			Ip:       "",
			Ban:      false,
		})
	}

	return serializer.Response{
		Data: result,
		Msg:  "获取到玩家信息",
	}
}

func (s *PlayerService) KicPlayerService() serializer.Response {
	proc, err := process.Pool.GetWorker(s.Id)
	if err != nil {
		return serializer.HandleErr(err, "获取进程失败")
	}
	_, err = proc.Command(process.KICK + s.Nickname)
	if err != nil {
		return serializer.HandleErr(err, "踢出失败")
	}

	return serializer.Response{
		Msg: "踢出成功",
	}
}

func (s *PlayerService) BlockPlayerService() serializer.Response {
	proc, err := process.Pool.GetWorker(s.Id)
	if err != nil {
		return serializer.HandleErr(err, "获取进程失败")
	}
	_, err = proc.Command(process.BAN + s.Nickname)
	if err != nil {
		return serializer.HandleErr(err, "BAN失败")
	}

	return serializer.Response{
		Msg: "BAN成功",
	}
}

func (s *PlayerService) DelPlayerService() serializer.Response {
	//t.Ip,t.Nickname
	//打开ban list文件并删除

	if err := utils.RemoveFromBanList(s.Ip); err != nil {
		return serializer.HandleErr(err, "删除失败")
	}

	if err := utils.RemoveFromBanList(s.Nickname); err != nil {
		return serializer.HandleErr(err, "删除失败")
	}

	return serializer.Response{
		Msg: "删除成功",
	}
}
