package service

import (
	"TSM-Server/cmd/tmd"
	"TSM-Server/internal/serializer"
	"TSM-Server/utils"
	"strings"
)

type ServerService struct {
	Time   string `json:"time"`
	Action string `json:"action"`
	Config string `json:"config"`
}

type Server struct {
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

// GetServerInfoService  wip
func (s *ServerService) GetServerInfoService() serializer.Response {
	var info Server
	info.Ip = utils.IpAddress()
	info.Seed = tmd.Command("seed")
	info.Online = tmd.Command("playing")
	info.Port, _ = utils.ReadServerConfig("port")
	info.Password, _ = utils.ReadServerConfig("password")
	info.World, _ = utils.ReadServerConfig("worldname")
	info.Time = tmd.Command("time")
	version := tmd.Command("version")
	if version == "game not start!" {
		info.ServerVersion = ""
		info.ClientVersion = ""
	} else {
		info.ClientVersion = strings.TrimSpace(strings.Split(version, "-")[0])
		info.ServerVersion = strings.TrimSpace(strings.Split(version, "-")[1])
	}

	return serializer.Response{
		Data: info,
		Msg:  "返回服务器基本信息",
	}
}

func (s *ServerService) SetTimeService() serializer.Response {
	//dawn noon dusk midnight
	tmd.Command(s.Time)
	return serializer.Response{
		Msg: "设置时间" + s.Time,
	}
}

// ServerActionService 错误处理
func (s *ServerService) ServerActionService() serializer.Response {
	ok := tmd.CheckStart()
	response := serializer.Response{
		Msg: "server has " + s.Action,
	}
	if ok {
		switch s.Action {
		case "start":
			break
		case "exit":
			tmd.Command("exit")
			break
		case "restart":
			tmd.Command("exit")
			ch := make(chan bool)
			go tmd.Start(ch, s.Config)
			<-ch
			break
		}
	} else {
		switch s.Action {
		case "start", "restart":
			ch := make(chan bool)
			go tmd.Start(ch, s.Config)
			<-ch
			break
		case "exit":
			break
		}
	}
	return response
}
