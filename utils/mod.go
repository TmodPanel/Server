package utils

import (
	"TSM-Server/cmd/setting"
	"fmt"
	"os"

	"github.com/goccy/go-json"
)

func DisableMod(modName string, instanceName string) error {
	enableFilePath := fmt.Sprintf("./data/instance/%s/enable.json", instanceName)
	file, err := os.ReadFile(enableFilePath)
	if err != nil {
		return err
	}

	var data []string
	err = json.Unmarshal(file, &data)
	if err != nil {
		return err
	}

	for i := 0; i < len(data); i++ {
		if data[i] == modName {
			data = append(data[:i], data[i+1:]...)
		}
	}
	file, _ = json.MarshalIndent(data, "", " ")
	return os.WriteFile(enableFilePath, file, 0644)
}

func EnableMod(modName string, instanceName string) error {
	enableFilePath := fmt.Sprintf("./data/instance/%s/enable.json", instanceName)
	file, err := os.ReadFile(enableFilePath)
	if err != nil {
		return err
	}

	var data []string
	err = json.Unmarshal(file, &data)
	if err != nil {
		return err
	}

	for i := 0; i < len(data); i++ {
		if data[i] == modName {
			return nil
		}
	}
	data = append(data, modName)

	file, _ = json.MarshalIndent(data, "", " ")
	return os.WriteFile(enableFilePath, file, 0644)
}

// DelMod wip
func DelMod(name string) error {
	//err := DisableMod(name)
	//if err != nil {
	//	return err
	//}
	return os.Remove(setting.ModPath + name + ".tmod")
}
