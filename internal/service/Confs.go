package service

import (
	"TSM-Server/internal/serializer"
	"TSM-Server/utils"
	"encoding/json"
	"fmt"
	"os"
)

type SchemeService struct {
	conf Scheme
}

type Scheme struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Maxnum    string `json:"maxnum"`
	Password  string `json:"password"`
	Port      string `json:"port"`
	Map       Map    `json:"map"`
	Language  string `json:"language"`
	Motd      string `json:"motd"`
	Priority  int    `json:"priority"`
	IP        string `json:"ip"`
	Npcstream int    `json:"npcstream"`
	Secure    bool   `json:"secure"`
	Banlist   bool   `json:"banlist"`
	Upnp      bool   `json:"upnp"`
}
type Map struct {
	Isnew     bool   `json:"isnew"`
	Size      string `json:"size"`
	Difficult int    `json:"difficult"`
	Seed      string `json:"seed"`
	Worldname string `json:"worldname"`
}

var (
	defaultScheme = Scheme{
		ID:       "1",
		Name:     "默认配置方案",
		Maxnum:   "8",
		Password: "",
		Port:     "7777",
		Map: Map{
			Isnew:     false,
			Size:      "Large",
			Difficult: 0,
			Seed:      "",
			Worldname: "Terraria",
		},
		Language:  "zh-Hans",
		Motd:      "Welcome to Terraria",
		Priority:  1,
		IP:        "",
		Npcstream: 0,
		Secure:    false,
		Banlist:   false,
		Upnp:      false,
	}
)

func (s SchemeService) GetSchemesInfoService() serializer.Response {
	var list []Scheme
	data, _ := os.ReadFile("./config/schemes/scheme.json")
	json.Unmarshal(data, &list)
	return serializer.Response{
		Data: list,
		Msg:  "获取到配置方案信息",
	}
}

func (s SchemeService) DelSchemeService() serializer.Response {
	var list []Scheme
	data, _ := os.ReadFile("./config/schemes/scheme.json")
	json.Unmarshal(data, &list)
	var newList []Scheme
	for i := 0; i < len(list); i++ {
		if list[i].ID == s.conf.ID {
			continue
		}
		newList = append(newList, list[i])
	}
	data, _ = json.Marshal(newList)
	os.WriteFile("./config/schemes/scheme.json", data, 0777)
	return serializer.Response{
		Msg: s.conf.Name + "已删除",
	}
}

func (s SchemeService) AddSchemeService() serializer.Response {
	var list []Scheme
	data, _ := os.ReadFile("./config/schemes/scheme.json")
	json.Unmarshal(data, &list)
	list = append(list, s.conf)
	data, _ = json.Marshal(list)
	os.WriteFile("./config/schemes/scheme.json", data, 0777)
	return serializer.Response{
		Msg: s.conf.Name + "已添加",
	}
}

func (s SchemeService) UpdateSchemeService() serializer.Response {
	var list []Scheme
	data, _ := os.ReadFile("./config/schemes/scheme.json")
	json.Unmarshal(data, &list)
	var newList []Scheme
	for i := 0; i < len(list); i++ {
		if list[i].ID == s.conf.ID {
			newList = append(newList, s.conf)
			continue
		}
		newList = append(newList, list[i])
	}
	data, _ = json.Marshal(newList)
	os.WriteFile("./config/schemes/scheme.json", data, 0777)
	return serializer.Response{
		Msg: s.conf.Name + "已更新",
	}
}

func (s SchemeService) ResetSchemeService() serializer.Response {
	var list []Scheme
	data, _ := os.ReadFile("./config/schemes/scheme.json")
	json.Unmarshal(data, &list)
	var newList []Scheme
	for i := 0; i < len(list); i++ {
		if list[i].ID == s.conf.ID {
			newList = append(newList, defaultScheme)
			continue
		}
		newList = append(newList, list[i])
	}
	data, _ = json.Marshal(newList)
	os.WriteFile("./config/schemes/scheme.json", data, 0777)
	return serializer.Response{
		Msg: s.conf.Name + "已重置",
	}
}

func GenServerConfig() {
	var list []Scheme
	var args []string
	data, _ := os.ReadFile("./config/schemes/scheme.json")
	json.Unmarshal(data, &list)
	for i := 0; i < len(list); i++ {
		file := list[i].Name
		os.Create("./config/schemes/" + file + ".txt")
		args = append(args, fmt.Sprint("maxplayers="+list[i].Maxnum))
		args = append(args, fmt.Sprint("port="+list[i].Port))
		args = append(args, fmt.Sprint("password="+list[i].Password))
		args = append(args, fmt.Sprint("motd="+list[i].Motd))
		args = append(args, fmt.Sprint("language="+list[i].Language))
		args = append(args, fmt.Sprintf("npcstream=%d", list[i].Npcstream))
		args = append(args, fmt.Sprintf("priority=%d", list[i].Priority))
		if list[i].Upnp {
			args = append(args, fmt.Sprintf("upnp=1"))
		}
		if list[i].Secure {
			args = append(args, fmt.Sprintf("secure=1"))
		}
		if list[i].Banlist {
			args = append(args, fmt.Sprintf("banlist=banlist.txt"))
		}
		if list[i].Map.Isnew {
			args = append(args, fmt.Sprint("worldname="+list[i].Map.Worldname))
			args = append(args, fmt.Sprint("autocreate="+list[i].Map.Size))
			args = append(args, fmt.Sprintf("difficulty=%d", list[i].Map.Difficult))
			if list[i].Map.Seed != "" {
				args = append(args, fmt.Sprint("seed="+list[i].Map.Seed))
			}
		} else {
			args = append(args, fmt.Sprint("worldname="+list[i].Map.Worldname))
		}
		utils.WriteServerConf(args, file)
		args = []string{}
	}
}
