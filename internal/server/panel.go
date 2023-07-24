package server

import (
	"TSM-Server/cmd/setting"
	"TSM-Server/internal/serializer"
	"time"

	"github.com/shirou/gopsutil/v3/host"
)

type Panel struct {
	Version     string `json:"version"`
	Port        int    `json:"port"`
	Platform    string `json:"platform"` // ex: windows
	Arch        string `json:"Arch"`     // ex: x86_64
	HostName    string `json:"hostname"`
	Online      int    `json:"online"`
	RunningTime string `json:"runningTime"`
	Uptime      string `json:"uptime"`
}

func (s *Panel) GetPanelInfo() serializer.Response {
	info := Panel{}

	info.Version = setting.Cfg.Version
	info.Port = setting.Cfg.Http.Port

	hostInfo, _ := host.Info()

	info.Platform = hostInfo.Platform

	info.Arch = hostInfo.KernelArch

	info.HostName = hostInfo.Hostname

	info.Online = 0

	info.RunningTime = time.Now().Format("2006-01-02 15:04:05")

	info.Uptime = "0"

	return serializer.Response{
		Data: info,
		Msg:  "返回服务器基本信息",
	}
}
