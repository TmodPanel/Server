package service

import (
	"TSM-Server/internal/serializer"
	"encoding/json"
	"io/ioutil"
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
	Map       string `json:"map"`
	Language  string `json:"language"`
	Motd      string `json:"motd"`
	Priority  int    `json:"priority"`
	IP        string `json:"ip"`
	Npcstream int    `json:"npcstream"`
	Secure    bool   `json:"secure"`
	Banlist   bool   `json:"banlist"`
	Upnp      bool   `json:"upnp"`
	Steam     bool   `json:"steam"`
	Lobby     bool   `json:"lobby"`
}

var (
	defaultValue = Scheme{
		ID:        "1",
		Name:      "默认配置方案",
		Maxnum:    "8",
		Password:  "",
		Port:      "7777",
		Map:       "空岛生存带师",
		Language:  "zh-Hans",
		Motd:      "Welcome to Terraria",
		Priority:  1,
		IP:        "",
		Npcstream: 0,
		Secure:    false,
		Banlist:   false,
		Upnp:      false,
		Steam:     false,
		Lobby:     false,
	}
)

func (s SchemeService) GetSchemesInfoService() serializer.Response {
	var list []Scheme
	data, _ := ioutil.ReadFile("./config/schemes/scheme.json")
	json.Unmarshal(data, &list)
	return serializer.Response{
		Data: list,
		Msg:  "获取到配置方案信息",
	}
}

func (s SchemeService) DelSchemeService() serializer.Response {
	var list []Scheme
	data, _ := ioutil.ReadFile("./config/schemes/scheme.json")
	json.Unmarshal(data, &list)
	var newList []Scheme
	for i := 0; i < len(list); i++ {
		if list[i].ID == s.conf.ID {
			continue
		}
		newList = append(newList, list[i])
	}
	data, _ = json.Marshal(newList)
	ioutil.WriteFile("./config/schemes/scheme.json", data, 0777)
	return serializer.Response{
		Msg: s.conf.Name + "已删除",
	}
}

func (s SchemeService) AddSchemeService() serializer.Response {
	var list []Scheme
	data, _ := ioutil.ReadFile("./config/schemes/scheme.json")
	json.Unmarshal(data, &list)
	list = append(list, s.conf)
	data, _ = json.Marshal(list)
	ioutil.WriteFile("./config/schemes/scheme.json", data, 0777)
	return serializer.Response{
		Msg: s.conf.Name + "已添加",
	}
}

func (s SchemeService) UpdateSchemeService() serializer.Response {
	var list []Scheme
	data, _ := ioutil.ReadFile("./config/schemes/scheme.json")
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
	ioutil.WriteFile("./config/schemes/scheme.json", data, 0777)
	return serializer.Response{
		Msg: s.conf.Name + "已更新",
	}
}

func (s SchemeService) ResetSchemeService() serializer.Response {
	var list []Scheme
	data, _ := ioutil.ReadFile("./config/schemes/scheme.json")
	json.Unmarshal(data, &list)
	var newList []Scheme
	for i := 0; i < len(list); i++ {
		if list[i].ID == s.conf.ID {
			newList = append(newList, defaultValue)
			continue
		}
		newList = append(newList, list[i])
	}
	data, _ = json.Marshal(newList)
	ioutil.WriteFile("./config/schemes/scheme.json", data, 0777)
	return serializer.Response{
		Msg: s.conf.Name + "已重置",
	}
}
