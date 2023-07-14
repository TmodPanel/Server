package server

import (
	"TSM-Server/internal/serializer"
	"encoding/json"
	"fmt"
	"os"
)

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

func GetSchemeInfo(instanceName string) serializer.Response {
	var data Scheme
	path := fmt.Sprintf("./data/instance/%s/scheme.json", instanceName)

	f, err := os.ReadFile(path)
	if err != nil {
		return serializer.HandleErr(err, "打开scheme.json失败")
	}

	if err := json.Unmarshal(f, &data); err != nil {
		return serializer.HandleErr(err, "解析scheme.json失败")
	}

	return serializer.Response{
		Data: data,
		Msg:  "获取到配置方案",
	}

}

func DelScheme(instanceName string) serializer.Response {

	path := fmt.Sprintf("./data/instance/%s/scheme.json", instanceName)

	if err := os.Remove(path); err != nil {
		return serializer.HandleErr(err, "删除配置方案失败")
	}

	return serializer.Response{
		Msg: "删除配置方案成功",
	}
}

func (s *Scheme) UpdateScheme(instanceName string) serializer.Response {
	path := fmt.Sprintf("./data/instance/%s/scheme.json", instanceName)

	buf, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		return serializer.HandleErr(err, "解析配置方案失败")
	}

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		return serializer.HandleErr(err, "打开配置方案失败")
	}

	if _, err := f.Write(buf); err != nil {
		return serializer.HandleErr(err, "写入配置方案失败")
	}

	return serializer.Response{
		Msg: "更新配置方案成功",
	}
}
