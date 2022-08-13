package service

import (
	"TSM-Server/cmd/tmd"
	"TSM-Server/internal/serializer"
	"TSM-Server/utils"
	"io"
	"net/http"
	"strings"
)

type ServerService struct {
	time   string
	action string
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
	resp, err := http.Get("https://myexternalip.com/raw")
	if err != nil {
		info.Ip = ""
	} else {
		body, _ := io.ReadAll(resp.Body)
		info.Ip = string(body)
	}
	info.Seed = tmd.Command("seed")
	info.Port = tmd.Command("port")
	//wip
	info.Online = tmd.Command("playing")
	info.Password = tmd.Command("password")
	info.Time = tmd.Command("time")
	info.World, _ = utils.ReadServerConfig("worldname")
	version := tmd.Command("version")
	info.ClientVersion = strings.Split(version, "-")[0]
	info.ServerVersion = strings.TrimSpace(strings.Split(version, "-")[1])

	return serializer.Response{
		Data:  info,
		Msg:   "返回服务器基本信息",
		Error: "",
	}
}

func (s *ServerService) SetTimeService() serializer.Response {
	//dawn noon dusk midnight
	tmd.Command(s.time)
	return serializer.Response{
		Data:  "",
		Msg:   "设置时间" + s.time,
		Error: "",
	}
}

// ServerActionService  wip
func (s *ServerService) ServerActionService() serializer.Response {
	//restart
	//exit
	//save
	if s.action == "exit" {
		tmd.Command("exit")
	} else if s.action == "start" {
		start := make(chan bool)
		tmd.Start(start)
	}
	return serializer.Response{
		Data:  "",
		Msg:   "action success",
		Error: "",
	}
}
