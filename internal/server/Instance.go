package server

import (
	"TSM-Server/internal/serializer"
	"TSM-Server/utils"
	"os"
	"path"
)

func DelInstance(instanceName string) serializer.Response {
	instancePath := path.Join("./data/instance/", instanceName)
	if err := os.RemoveAll(instancePath); err != nil {
		return serializer.HandleErr(err, "删除实例失败")
	}
	return serializer.Response{
		Msg: "删除实例成功",
	}
}

func AddInstance(instanceName string) serializer.Response {
	instancePath := path.Join("./data/instance/", instanceName)
	if err := os.Mkdir(instancePath, os.ModePerm); err != nil {
		return serializer.HandleErr(err, "创建文件夹失败")
	}

	files := map[string]string{
		"scheme.json":      "./data/instance/default.json",
		"serverconfig.txt": "./data/instance/defaultconfig.txt",
		"banlist.txt":      "",
		"enable.json":      "",
	}

	for k, v := range files {
		if v == "" {
			if _, err := os.Create(path.Join(instancePath, k)); err != nil {
				return serializer.HandleErr(err, "创建文件失败")
			}
			continue
		}
		if err := utils.CopyFile(v, path.Join(instancePath, k)); err != nil {
			return serializer.HandleErr(err, "复制文件失败")
		}
	}

	if err := utils.WriteToFile(path.Join(instancePath, "enable.json"), "[]", 0644); err != nil {
		return serializer.HandleErr(err, "写入空的json失败")
	}

	return serializer.Response{
		Msg: "添加新的实例成功",
	}

}

func GetInstanceInfo(instanceName string) serializer.Response {
	res := GetSchemeInfo(instanceName)
	if res.Error != "" {
		return res
	}

	maps, _ := GetMapList()
	mods := GetModInfo(instanceName)

	return serializer.Response{
		Data: map[string]interface{}{
			"scheme": res.Data,
			//"boss":   bossList,
			"maps": maps,
			"mods": mods.Data,
		},
		Msg: "获取实例信息成功",
	}
}

func StartGame(instanceName string) serializer.Response {
	//if err := Run(instanceName); err != nil {
	//	return serializer.HandleErr(err, "启动失败")
	//}
	return serializer.Response{
		Msg: "启动成功",
	}
}

func StopGame(instanceName string) serializer.Response {
	//if err := Stop(instanceName); err != nil {
	//	return serializer.HandleErr(err, "停止失败")
	//}
	return serializer.Response{
		Msg: "停止成功",
	}
}

func RestartGame(instanceName string) serializer.Response {
	//if err := Restart(instanceName); err != nil {
	//	return serializer.HandleErr(err, "重启失败")
	//}
	return serializer.Response{
		Msg: "重启成功",
	}
}

func SetGameTime(instanceName string, time string) serializer.Response {
	//if err := SetTime(instanceName, time); err != nil {
	//	return serializer.HandleErr(err, "设置失败")
	//}
	return serializer.Response{
		Msg: "设置成功",
	}
}
