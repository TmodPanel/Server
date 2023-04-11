package service

import (
	"TSM-Server/cmd/process"
	"TSM-Server/internal/serializer"
	"TSM-Server/utils"
)

type GameService struct {
	Id     int    `json:"id"`
	Time   string `json:"time"`
	Action string `json:"action"`
	Config string `json:"config"`
}

type Game struct {
	Ip            string `json:"ip"`
	Online        string `json:"online"`
	Password      string `json:"password"`
	Port          string `json:"port"`
	World         string `json:"world"`
	Seed          string `json:"seed"`
	Time          string `json:"time"`
	ServerVersion string `json:"serverVersion"`
	ClientVersion string `json:"clientVersion"`
}

const (
	STOP    = "stop"
	START   = "start"
	RESTART = "restart"
)

// GetGameInfoService   wip
func (s *GameService) GetGameInfoService() serializer.Response {
	proc, err := process.Pool.GetWorker(s.Id)
	if err != nil {
		return serializer.HandleErr(err, "获取进程失败")
	}

	var info Game
	info.Ip = utils.IpAddress()
	info.Seed, _ = proc.Command(process.SEED)
	info.Online, _ = proc.Command(process.PLAYING)
	info.Port, _ = utils.ReadServerConfig(process.PORT)
	info.Password, _ = utils.ReadServerConfig(process.PASSWORD)
	info.World, _ = utils.ReadServerConfig(process.WORLD)
	info.Time, _ = proc.Command(process.TIME)
	version, _ := proc.Command(process.VERSION)
	info.ServerVersion = version
	info.ClientVersion = version

	return serializer.Response{
		Data: info,
		Msg:  "返回服务器基本信息",
	}
}

func (s *GameService) SetTimeService() serializer.Response {
	proc, err := process.Pool.GetWorker(s.Id)
	if err != nil {
		return serializer.HandleErr(err, "获取进程失败")
	}

	command, err := proc.Command(s.Time)
	return serializer.Response{
		Msg:   command,
		Error: err.Error(),
	}
}

// ServerActionService 错误处理
func (s *GameService) ServerActionService() serializer.Response {

	switch s.Action {
	case STOP:
		if err := process.Pool.Stop(s.Id); err != nil {
			return serializer.HandleErr(err, "停止失败")
		}
		break
	case START:
		if err := process.Pool.Run(s.Id); err != nil {
			return serializer.HandleErr(err, "启动失败")
		}
		break
	case RESTART:
		if err := process.Pool.Restart(s.Id); err != nil {
			return serializer.HandleErr(err, "重启失败")
		}
		break
	}
	return serializer.Response{
		Msg: "执行命令" + s.Action,
	}
}
