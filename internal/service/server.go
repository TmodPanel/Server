package service

import (
	"TSM-Server/internal/serializer"
	"TSM-Server/utils"
	"runtime"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

type ServerService struct {
	Id int `json:"id"` // according id to the scheme
}

type Server struct {
	Cpu struct {
		Percent float64 `json:"percent"`
	}
	Mem struct {
		Total int `json:"total"`
		Used  int `json:"used"`
	}
	Disk struct {
		Total int `json:"total"`
		Used  int `json:"used"`
	}
	Ip string `json:"ip"`

	Os              string `json:"os"`              // ex: linux
	PlatformVersion string `json:"platformVersion"` // ex: Ubuntu 20.04
	PlatformArch    string `json:"platformArch"`    // ex: x86_64
}

func (s *ServerService) GetServerInfoService() serializer.Response {
	var info Server
	info.Ip = utils.IpAddress()
	memory, _ := mem.VirtualMemory()
	info.Mem.Total = int(memory.Total / 1024 / 1024)
	info.Mem.Used = int(memory.Used / 1024 / 1024)
	diskInfo, _ := disk.Usage("/")
	info.Disk.Total = int(diskInfo.Total / 1024 / 1024)
	info.Disk.Used = int(diskInfo.Used / 1024 / 1024)
	percent, _ := cpu.Percent(0, false)
	info.Cpu.Percent = percent[0]
	info.Os = runtime.GOOS
	hostInfo, _ := host.Info()
	info.PlatformVersion = hostInfo.PlatformVersion
	info.PlatformArch = hostInfo.KernelArch
	return serializer.Response{
		Data:   info,
		Msg:    "返回服务器基本信息",
		Status: 200,
	}
}
