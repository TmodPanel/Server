package server

import (
	process2 "TSM-Server/internal/process"
	"TSM-Server/internal/serializer"
	"encoding/json"
	"os"
)

type Player struct {
	Id           int    `json:"id"`
	Ip           string `json:"ip"`
	Ban          bool   `json:"ban"`
	Nickname     string `json:"nickname"`
	RegisterTime string `json:"registerTime"`
	LastLogin    string `json:"lastLogin"`
}

func (s *Player) GetPlayerInfo() serializer.Response {
	list, err := readPlayer()
	if err != nil {
		return serializer.HandleErr(err, "获取玩家列表失败")
	}

	return serializer.Response{
		Data: list,
		Msg:  "获取玩家列表成功",
	}
}

func (s *Player) KicPlayer() serializer.Response {
	proc, err := process2.Pool.GetWorker(s.Id)
	if err != nil {
		return serializer.HandleErr(err, "获取进程失败")
	}
	_, err = proc.Command(process2.KICK + s.Nickname)
	if err != nil {
		return serializer.HandleErr(err, "踢出失败")
	}

	return serializer.Response{
		Msg: "踢出成功",
	}
}

func (s *Player) BlockPlayer() serializer.Response {
	proc, err := process2.Pool.GetWorker(s.Id)
	if err != nil {
		return serializer.HandleErr(err, "获取进程失败")
	}
	_, err = proc.Command(process2.BAN + s.Nickname)
	if err != nil {
		return serializer.HandleErr(err, "BAN失败")
	}

	return serializer.Response{
		Msg: "BAN成功",
	}
}

//func (s *Player) DelPlayer() serializer.Response {
//	//t.Ip,t.Nickname
//	//打开ban list文件并删除
//
//	if err := utils.RemoveFromBanList(s.Ip); err != nil {
//		return serializer.HandleErr(err, "删除失败")
//	}
//
//	if err := utils.RemoveFromBanList(s.Nickname); err != nil {
//		return serializer.HandleErr(err, "删除失败")
//	}
//
//	return serializer.Response{
//		Msg: "删除成功",
//	}
//}

//func removeFromBanList(name string) error {
//	lines, err := read("./config/banlist.txt")
//	if err != nil {
//		return err
//	}
//	return write("./config/banlist.txt", lines, name)
//}

func readPlayer() ([]Player, error) {
	catalog, err := os.ReadDir("./data/player/")

	if err != nil {
		return nil, err
	}
	var players []Player
	for _, v := range catalog {
		// 读取文件内容
		file, err := os.ReadFile("./data/player/" + v.Name())
		if err != nil {
			return nil, err
		}
		var player Player
		json.Unmarshal(file, &player)
		players = append(players, player)
	}
	return players, nil
}
