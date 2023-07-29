package server

import (
	"TSM-Server/internal/serializer"
	"TSM-Server/utils"
	"encoding/json"
	"fmt"
	"os"
)

type Mod struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Enable  bool   `json:"enable"`
	Version string `json:"version"`
	Size    string `json:"size"`
}

var modJSONPath = "./data/mod/mods.json"

func GetModInfo(instanceName string) serializer.Response {
	enableFilePath := fmt.Sprintf("./data/instance/%s/enable.json", instanceName)

	f, err := os.ReadFile(modJSONPath)
	if err != nil {
		return serializer.HandleErr(err, "打开mod.json失败")
	}

	var list []Mod
	if err := json.Unmarshal(f, &list); err != nil {
		return serializer.HandleErr(err, "解析mods.json失败")
	}

	enableList, err := os.ReadFile(enableFilePath)
	if err != nil {
		return serializer.HandleErr(err, "打开enable.json失败")
	}

	var enableData []string
	if err := json.Unmarshal(enableList, &enableData); err != nil {
		return serializer.HandleErr(err, "解析enable.json失败")
	}

	// 根据enable.json中的mod name，将mods.json中的enable字段设置为true
	var enableSet []Mod
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(enableData); j++ {
			if list[i].Name == enableData[j] {
				list[i].Enable = true
			}
		}
		enableSet = append(enableSet, list[i])
	}

	return serializer.Response{
		Data: enableSet,
		Msg:  "获取到mod信息",
	}
}

func (s *Mod) EnableMod(instanceName string) serializer.Response {
	flag, modName := exist(s.ID)
	if !flag {
		return serializer.Response{
			Msg: "mod不存在",
		}
	}

	err := utils.EnableMod(modName, instanceName)
	if err != nil {
		return serializer.HandleErr(err, "启用失败")
	}

	return serializer.Response{
		Msg: "启用成功",
	}
}

// DelMod wip
func (s *Mod) DelMod() serializer.Response {
	flag, _ := exist(s.ID)
	if !flag {
		return serializer.Response{
			Msg: "mod不存在",
		}
	}

	err := utils.DelMod(s.Name)
	if err != nil {
		return serializer.HandleErr(err, "删除失败")
	}

	return serializer.Response{
		Msg: "删除成功",
	}
}

func (s *Mod) DisableMod(instanceName string) serializer.Response {
	flag, modName := exist(s.ID)
	if !flag {
		return serializer.Response{
			Msg: "mod不存在",
		}
	}

	err := utils.DisableMod(modName, instanceName)
	if err != nil {
		return serializer.HandleErr(err, "禁用失败")
	}

	return serializer.Response{
		Msg: "禁用成功",
	}
}

func exist(id string) (bool, string) {
	//though read mod path to check mod exist
	//files, err := os.ReadDir(setting.ModPath)
	//if err != nil {
	//	return false
	//}
	//for i := 0; i < len(files); i++ {
	//	if files[i].Name() == name+".tmod" {
	//		return true
	//	}
	//}
	//return false

	// though read mods.json to check mod exist
	f, err := os.ReadFile(modJSONPath)
	if err != nil {
		return false, ""
	}
	var list []Mod
	err = json.Unmarshal(f, &list)
	if err != nil {
		return false, ""
	}

	for i := 0; i < len(list); i++ {
		if list[i].ID == id {
			return true, list[i].Name
		}
	}
	return false, ""
}
